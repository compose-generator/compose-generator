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

	// Docker Swarm compatibility (default: no)
	//dockerSwarm := utils.YesNoQuestion("Should your compose file be used for distributed deployment with Docker Swarm?", false)
	//utils.Pl(dockerSwarm)

	// Predefined stack (default: yes)
	/*usePredefinedStack := utils.YesNoQuestion("Do you want to use a predefined stack?", true)
	if usePredefinedStack {
		generateFromPredefinedTemplate(projectName, flagAdvanced, flagWithInstructions, flagWithDockerfile)
	} else {
		generateFromScratch(projectName, flagAdvanced, flagForce)
	}*/
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

	// Load stacks from templates
	templateData := parser.ParsePredefinedServices()

	// Ask for production
	alsoProduction := utils.YesNoQuestion("Also generate production configuration?", false)

	// Ask for compose file version
	compose_version := "3.9"
	if flagAdvanced {
		compose_version = utils.TextQuestionWithDefault("Docker compose file version:", compose_version)
	}

	// Initialize varMap and volumeMap
	varMap := make(map[string]string)
	varMap["PROJECT_NAME"] = projectName
	varMap["PROJECT_NAME_CONTAINER"] = strings.ReplaceAll(strings.ToLower(projectName), " ", "-")
	volumeMap := make(map[string]string)

	// Ask for frontends
	templateData, varMap, volumeMap = askForStackComponent(templateData, varMap, volumeMap, "frontend", true, "Which frontend framework do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for backends
	templateData, varMap, volumeMap = askForStackComponent(templateData, varMap, volumeMap, "backend", true, "Which backend framework do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for databases
	templateData, varMap, volumeMap = askForStackComponent(templateData, varMap, volumeMap, "database", true, "Which database engine do you want to use?", flagAdvanced, flagWithDockerfile)

	// Ask for db admin tools
	templateData, varMap, volumeMap = askForStackComponent(templateData, varMap, volumeMap, "db-admin", true, "Which db admin tool do you want to use?", flagAdvanced, flagWithDockerfile)

	if alsoProduction {
		// Ask for proxies
		templateData, varMap, volumeMap = askForStackComponent(templateData, varMap, volumeMap, "proxy", false, "Which reverse proxy you want to use?", flagAdvanced, flagWithDockerfile)

		// Ask for proxy tls helpers
		templateData, varMap, volumeMap = askForStackComponent(templateData, varMap, volumeMap, "tls-helper", false, "Which tls helper you want to use?", flagAdvanced, flagWithDockerfile)
	}

	// Generate configuration
	fmt.Print("Generating configuration ... ")
	var composeFileDev model.ComposeFile
	composeFileDev.Version = compose_version
	composeFileDev.Services = make(map[string]model.Service)
	var composeFileProd model.ComposeFile
	composeFileProd.Version = compose_version
	composeFileProd.Services = make(map[string]model.Service)

	// Delete old files
	dstPath := "."
	if flagWithInstructions {
		os.Remove(dstPath + "/README.md")
	}
	os.Remove(dstPath + "/environment.env")
	for templateType, templates := range templateData {
		for _, template := range templates {
			srcPath := utils.GetPredefinedServicesPath() + "/" + template.Dir
			// Apply all existing files of service template
			for _, f := range template.Files {
				switch f.Type {
				case "docs", "env":
					if (f.Type == "docs" && flagWithInstructions) || f.Type == "env" {
						// Append to existing file
						file_out, err1 := os.OpenFile(filepath.Join(dstPath, f.Path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						file_in, err2 := ioutil.ReadFile(filepath.Join(srcPath, f.Path))
						if err1 == nil && err2 == nil {
							file_out.WriteString(string(file_in) + "\n\n")
						}
						defer file_out.Close()
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
					bytes, _ := ioutil.ReadAll(yamlFile)
					replaced := utils.ReplaceVarsInString(string(bytes), varMap)
					service := model.Service{}
					yaml.Unmarshal([]byte(replaced), &service)
					// Add service to compose files
					if templateType != "proxy" && templateType != "tls-helper" {
						composeFileDev.Services[templateType] = service
					}
					composeFileProd.Services[templateType] = service
				}
			}
			// Generate / copy volumes
		}
	}
	utils.Done()

	// Write dev compose file
	utils.P("Saving dev configuration ... ")
	output, err1 := yaml.Marshal(&composeFileDev)
	err2 := ioutil.WriteFile("./docker-compose.yml", output, 0777)
	if err1 != nil || err2 != nil {
		utils.Error("Could not write yaml to compose file.", true)
	}
	utils.Done()

	if alsoProduction {
		// Write prod compose file
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

	/*for _, q := range templateData[index].Questions {
		if !q.Advanced || (q.Advanced && flagAdvanced) {
			switch q.Type {
			case 1: // Yes/No
				defaultValue, _ := strconv.ParseBool(q.DefaultValue)
				envMap[q.EnvVar] = strconv.FormatBool(utils.YesNoQuestion(q.Text, defaultValue))
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
					envMap[q.EnvVar] = utils.TextQuestionWithDefaultAndValidator(q.Text, q.DefaultValue, customValidator)
				} else {
					envMap[q.EnvVar] = utils.TextQuestionWithDefault(q.Text, q.DefaultValue)
				}
			}
		} else {
			envMap[q.EnvVar] = q.DefaultValue
		}
	}

	// Ask for custom volume paths
	volumesMap := make(map[string]string)
	if len(templateData[index].Volumes) > 0 {
		for _, v := range templateData[index].Volumes {
			if !v.Advanced || (v.Advanced && flagAdvanced) {
				if !v.WithDockerfile || (v.WithDockerfile && flagWithDockerfile) {
					envMap[v.EnvVar] = utils.TextQuestionWithDefault(v.Text, v.DefaultValue)
				} else {
					envMap[v.EnvVar] = v.DefaultValue
				}
			} else {
				envMap[v.EnvVar] = v.DefaultValue
			}
			volumesMap[v.DefaultValue] = envMap[v.EnvVar]
		}
		utils.Pel()
	}

	// Copy template files
	fmt.Print("Copying predefined template ... ")
	srcPath := utils.GetPredefinedServicesPath() + "/" + templateData[index].Dir
	dstPath := "."

	var err error
	for _, f := range templateData[index].Files {
		switch f.Type {
		case "compose", "env":
			os.Remove(dstPath + "/" + f.Path)
			if utils.FileExists(srcPath + "/" + f.Path) {
				err = copy.Copy(srcPath+"/"+f.Path, dstPath+"/"+f.Path)
			}
		case "docs", "docker":
			if (flagWithInstructions && f.Type == "docs") || (flagWithDockerfile && f.Type == "docker") {
				os.Remove(dstPath + "/" + f.Path)
				if utils.FileExists(srcPath + "/" + f.Path) {
					err = copy.Copy(srcPath+"/"+f.Path, dstPath+"/"+f.Path)
				}
			}
		}
	}
	if err != nil {
		utils.Error("Could not copy predefined template files.", true)
	}
	utils.Done()

	// Create volumes
	fmt.Print("Creating volumes ... ")
	for src, dst := range volumesMap {
		os.RemoveAll(dst)
		src = srcPath + src[1:]

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
	}
	utils.Done()

	// Replace variables
	fmt.Print("Applying customizations ... ")
	utils.ReplaceVarsInFile("./docker-compose.yml", envMap)
	if utils.FileExists("./environment.env") {
		utils.ReplaceVarsInFile("./environment.env", envMap)
	}
	if flagWithDockerfile && utils.FileExists("./Dockerfile") {
		utils.ReplaceVarsInFile("./Dockerfile", envMap)
	}
	utils.Done()

	if utils.FileExists("./environment.env") {
		// Generate secrets
		fmt.Print("Generating secrets ... ")
		secretsMap := utils.GenerateSecrets("./environment.env", templateData[index].Secrets)
		utils.Done()
		// Print secrets to console
		utils.Pel()
		utils.Pl("Following secrets were automatically generated:")
		for key, secret := range secretsMap {
			fmt.Print("   " + key + ": ")
			color.Yellow(secret)
		}
	}*/
}

func askForStackComponent(
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
	volumeMap map[string]string,
	component string,
	multiSelect bool,
	question string,
	flagAdvanced bool,
	flagWithDockerfile bool,
) (map[string][]model.ServiceTemplateConfig, map[string]string, map[string]string) {
	templates := templateData[component]
	items := parser.TemplateListToTemplateLabelList(templates)
	templateData[component] = []model.ServiceTemplateConfig{}
	if multiSelect {
		templateSelections := utils.MultiSelectMenuQuestionIndex(question, items)
		for _, index := range templateSelections {
			utils.Pel()
			templateData[component] = append(templateData[component], templates[index])
			varMap = getVarMapFromQuestions(varMap, templates[index].Questions, flagAdvanced)
			varMap, volumeMap = getVolumeMapFromVolumes(varMap, volumeMap, templates[index].Volumes, flagAdvanced, flagWithDockerfile)
		}
	} else {
		templateSelection := utils.MenuQuestionIndex(question, items)
		templateData[component] = append(templateData[component], templates[templateSelection])
		varMap = getVarMapFromQuestions(varMap, templates[templateSelection].Questions, flagAdvanced)
		varMap, volumeMap = getVolumeMapFromVolumes(varMap, volumeMap, templates[templateSelection].Volumes, flagAdvanced, flagWithDockerfile)
	}
	utils.Pel()
	return templateData, varMap, volumeMap
}

func getVarMapFromQuestions(varMap map[string]string, questions []model.Question, flagAdvanced bool) map[string]string {
	for _, q := range questions {
		if !q.Advanced || (q.Advanced && flagAdvanced) {
			switch q.Type {
			case 1: // Yes/No
				defaultValue, _ := strconv.ParseBool(q.DefaultValue)
				varMap[q.Variable] = strconv.FormatBool(utils.YesNoQuestion(q.Text, defaultValue))
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
					varMap[q.Variable] = utils.TextQuestionWithDefaultAndValidator(q.Text, q.DefaultValue, customValidator)
				} else {
					varMap[q.Variable] = utils.TextQuestionWithDefault(q.Text, q.DefaultValue)
				}
			}
		} else {
			varMap[q.Variable] = q.DefaultValue
		}
	}
	return varMap
}

func getVolumeMapFromVolumes(varMap map[string]string, volumeMap map[string]string, volumes []model.Volume, flagAdvanced bool, flagWithDockerfile bool) (map[string]string, map[string]string) {
	for _, v := range volumes {
		if !v.Advanced || (v.Advanced && flagAdvanced) {
			if !v.WithDockerfile || (v.WithDockerfile && flagWithDockerfile) {
				varMap[v.Variable] = utils.TextQuestionWithDefault(v.Text, v.DefaultValue)
			} else {
				varMap[v.Variable] = v.DefaultValue
			}
		} else {
			varMap[v.Variable] = v.DefaultValue
		}
		volumeMap[v.DefaultValue] = varMap[v.Variable]
	}
	return varMap, volumeMap
}

/*func generateFromScratch(projectName string, flagAdvanced bool, flagForce bool) {
	utils.ClearScreen()

	// Create custom stack
	utils.Heading("Okay. Let's create a custom stack for you!")
	utils.Pel()

	services := make(map[string]model.Service)
	i := 1
	for another := true; another; another = utils.YesNoQuestion("Generate another service?", true) {
		utils.ClearScreen()
		color.Blue("Service no. " + strconv.Itoa(i) + ":")
		service, serviceName, _ := AddService(services, flagAdvanced, flagForce, true)
		services[serviceName] = service
		utils.ClearScreen()
		color.Green("✓ Created Service '" + serviceName + "'")
		i++
	}
	utils.ClearScreen()

	// Ask for the dependencies
	var serviceNames []string
	for name := range services {
		serviceNames = append(serviceNames, name)
	}
	for serviceName, service := range services {
		currentService := service
		service.DependsOn = utils.MultiSelectMenuQuestion("On which services should your service '"+serviceName+"' depend?", utils.RemoveStringFromSlice(serviceNames, serviceName))
		services[serviceName] = currentService
	}

	// Generate compose file
	utils.P("Generating compose file ... ")
	composeFile := model.ComposeFile{}
	composeFile.Version = "3.9"
	composeFile.Services = services
	utils.Done()

	// Write to file
	utils.P("Saving compose file ... ")
	output, err1 := yaml.Marshal(&composeFile)
	err2 := ioutil.WriteFile("./docker-compose.yml", output, 0777)
	if err1 != nil || err2 != nil {
		utils.Error("Could not write yaml to compose file.", true)
	}
	utils.Done()
}*/
