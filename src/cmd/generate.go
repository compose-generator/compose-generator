package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	dcu "github.com/compose-generator/dcu"
	dcu_model "github.com/compose-generator/dcu/model"
	"github.com/fatih/color"
	"github.com/go-playground/validator"
	"github.com/otiai10/copy"
	yaml "gopkg.in/yaml.v3"

	"compose-generator/model"
	"compose-generator/parser"
	"compose-generator/util"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate a docker compose configuration
func Generate(configPath string, flagAdvanced bool, flagRun bool, flagDetached bool, flagForce bool, flagWithInstructions bool, flagWithDockerfile bool) {
	// Clear screen if in interactive mode
	if configPath == "" {
		util.ClearScreen()
	}

	// Load config file if available
	var configFile model.GenerateConfig
	projectName := "Example Project"
	if configPath != "" {
		if util.FileExists(configPath) {
			yamlFile, err1 := os.Open(configPath)
			content, err2 := ioutil.ReadAll(yamlFile)
			if err1 != nil {
				util.Error("Could not load config file. Permissions granted?", err1, true)
			}
			if err2 != nil {
				util.Error("Could not load config file. Permissions granted?", err2, true)
			}
			// Parse yaml
			yaml.Unmarshal(content, &configFile)
			projectName = configFile.ProjectName
		} else {
			util.Error("Config file could not be found", nil, true)
		}
	} else {
		// Welcome Message
		util.Heading("Welcome to Compose Generator! 👋")
		util.Pl("Please continue by answering a few questions:")
		util.Pel()

		// Ask for project name
		projectName = util.TextQuestion("What is the name of your project:")
		if projectName == "" {
			util.Error("Error. You must specify a project name!", nil, true)
		}
	}

	// Generate dynamic stack
	generateDynamicStack(configFile, projectName, flagAdvanced, flagForce, flagWithInstructions, flagWithDockerfile)

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached)
	} else {
		// Print success message
		util.Pel()
		util.SuccessMessage("🎉 Done! You now can execute \"$ docker-compose up\" to launch your app! 🎉")
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
		util.ClearScreen()
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
	util.P("Generating configuration ... ")
	composeFileDev, composeFileProd, varFiles, secrets, dockerfileMap, instString, envString := processUserInput(templateData, varMap, volMap, composeVersion, flagWithInstructions, flagWithDockerfile)
	varFiles = append(varFiles, "docker-compose.yml")
	varFiles = append(varFiles, "README.md")
	if alsoProduction {
		varFiles = append(varFiles, "docker-compose-prod.yml")
	}
	util.Done()

	// Execute safety checks
	if !flagForce {
		var existingFiles []string
		// Check files
		for _, file := range varFiles {
			if util.FileExists(file) {
				existingFiles = util.AppendStringToSliceIfMissing(existingFiles, file)
			}
		}
		// Check volumes
		for _, vol := range volMap {
			if util.FileExists(vol) {
				existingFiles = util.AppendStringToSliceIfMissing(existingFiles, vol)
			}
		}
		if len(existingFiles) > 0 {
			util.PrintSafetyWarning(len(existingFiles))
		}
	}

	// Write README.md
	if ioutil.WriteFile("./README.md", []byte(instString), 0777) != nil {
		util.Error("Could not write yaml to README file", nil, true)
	}

	// Write environment.env file
	if len(envString) > 0 {
		if ioutil.WriteFile("./environment.env", []byte(envString), 0777) != nil {
			util.Error("Could not write yaml to environment file", nil, true)
		} else {
			// Add environment.env file to .gitignore
			util.AddFileToGitignore("./environment.env")
		}
	}

	// Write dev compose file
	util.P("Saving dev configuration ... ")
	if err := dcu.SerializeToFile(composeFileDev, "./docker-compose.yml"); err != nil {
		util.Error("Could not write yaml to compose file.", err, true)
	}
	util.Done()

	// Write prod compose file
	if alsoProduction {
		util.P("Saving prod configuration ... ")
		if err := dcu.SerializeToFile(composeFileProd, "./docker-compose-prod.yml"); err != nil {
			util.Error("Could not write yaml to compose file.", err, true)
		}
		util.Done()
	}

	// Create / copy volumes
	util.P("Creating volumes ... ")
	for src, dst := range volMap {
		os.RemoveAll(dst)
		if util.FileExists(src) {
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
				util.Error("Could not copy volume files", err, true)
			}
		} else {
			// Create empty volume
			os.MkdirAll(dst, 0777)
		}
	}
	util.Done()

	// Copy dockerfiles
	for src, dst := range dockerfileMap {
		content, err1 := ioutil.ReadFile(src)
		if err1 != nil {
			util.Error("Could not read Dockerfile "+src, err1, true)
		}
		newContent := util.EvaluateConditionalSections(string(content), templateData, varMap)
		os.MkdirAll(filepath.Dir(dst), 0700)
		if ioutil.WriteFile(dst, []byte(newContent), 0777) != nil {
			util.Error("Could not write to Dockerfile "+dst, nil, true)
		}
	}

	// Replace variables
	util.P("Applying customizations ... ")
	for _, path := range varFiles {
		util.ReplaceVarsInFile(path, varMap)
	}
	util.Done()

	if flagWithDockerfile {
		// Create demo applications
		executeListOfCommands(templateData, &varMap, "DemoAppInitCmd")
	}

	// Intialize services
	executeListOfCommands(templateData, &varMap, "ServiceInitCmd")

	if len(secrets) > 0 {
		// Generate secrets
		util.P("Generating secrets ... ")
		secretsMap := util.GenerateSecrets("./environment.env", secrets)
		util.Done()
		// Print secrets to console
		util.Pel()
		util.Pl("Following secrets were automatically generated:")
		for key, secret := range secretsMap {
			util.P("🔑   " + util.ReplaceVarsInString(key, varMap) + ": ")
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
	alsoProduction := util.YesNoQuestion("Also generate production configuration?", false)

	// Ask for compose file version
	composeVersion := "3.9"
	if flagAdvanced {
		composeVersion = util.TextQuestionWithDefault("Docker compose file version:", composeVersion)
	}

	// Initialized port and volume lists
	usedPorts := []int{}
	usedVolumes := []string{}

	// Ask for frontends
	askForStackComponent(templateData, varMap, volMap, &usedPorts, &usedVolumes, "frontend", true, "Which frontend framework do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for backends
	askForStackComponent(templateData, varMap, volMap, &usedPorts, &usedVolumes, "backend", true, "Which backend framework do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for databases
	databaseCount := askForStackComponent(templateData, varMap, volMap, &usedPorts, &usedVolumes, "database", true, "Which database engine do you want to use?", flagAdvanced, flagWithDockerfile)

	if databaseCount > 0 {
		// Ask for db admin tools
		askForStackComponent(templateData, varMap, volMap, &usedPorts, &usedVolumes, "db-admin", true, "Which db admin tool do you want to use?", flagAdvanced, flagWithDockerfile)
	} else {
		(*templateData)["db-admin"] = []model.ServiceTemplateConfig{}
	}

	if alsoProduction {
		// Ask for proxies
		askForStackComponent(templateData, varMap, volMap, &usedPorts, &usedVolumes, "proxy", false, "Which reverse proxy you want to use?", flagAdvanced, flagWithDockerfile)

		// Ask for proxy tls helpers
		askForStackComponent(templateData, varMap, volMap, &usedPorts, &usedVolumes, "tls-helper", false, "Which tls helper you want to use?", flagAdvanced, flagWithDockerfile)
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
) (dcu_model.ComposeFile, dcu_model.ComposeFile, []string, []model.Secret, map[string]string, string, string) {
	// Prepare compose files
	var composeFileDev dcu_model.ComposeFile
	composeFileDev.Version = composeVersion
	composeFileDev.Services = make(map[string]dcu_model.Service)
	var composeFileProd dcu_model.ComposeFile
	composeFileProd.Version = composeVersion
	composeFileProd.Services = make(map[string]dcu_model.Service)

	// Loop through selected templates
	dstPath := "."
	var secrets []model.Secret
	var varFiles []string
	var networks []string
	dockerfileMap := make(map[string]string)
	var instString string
	var envString string

	// Read instructions header
	fileIn, err := ioutil.ReadFile(filepath.Join(util.GetPredefinedServicesPath(), "INSTRUCTIONS_HEADER.md"))
	if err != nil {
		util.Error("Cannot load instructions header file", err, true)
	}
	instString = instString + string(fileIn) + "\n\n"

	// Loop through templates
	for templateType, templates := range templateData {
		for _, template := range templates {
			srcPath := util.GetPredefinedServicesPath() + "/" + template.Dir
			// Apply all existing files of service template
			for _, f := range template.Files {
				switch f.Type {
				case "docs":
					if flagWithInstructions {
						// Append content to existing file
						fileIn, err := ioutil.ReadFile(filepath.Join(srcPath, f.Path))
						if err != nil {
							util.Error("Cannot read instructions file for template: "+template.Label, err, false)
						}
						instString = instString + string(fileIn) + "\n\n"
					}
				case "env":
					// Append content to existing file
					outPath := filepath.Join(dstPath, f.Path)
					fileIn, err := ioutil.ReadFile(filepath.Join(srcPath, f.Path))
					if err != nil {
						util.Error("Cannot read environment file for template: "+template.Label, err, false)
					}
					envString = envString + string(fileIn) + "\n\n"
					varFiles = util.AppendStringToSliceIfMissing(varFiles, outPath)
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
					content := util.EvaluateConditionalSections(string(contentBytes), templateData, varMap)
					// Replace variables
					content = util.ReplaceVarsInString(content, varMap)
					// Parse yaml
					service := dcu_model.Service{}
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
						if len(templateData["backend"]) == 0 {
							for _, template := range templateData["database"] {
								service.DependsOn = append(service.DependsOn, "database-"+template.Name)
							}
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
		composeFileDev.Networks = make(map[string]dcu_model.Network)
		composeFileProd.Networks = make(map[string]dcu_model.Network)
		for _, n := range networks {
			composeFileDev.Networks[n] = dcu_model.Network{}
			composeFileProd.Networks[n] = dcu_model.Network{}
		}
	}
	return composeFileDev, composeFileProd, varFiles, secrets, dockerfileMap, instString, envString
}

func askForStackComponent(
	templateData *map[string][]model.ServiceTemplateConfig,
	varMap *map[string]string,
	volMap *map[string]string,
	usedPorts *[]int,
	usedVolumes *[]string,
	component string,
	multiSelect bool,
	question string,
	flagAdvanced bool,
	flagWithDockerfile bool,
) (componentCount int) {
	templates := (*templateData)[component]
	items := templateListToLabelList(templates)
	itemsPreselected := templateListToPreselectedLabelList(templates, templateData)
	(*templateData)[component] = []model.ServiceTemplateConfig{}
	if multiSelect {
		templateSelections := util.MultiSelectMenuQuestionIndex(question, items, itemsPreselected)
		for _, index := range templateSelections {
			util.Pel()
			(*templateData)[component] = append((*templateData)[component], templates[index])
			getVarMapFromQuestions(varMap, usedPorts, templates[index].Questions, flagAdvanced)
			getVolumeMapFromVolumes(varMap, volMap, usedVolumes, templates[index], flagAdvanced, flagWithDockerfile)
			componentCount++
		}
	} else {
		templateSelection := util.MenuQuestionIndex(question, items)
		(*templateData)[component] = append((*templateData)[component], templates[templateSelection])
		getVarMapFromQuestions(varMap, usedPorts, templates[templateSelection].Questions, flagAdvanced)
		getVolumeMapFromVolumes(varMap, volMap, usedVolumes, templates[templateSelection], flagAdvanced, flagWithDockerfile)
		componentCount = 1
	}
	util.Pel()
	return
}

func getVarMapFromQuestions(
	varMap *map[string]string,
	usedPorts *[]int,
	questions []model.Question,
	flagAdvanced bool,
) {
	for _, q := range questions {
		defaultValue := util.ReplaceVarsInString(q.DefaultValue, *varMap)
		if !q.Advanced || (q.Advanced && flagAdvanced) {
			switch q.Type {
			case 1: // Yes/No
				defaultValue, _ := strconv.ParseBool(defaultValue)
				(*varMap)[q.Variable] = strconv.FormatBool(util.YesNoQuestion(q.Text, defaultValue))
			case 2: // Text
				if q.Validator != "" {
					var customValidator survey.Validator
					switch q.Validator {
					case "port":
						customValidator = util.PortValidator
						// Check if port was already assigned
						port, _ := strconv.Atoi(defaultValue)
						for util.SliceContainsInt(*usedPorts, port) {
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
					answer := util.TextQuestionWithDefaultAndValidator(q.Text, defaultValue, customValidator)
					(*varMap)[q.Variable] = answer
					if q.Validator == "port" {
						port, _ := strconv.Atoi(answer)
						*usedPorts = append(*usedPorts, port)
					}
				} else {
					(*varMap)[q.Variable] = util.TextQuestionWithDefault(q.Text, defaultValue)
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
	usedVolumes *[]string,
	template model.ServiceTemplateConfig,
	flagAdvanced bool,
	flagWithDockerfile bool,
) {
	srcPath := filepath.Join(util.GetPredefinedServicesPath(), template.Dir)
	for _, v := range template.Volumes {
		defaultValue := v.DefaultValue
		if util.SliceContainsString(*usedVolumes, defaultValue) {
			defaultValue = defaultValue + "-" + template.Name
		}
		if !v.Advanced || (v.Advanced && flagAdvanced) {
			if !v.WithDockerfile || (v.WithDockerfile && flagWithDockerfile) {
				(*varMap)[v.Variable] = util.TextQuestionWithDefault(v.Text, defaultValue)
			} else {
				(*varMap)[v.Variable] = defaultValue
			}
		} else {
			(*varMap)[v.Variable] = defaultValue
		}
		*usedVolumes = append(*usedVolumes, defaultValue)
		(*volMap)[filepath.Join(srcPath, v.DefaultValue)] = (*varMap)[v.Variable]
	}
}

func executeListOfCommands(
	templateData map[string][]model.ServiceTemplateConfig,
	varMap *map[string]string,
	field string,
) {
	for _, templates := range templateData {
		for _, template := range templates {
			var commands []string
			r := reflect.ValueOf(template)
			f := reflect.Indirect(r).FieldByName(field)
			for _, cmd := range f.Interface().([]string) {
				commands = append(commands, util.ReplaceVarsInString(cmd, *varMap))
			}
			if len(commands) > 0 {
				util.P("Generating demo application for '" + template.Label + "' (may take a while) ... ")
				util.ExecuteOnLinux(strings.Join(commands, "; "))
				util.Done()
			}
		}
	}
}

func templateListToLabelList(templates []model.ServiceTemplateConfig) (labels []string) {
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
			if util.EvaluateCondition(c, *templateData, nil) {
				fulfilled = true
			}
		}
		if fulfilled {
			labels = append(labels, t.Label)
		}
	}
	return
}
