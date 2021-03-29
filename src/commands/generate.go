package commands

import (
	"errors"
	"fmt"
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
func Generate(flagAdvanced bool, flagRun bool, flagDetached bool, flagForce bool, flagWithInstructions bool, flagWithDockerfile bool) {
	utils.ClearScreen()

	// Execute SafetyFileChecks
	if !flagForce {
		utils.ExecuteSafetyFileChecks(flagWithInstructions, flagWithDockerfile)
	}

	// Welcome Message
	utils.Heading("Welcome to Compose Generator!")
	utils.Pl("Please continue by answering a few questions:")
	utils.Pel()

	// Project name
	projectName := utils.TextQuestion("What is the name of your project:")
	if projectName == "" {
		utils.Error("Error. You must specify a project name!", true)
	}

	// Generate dynamic stack
	generateDynamicStack(projectName, flagAdvanced, flagWithInstructions, flagWithDockerfile)

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

func generateDynamicStack(projectName string, flagAdvanced bool, flagWithInstructions bool, flagWithDockerfile bool) {
	utils.ClearScreen()

	// Initialize varMap and volumeMap
	varMap := make(map[string]string)
	varMap["PROJECT_NAME"] = projectName
	varMap["PROJECT_NAME_CONTAINER"] = strings.ReplaceAll(strings.ToLower(projectName), " ", "-")
	volMap := make(map[string]string)

	// Load configurations of service templates
	templateData := parser.ParsePredefinedServices()

	// Ask user decisions
	composeVersion, alsoProduction := askForUserInput(&templateData, &varMap, &volMap, flagAdvanced, flagWithDockerfile)

	// Delete old files
	fmt.Print("Cleaning up ... ")
	if flagWithInstructions {
		os.Remove("./README.md")
	}
	os.Remove("./environment.env")
	utils.Done()

	// Generate configuration
	fmt.Print("Generating configuration ... ")
	composeFileDev, composeFileProd, secrets := processUserInput(templateData, varMap, volMap, composeVersion, flagWithInstructions, flagWithDockerfile)
	utils.Done()

	// Write dev compose file
	utils.P("Saving dev configuration ... ")
	output, err1 := yaml.Marshal(&composeFileDev)
	err2 := ioutil.WriteFile("./docker-compose.yml", output, 0777)
	if err1 != nil || err2 != nil {
		utils.Error("Could not write yaml to compose file.", true)
	}
	utils.Done()

	// Write prod compose file
	if alsoProduction {
		utils.P("Saving prod configuration ... ")
		output, err1 := yaml.Marshal(&composeFileProd)
		err2 := ioutil.WriteFile("./docker-compose-prod.yml", output, 0777)
		if err1 != nil || err2 != nil {
			utils.Error("Could not write yaml to compose file.", true)
		}
		utils.Done()
	}

	// Replace variables
	fmt.Print("Applying customizations ... ")
	utils.ReplaceVarsInFile("./docker-compose.yml", varMap)
	if alsoProduction {
		utils.ReplaceVarsInFile("./docker-compose-prod.yml", varMap)
	}
	if utils.FileExists("./environment.env") {
		utils.ReplaceVarsInFile("./environment.env", varMap)
	}
	utils.Done()

	// Create / copy volumes
	fmt.Print("Creating volumes ... ")
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
				utils.Error("Could not copy volume files.", true)
			}
		} else {
			// Create empty volume
			os.MkdirAll(dst, 0777)
		}
	}
	utils.Done()

	// Create example applications
	fmt.Print("Generating demo applications (may take a while) ... ")
	for _, templates := range templateData {
		for _, template := range templates {
			for _, cmd := range template.ExampleAppInitCmd {
				utils.ExecuteAndWait(strings.Split(cmd, " ")...)
			}
		}
	}
	utils.Done()

	if utils.FileExists("./environment.env") {
		// Generate secrets
		fmt.Print("Generating secrets ... ")
		secretsMap := utils.GenerateSecrets("./environment.env", secrets)
		utils.Done()
		// Print secrets to console
		utils.Pel()
		utils.Pl("Following secrets were automatically generated:")
		for key, secret := range secretsMap {
			fmt.Print("   " + key + ": ")
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

	// Ask for frontends
	askForStackComponent(templateData, varMap, volMap, "frontend", true, "Which frontend framework do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for backends
	askForStackComponent(templateData, varMap, volMap, "backend", true, "Which backend framework do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for databases
	askForStackComponent(templateData, varMap, volMap, "database", true, "Which database engine do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for db admin tools
	askForStackComponent(templateData, varMap, volMap, "db-admin", true, "Which db admin tool do you want to use?", flagAdvanced, flagWithDockerfile)

	if alsoProduction {
		// Ask for proxies
		askForStackComponent(templateData, varMap, volMap, "proxy", false, "Which reverse proxy you want to use?", flagAdvanced, flagWithDockerfile)

		// Ask for proxy tls helpers
		askForStackComponent(templateData, varMap, volMap, "tls-helper", false, "Which tls helper you want to use?", flagAdvanced, flagWithDockerfile)
	} else {
		(*templateData)["proxy"] = []model.ServiceTemplateConfig{}
		(*templateData)["tls-helper"] = []model.ServiceTemplateConfig{}
	}
	return composeVersion, alsoProduction
}

func processUserInput(
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
	volumeMap map[string]string,
	composeVersion string,
	flagWithInstructions bool,
	flagWithDockerfile bool,
) (model.ComposeFile, model.ComposeFile, []model.Secret) {
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
	var networks []string
	for templateType, templates := range templateData {
		for _, template := range templates {
			srcPath := utils.GetPredefinedServicesPath() + "/" + template.Dir
			// Apply all existing files of service template
			for _, f := range template.Files {
				switch f.Type {
				case "docs", "env":
					if (f.Type == "docs" && flagWithInstructions) || f.Type == "env" {
						// Append content to existing file
						fileOut, err1 := os.OpenFile(filepath.Join(dstPath, f.Path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						fileIn, err2 := ioutil.ReadFile(filepath.Join(srcPath, f.Path))
						replaced := utils.ReplaceVarsInString(string(fileIn), varMap)
						if err1 == nil && err2 == nil {
							fileOut.WriteString(replaced + "\n\n")
						}
						defer fileOut.Close()
					}
				case "docker":
					if flagWithDockerfile {
						// Copy dockerfile
						os.Remove(filepath.Join(dstPath, f.Path))
						copy.Copy(filepath.Join(srcPath, f.Path), filepath.Join(dstPath, f.Path))
						utils.ReplaceVarsInFile(filepath.Join(dstPath, f.Path), varMap)
					}
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
						if len(templateData["backend"]) > 0 {
							service.DependsOn = []string{"backend"}
						}
					case "backend":
						if len(templateData["database"]) > 0 {
							service.DependsOn = []string{"database"}
						}
					case "db-admin":
						if len(templateData["database"]) > 0 {
							service.DependsOn = []string{"database"}
						}
					case "tls-helper":
						if len(templateData["proxy"]) > 0 {
							service.DependsOn = []string{"proxy"}
						}
					}
					// Add service to compose files
					if templateType != "proxy" && templateType != "tls-helper" {
						composeFileDev.Services[templateType] = service
					}
					composeFileProd.Services[templateType] = service
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
	return composeFileDev, composeFileProd, secrets
}

func askForStackComponent(
	templateData *map[string][]model.ServiceTemplateConfig,
	varMap *map[string]string,
	volMap *map[string]string,
	component string,
	multiSelect bool,
	question string,
	flagAdvanced bool,
	flagWithDockerfile bool,
) {
	templates := (*templateData)[component]
	items := parser.TemplateListToTemplateLabelList(templates)
	(*templateData)[component] = []model.ServiceTemplateConfig{}
	if multiSelect {
		templateSelections := utils.MultiSelectMenuQuestionIndex(question, items)
		for _, index := range templateSelections {
			utils.Pel()
			(*templateData)[component] = append((*templateData)[component], templates[index])
			getVarMapFromQuestions(varMap, templates[index].Questions, flagAdvanced)
			getVolumeMapFromVolumes(varMap, volMap, templates[index], flagAdvanced, flagWithDockerfile)
		}
	} else {
		templateSelection := utils.MenuQuestionIndex(question, items)
		(*templateData)[component] = append((*templateData)[component], templates[templateSelection])
		getVarMapFromQuestions(varMap, templates[templateSelection].Questions, flagAdvanced)
		getVolumeMapFromVolumes(varMap, volMap, templates[templateSelection], flagAdvanced, flagWithDockerfile)
	}
	utils.Pel()
}

func getVarMapFromQuestions(
	varMap *map[string]string,
	questions []model.Question,
	flagAdvanced bool,
) {
	for _, q := range questions {
		if !q.Advanced || (q.Advanced && flagAdvanced) {
			switch q.Type {
			case 1: // Yes/No
				defaultValue, _ := strconv.ParseBool(q.DefaultValue)
				(*varMap)[q.Variable] = strconv.FormatBool(utils.YesNoQuestion(q.Text, defaultValue))
			case 2: // Text
				if q.Validator != "" {
					var customValidator survey.Validator
					switch q.Validator {
					case "port":
						customValidator = utils.PortValidator
					default:
						customValidator = func(val interface{}) error {
							validate := validator.New()
							if validate.Var(val.(string), "required,"+q.Validator) != nil {
								return errors.New("please provide a valid input")
							}
							return nil
						}
					}
					(*varMap)[q.Variable] = utils.TextQuestionWithDefaultAndValidator(q.Text, q.DefaultValue, customValidator)
				} else {
					(*varMap)[q.Variable] = utils.TextQuestionWithDefault(q.Text, q.DefaultValue)
				}
			}
		} else {
			(*varMap)[q.Variable] = q.DefaultValue
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
				fmt.Println("Name: '" + template.Name + "'")
				fmt.Println("t: '" + condition[12:] + "'")
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
