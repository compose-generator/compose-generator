package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Knetic/govaluate"
	"github.com/fatih/color"
	"github.com/go-playground/validator"
	"github.com/otiai10/copy"
	yaml "gopkg.in/yaml.v3"

	"compose-generator/model"
	"compose-generator/parser"
	"compose-generator/utils"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate a docker compose configuration
func Generate(configPath string, flagAdvanced bool, flagRun bool, flagDetached bool, flagForce bool, flagWithInstructions bool, flagWithDockerfile bool) {
	// Clear screen if in interactive mode
	if configPath == "" {
		utils.ClearScreen()
	}

	// Load config file if available
	var configFile model.GenerateConfig
	projectName := "Example Project"
	if configPath != "" {
		if utils.FileExists(configPath) {
			yamlFile, err1 := os.Open(configPath)
			content, err2 := ioutil.ReadAll(yamlFile)
			if err1 != nil {
				utils.Error("Could not load config file. Permissions granted?", err1, true)
			}
			if err2 != nil {
				utils.Error("Could not load config file. Permissions granted?", err2, true)
			}
			// Parse yaml
			yaml.Unmarshal(content, &configFile)
			projectName = configFile.ProjectName
		} else {
			utils.Error("Config file could not be found", nil, true)
		}
	} else {
		// Welcome Message
		utils.Heading("Welcome to Compose Generator! 👋")
		utils.Pl("Please continue by answering a few questions:")
		utils.Pel()

		// Ask for project name
		projectName = utils.TextQuestion("What is the name of your project:")
		if projectName == "" {
			utils.Error("Error. You must specify a project name!", nil, true)
		}
	}

	// Generate dynamic stack
	generateDynamicStack(configFile, projectName, flagAdvanced, flagForce, flagWithInstructions, flagWithDockerfile)

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		utils.DockerComposeUp(flagDetached)
	} else {
		// Print success message
		utils.Pel()
		utils.SuccessMessage("🎉 Done! You now can execute \"$ docker-compose up\" to launch your app! 🎉")
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func generateDynamicStack(
	configFile model.GenerateConfig,
	projectName string,
	flagAdvanced bool,
	flagForce bool,
	flagWithInstructions bool,
	flagWithDockerfile bool,
) {
	// Clear screen if in interactive mode
	if configFile.ProjectName == "" {
		utils.ClearScreen()
	}

	// Initialize varMap and volumeMap
	varMap := make(map[string]string)
	varMap["PROJECT_NAME"] = projectName
	varMap["PROJECT_NAME_CONTAINER"] = strings.ReplaceAll(strings.ToLower(projectName), " ", "-")
	volMap := make(map[string]string)

	// Load configurations of service templates
	templateData := parser.ParsePredefinedServices()

	composeVersion := configFile.ComposeVersion
	alsoProduction := configFile.AlsoProduction
	if configFile.ProjectName != "" {
		// Provide selected template data from config file
		serviceConfig := configFile.ServiceConfig
		selectedTemplateData := map[string][]model.ServiceTemplateConfig{}
		for templateType, templates := range templateData {
			selectedTemplateData[templateType] = []model.ServiceTemplateConfig{}
			for _, template := range templates {
				// Loop through services
				for _, service := range serviceConfig {
					if service.Type == templateType && strings.HasSuffix(template.Dir, service.Service) {
						selectedTemplateData[templateType] = append(selectedTemplateData[templateType], template)
						// Loop through questions and add default values to varMap
						for _, question := range template.Questions {
							varMap[question.Variable] = question.DefaultValue
						}
						// Override with params
						for varName, varValue := range service.Params {
							varMap[varName] = varValue
							// Loop through volumes
							for _, volume := range template.Volumes {
								if volume.Variable == varName {
									volMap[filepath.Join(template.Dir, volume.DefaultValue)] = varValue
									break
								}
							}
						}
						break
					}
				}
			}
		}
		templateData = selectedTemplateData
	} else {
		// Ask user decisions
		composeVersion, alsoProduction = askForUserInput(&templateData, &varMap, &volMap, flagAdvanced, flagWithDockerfile)
	}

	// Generate configuration
	utils.P("Generating configuration ... ")
	composeFileDev, composeFileProd, varFiles, secrets, dockerfileMap, instString, envString := processUserInput(templateData, varMap, volMap, composeVersion, flagWithInstructions, flagWithDockerfile)
	varFiles = append(varFiles, "docker-compose.yml")
	if alsoProduction {
		varFiles = append(varFiles, "docker-compose-prod.yml")
	}
	utils.Done()

	// Execute safety checks
	if !flagForce {
		var existingFiles []string
		// Check files
		for _, file := range varFiles {
			if utils.FileExists(file) {
				existingFiles = utils.AppendStringToSliceIfMissing(existingFiles, file)
			}
		}
		// Check volumes
		for _, vol := range volMap {
			if utils.FileExists(vol) {
				existingFiles = utils.AppendStringToSliceIfMissing(existingFiles, vol)
			}
		}
		if len(existingFiles) > 0 {
			utils.PrintSafetyWarning(len(existingFiles))
		}
	}

	// Write README & environment file
	if ioutil.WriteFile("./README.md", []byte(instString), 0777) != nil {
		utils.Error("Could not write yaml to README file.", nil, true)
	}
	if ioutil.WriteFile("./environment.env", []byte(envString), 0777) != nil {
		utils.Error("Could not write yaml to environment file.", nil, true)
	}

	// Copy dockerfiles
	for src, dst := range dockerfileMap {
		os.Remove(dst)
		copy.Copy(src, dst)
	}

	// Write dev compose file
	utils.P("Saving dev configuration ... ")
	output, err1 := yaml.Marshal(&composeFileDev)
	err2 := ioutil.WriteFile("./docker-compose.yml", output, 0777)
	if err1 != nil {
		utils.Error("Could not write yaml to compose file.", err1, true)
	}
	if err2 != nil {
		utils.Error("Could not write yaml to compose file.", err2, true)
	}
	utils.Done()

	// Write prod compose file
	if alsoProduction {
		utils.P("Saving prod configuration ... ")
		output, err1 := yaml.Marshal(&composeFileProd)
		err2 := ioutil.WriteFile("./docker-compose-prod.yml", output, 0777)
		if err1 != nil {
			utils.Error("Could not write yaml to compose file.", err1, true)
		}
		if err2 != nil {
			utils.Error("Could not write yaml to compose file.", err2, true)
		}
		utils.Done()
	}

	// Create / copy volumes
	utils.P("Creating volumes ... ")
	for src, dst := range volMap {
		os.RemoveAll(dst)
		if utils.FileExists(src) {
			// Copy contents of volume
			opt := copy.Options{
				Skip: func(src string) (bool, error) {
					return strings.HasSuffix(src, ".gitkeep"), nil
				},
				OnDirExists: func(src string, dst string) copy.DirExistsAction {
					return copy.Replace
				},
			}
			err := copy.Copy(src, dst, opt)
			if err != nil {
				utils.Error("Could not copy volume files", err, true)
			}
		} else {
			// Create empty volume
			os.MkdirAll(dst, 0777)
		}
	}
	utils.Done()

	// Replace variables
	utils.P("Applying customizations ... ")
	for _, path := range varFiles {
		utils.ReplaceVarsInFile(path, varMap)
	}
	utils.Done()

	if flagWithDockerfile {
		// Create example applications
		for _, templates := range templateData {
			for _, template := range templates {
				var commands []string
				for _, cmd := range template.ExampleAppInitCmd {
					commands = append(commands, utils.ReplaceVarsInString(cmd, varMap))
				}
				if len(commands) > 0 {
					utils.P("Generating demo applications (may take a while) ... ")
					utils.ExecuteOnLinux(strings.Join(commands, "; "))
					utils.Done()
				}
			}
		}
	}

	if len(secrets) > 0 {
		// Generate secrets
		utils.P("Generating secrets ... ")
		secretsMap := utils.GenerateSecrets("./environment.env", secrets)
		utils.Done()
		// Print secrets to console
		utils.Pel()
		utils.Pl("Following secrets were automatically generated:")
		for key, secret := range secretsMap {
			utils.P("🔑   " + utils.ReplaceVarsInString(key, varMap) + ": ")
			color.Yellow(secret)
		}
	}
}

func askForUserInput(
	templateData *map[string][]model.ServiceTemplateConfig,
	varMap *map[string]string,
	volMap *map[string]string,
	flagAdvanced bool,
	flagWithDockerfile bool,
) (string, bool) {
	// Ask for production
	alsoProduction := utils.YesNoQuestion("Also generate production configuration?", false)

	// Ask for compose file version
	composeVersion := "3.9"
	if flagAdvanced {
		composeVersion = utils.TextQuestionWithDefault("Docker compose file version:", composeVersion)
	}

	// Initialized ports list
	usedPorts := []int{}

	// Ask for frontends
	askForStackComponent(templateData, varMap, volMap, &usedPorts, "frontend", true, "Which frontend framework do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for backends
	askForStackComponent(templateData, varMap, volMap, &usedPorts, "backend", true, "Which backend framework do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for databases
	databaseCount := askForStackComponent(templateData, varMap, volMap, &usedPorts, "database", true, "Which database engine do you want to use?", flagAdvanced, flagWithDockerfile)

	if databaseCount > 0 {
		// Ask for db admin tools
		askForStackComponent(templateData, varMap, volMap, &usedPorts, "db-admin", true, "Which db admin tool do you want to use?", flagAdvanced, flagWithDockerfile)
	} else {
		(*templateData)["db-admin"] = []model.ServiceTemplateConfig{}
	}

	if alsoProduction {
		// Ask for proxies
		askForStackComponent(templateData, varMap, volMap, &usedPorts, "proxy", false, "Which reverse proxy you want to use?", flagAdvanced, flagWithDockerfile)

		// Ask for proxy tls helpers
		askForStackComponent(templateData, varMap, volMap, &usedPorts, "tls-helper", false, "Which tls helper you want to use?", flagAdvanced, flagWithDockerfile)
	} else {
		(*templateData)["proxy"] = []model.ServiceTemplateConfig{}
		(*templateData)["tls-helper"] = []model.ServiceTemplateConfig{}
	}
	return composeVersion, alsoProduction
}

func processUserInput(
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
	volMap map[string]string,
	composeVersion string,
	flagWithInstructions bool,
	flagWithDockerfile bool,
) (model.ComposeFile, model.ComposeFile, []string, []model.Secret, map[string]string, string, string) {
	// Prepare compose files
	var composeFileDev model.ComposeFile
	composeFileDev.Version = composeVersion
	composeFileDev.Services = make(map[string]model.Service)
	var composeFileProd model.ComposeFile
	composeFileProd.Version = composeVersion
	composeFileProd.Services = make(map[string]model.Service)

	// Loop through selected templates
	dstPath := "."
	var secrets []model.Secret
	var varFiles []string
	var networks []string
	dockerfileMap := make(map[string]string)
	var instString string
	var envString string
	for templateType, templates := range templateData {
		for _, template := range templates {
			srcPath := utils.GetPredefinedServicesPath() + "/" + template.Dir
			// Apply all existing files of service template
			for _, f := range template.Files {
				switch f.Type {
				case "docs":
					if flagWithInstructions {
						// Append content to existing file
						outPath := filepath.Join(dstPath, f.Path)
						fileIn, err := ioutil.ReadFile(filepath.Join(srcPath, f.Path))
						if err != nil {
							utils.Error("Cannot read instructions file for template: "+template.Label, err, false)
						}
						instString = instString + string(fileIn) + "\n\n"
						varFiles = utils.AppendStringToSliceIfMissing(varFiles, outPath)
					}
				case "env":
					// Append content to existing file
					outPath := filepath.Join(dstPath, f.Path)
					fileIn, err := ioutil.ReadFile(filepath.Join(srcPath, f.Path))
					if err != nil {
						utils.Error("Cannot read environment file for template: "+template.Label, err, false)
					}
					envString = envString + string(fileIn) + "\n\n"
					varFiles = utils.AppendStringToSliceIfMissing(varFiles, outPath)
				case "docker":
					if flagWithDockerfile {
						// Check if Dockerfile is inside of a volume
						absDockerfileSrc, _ := filepath.Abs(filepath.Join(srcPath, f.Path))
						dockerfileDst := filepath.Join(dstPath, f.Path)
						for volSrc, volDst := range volMap {
							absVolSrc, _ := filepath.Abs(volSrc)
							if strings.Contains(absDockerfileSrc, absVolSrc) {
								dockerfileDst = volDst + absDockerfileSrc[len(absVolSrc):]
							}
						}
						dockerfileMap[absDockerfileSrc] = dockerfileDst
						varFiles = append(varFiles, dockerfileDst)
					}
				case "config":
					// Check if config file is inside of a volume
					absConfigSrc, _ := filepath.Abs(filepath.Join(srcPath, f.Path))
					configDst := filepath.Join(dstPath, f.Path)
					for volSrc, volDst := range volMap {
						absVolSrc, _ := filepath.Abs(volSrc)
						if strings.Contains(absConfigSrc, absVolSrc) {
							configDst = volDst + absConfigSrc[len(absVolSrc):]
						}
					}
					varFiles = append(varFiles, configDst)
				case "service":
					// Load service file
					yamlFile, _ := os.Open(filepath.Join(srcPath, f.Path))
					contentBytes, _ := ioutil.ReadAll(yamlFile)
					// Evaluate conditional sections
					content := evaluateConditionalSections(string(contentBytes), templateData, varMap)
					// Replace variables
					content = utils.ReplaceVarsInString(content, varMap)
					// Parse yaml
					service := model.Service{}
					yaml.Unmarshal([]byte(content), &service)
					// Get networks
					networks = append(networks, service.Networks...)
					// Add depends on
					switch templateType {
					case "frontend":
						service.DependsOn = []string{}
						for _, template := range templateData["backend"] {
							service.DependsOn = append(service.DependsOn, "backend-"+template.Name)
						}
					case "backend":
						service.DependsOn = []string{}
						for _, template := range templateData["database"] {
							service.DependsOn = append(service.DependsOn, "database-"+template.Name)
						}
					case "db-admin":
						service.DependsOn = []string{}
						for _, template := range templateData["database"] {
							service.DependsOn = append(service.DependsOn, "database-"+template.Name)
						}
					case "tls-helper":
						service.DependsOn = []string{}
						for _, template := range templateData["proxy"] {
							service.DependsOn = append(service.DependsOn, "proxy-"+template.Name)
						}
					}
					// Add service to compose files
					if templateType != "proxy" && templateType != "tls-helper" {
						composeFileDev.Services[templateType+"-"+template.Name] = service
					}
					composeFileProd.Services[templateType+"-"+template.Name] = service
				}
			}
			// Get secrets
			secrets = append(secrets, template.Secrets...)
		}
	}
	// Apply networks
	if len(networks) > 0 {
		composeFileDev.Networks = make(map[string]model.Network)
		composeFileProd.Networks = make(map[string]model.Network)
		for _, n := range networks {
			composeFileDev.Networks[n] = model.Network{}
			composeFileProd.Networks[n] = model.Network{}
		}
	}
	return composeFileDev, composeFileProd, varFiles, secrets, dockerfileMap, instString, envString
}

func askForStackComponent(
	templateData *map[string][]model.ServiceTemplateConfig,
	varMap *map[string]string,
	volMap *map[string]string,
	usedPorts *[]int,
	component string,
	multiSelect bool,
	question string,
	flagAdvanced bool,
	flagWithDockerfile bool,
) (componentCount int) {
	templates := (*templateData)[component]
	items := templateListToTemplateLabelList(templates)
	itemsPreselected := templateListToPreselectedLabelList(templates, templateData)
	(*templateData)[component] = []model.ServiceTemplateConfig{}
	if multiSelect {
		templateSelections := utils.MultiSelectMenuQuestionIndex(question, items, itemsPreselected)
		for _, index := range templateSelections {
			utils.Pel()
			(*templateData)[component] = append((*templateData)[component], templates[index])
			getVarMapFromQuestions(varMap, usedPorts, templates[index].Questions, flagAdvanced)
			getVolumeMapFromVolumes(varMap, volMap, templates[index], flagAdvanced, flagWithDockerfile)
			componentCount++
		}
	} else {
		templateSelection := utils.MenuQuestionIndex(question, items)
		(*templateData)[component] = append((*templateData)[component], templates[templateSelection])
		getVarMapFromQuestions(varMap, usedPorts, templates[templateSelection].Questions, flagAdvanced)
		getVolumeMapFromVolumes(varMap, volMap, templates[templateSelection], flagAdvanced, flagWithDockerfile)
		componentCount = 1
	}
	utils.Pel()
	return
}

func getVarMapFromQuestions(
	varMap *map[string]string,
	usedPorts *[]int,
	questions []model.Question,
	flagAdvanced bool,
) {
	for _, q := range questions {
		defaultValue := utils.ReplaceVarsInString(q.DefaultValue, *varMap)
		if !q.Advanced || (q.Advanced && flagAdvanced) {
			switch q.Type {
			case 1: // Yes/No
				defaultValue, _ := strconv.ParseBool(defaultValue)
				(*varMap)[q.Variable] = strconv.FormatBool(utils.YesNoQuestion(q.Text, defaultValue))
			case 2: // Text
				if q.Validator != "" {
					var customValidator survey.Validator
					switch q.Validator {
					case "port":
						customValidator = utils.PortValidator
						// Check if port was already assigned
						port, _ := strconv.Atoi(defaultValue)
						for utils.SliceContainsInt(*usedPorts, port) {
							port = port + 1
						}
						defaultValue = strconv.Itoa(port)
					default:
						customValidator = func(val interface{}) error {
							validate := validator.New()
							if validate.Var(val.(string), "required,"+q.Validator) != nil {
								return errors.New("please provide a valid input")
							}
							return nil
						}
					}
					answer := utils.TextQuestionWithDefaultAndValidator(q.Text, defaultValue, customValidator)
					(*varMap)[q.Variable] = answer
					if q.Validator == "port" {
						port, _ := strconv.Atoi(answer)
						*usedPorts = append(*usedPorts, port)
					}
				} else {
					(*varMap)[q.Variable] = utils.TextQuestionWithDefault(q.Text, defaultValue)
				}
			}
		} else {
			(*varMap)[q.Variable] = defaultValue
		}
	}
}

func getVolumeMapFromVolumes(
	varMap *map[string]string,
	volMap *map[string]string,
	template model.ServiceTemplateConfig,
	flagAdvanced bool,
	flagWithDockerfile bool,
) {
	srcPath := filepath.Join(utils.GetPredefinedServicesPath(), template.Dir)
	for _, v := range template.Volumes {
		if !v.Advanced || (v.Advanced && flagAdvanced) {
			if !v.WithDockerfile || (v.WithDockerfile && flagWithDockerfile) {
				(*varMap)[v.Variable] = utils.TextQuestionWithDefault(v.Text, v.DefaultValue)
			} else {
				(*varMap)[v.Variable] = v.DefaultValue
			}
		} else {
			(*varMap)[v.Variable] = v.DefaultValue
		}
		(*volMap)[filepath.Join(srcPath, v.DefaultValue)] = (*varMap)[v.Variable]
	}
}

func evaluateConditionalSections(
	content string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) string {
	rows := strings.Split(content, "\n")
	uncommenting := false
	for i, row := range rows {
		if strings.HasPrefix(row, "#! if ") {
			// Conditional section found -> check condition
			condition := row[6:strings.Index(row, " {")]
			uncommenting = evaluateCondition(condition, templateData, varMap)
		} else if strings.HasPrefix(row, "#! }") {
			uncommenting = false
		} else if uncommenting {
			rows[i] = row[3:]
		}
	}
	return strings.Join(rows, "\n")
}

func evaluateCondition(
	condition string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) bool {
	if strings.HasPrefix(condition, "has service ") {
		for _, templates := range templateData {
			for _, template := range templates {
				if template.Name == condition[12:] {
					return true
				}
			}
		}
	} else if strings.HasPrefix(condition, "has ") {
		return len(templateData[condition[4:]]) > 0
	} else if strings.HasPrefix(condition, "var.") {
		condition = condition[4:]
		expr, err1 := govaluate.NewEvaluableExpression(condition)
		parameters := make(map[string]interface{}, 8)
		for varName, varValue := range varMap {
			parameters[varName] = varValue
		}
		result, err2 := expr.Evaluate(parameters)
		return result.(bool) && err1 != nil && err2 != nil
	}
	return false
}

func templateListToTemplateLabelList(templates []model.ServiceTemplateConfig) (labels []string) {
	for _, t := range templates {
		labels = append(labels, t.Label)
	}
	return
}

func templateListToPreselectedLabelList(templates []model.ServiceTemplateConfig, templateData *map[string][]model.ServiceTemplateConfig) (labels []string) {
	for _, t := range templates {
		conditions := strings.Split(t.Preselected, "|")
		fulfilled := false
		for _, c := range conditions {
			if evaluateCondition(c, *templateData, nil) {
				fulfilled = true
			}
		}
		if fulfilled {
			labels = append(labels, t.Label)
		}
	}
	return
}
