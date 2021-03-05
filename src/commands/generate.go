package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/otiai10/copy"
	yaml "gopkg.in/yaml.v3"

	"compose-generator/model"
	"compose-generator/parser"
	"compose-generator/utils"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate a docker compose configuration
func Generate(flagAdvanced bool, flagRun bool, flagDetached bool, flagForce bool) {
	// Execute SafetyFileChecks
	if !flagForce {
		utils.ExecuteSafetyFileChecks()
	}

	// Welcome Message
	utils.Heading("Welcome to Compose Generator!")
	fmt.Println("Please continue by answering a few questions:")
	fmt.Println()

	// Project name
	projectName := utils.TextQuestion("What is the name of your project: ")
	if projectName == "" {
		utils.Error("Error. You must specify a project name!", true)
	}

	// Docker Swarm compatibility (default: no)
	//dockerSwarm := utils.YesNoQuestion("Should your compose file be used for distributed deployment with Docker Swarm?", false)
	//fmt.Println(dockerSwarm)

	// Predefined stack (default: yes)
	usePredefinedStack := utils.YesNoQuestion("Do you want to use a predefined stack?", true)
	if usePredefinedStack {
		generateFromPredefinedTemplate(projectName, flagAdvanced)
	} else {
		generateFromScratch(projectName, flagAdvanced, flagForce)
	}

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		utils.DockerComposeUp(flagDetached)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func generateFromPredefinedTemplate(projectName string, flagAdvanced bool) {
	// Load stacks from templates
	templateData := parser.ParsePredefinedTemplates()
	// Predefined stack menu
	var items []string
	for _, t := range templateData {
		items = append(items, t.Label)
	}
	index := utils.MenuQuestionIndex("Predefined software stack", items)
	fmt.Println()

	// Ask configured questions to the user
	envMap := make(map[string]string)
	envMap["PROJECT_NAME"] = projectName
	envMap["PROJECT_NAME_CONTAINER"] = strings.ReplaceAll(strings.ToLower(projectName), " ", "-")

	for _, q := range templateData[index].Questions {
		if !q.Advanced || (q.Advanced && flagAdvanced) {
			switch q.Type {
			case 1: // Yes/No
				defaultValue, _ := strconv.ParseBool(q.DefaultValue)
				envMap[q.EnvVar] = strconv.FormatBool(utils.YesNoQuestion(q.Text, defaultValue))
			case 2: // Text
				envMap[q.EnvVar] = utils.TextQuestionWithDefault(q.Text, q.DefaultValue)
			}
		} else {
			envMap[q.EnvVar] = q.DefaultValue
		}
	}
	fmt.Println()
	// Ask for custom volume paths
	volumesMap := make(map[string]string)
	for _, v := range templateData[index].Volumes {
		if !v.Advanced || (v.Advanced && flagAdvanced) {
			envMap[v.EnvVar] = utils.TextQuestionWithDefault(v.Text, v.DefaultValue)
		} else {
			envMap[v.EnvVar] = v.DefaultValue
		}
		volumesMap[v.DefaultValue] = envMap[v.EnvVar]
	}

	// Copy template files
	fmt.Print("Copying template ...")
	srcPath := utils.GetPredefinedTemplatesPath() + "/" + templateData[index].Dir
	dstPath := "."

	os.Remove(dstPath + "/docker-compose.yml")
	os.Remove(dstPath + "/environment.env")

	err1 := copy.Copy(srcPath+"/docker-compose.yml", dstPath+"/docker-compose.yml")
	err2 := copy.Copy(srcPath+"/environment.env", dstPath+"/environment.env")
	if err1 != nil || err2 != nil {
		utils.Error("Could not copy template files.", true)
	}

	color.Green(" done")

	// Create volumes
	fmt.Print("Creating volumes ...")

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

	color.Green(" done")

	// Replace variables
	fmt.Print("Applying customizations ...")
	utils.ReplaceVarsInFile("./docker-compose.yml", envMap)
	utils.ReplaceVarsInFile("./environment.env", envMap)
	color.Green(" done")

	// Generate secrets
	fmt.Print("Generating secrets ...")
	secretsMap := utils.GenerateSecrets("./environment.env", templateData[index].Secrets)
	color.Green(" done")
	// Print secrets to console
	fmt.Println()
	fmt.Println("Following secrets were automatically generated:")
	for key, secret := range secretsMap {
		fmt.Print("   " + key + ": ")
		color.Yellow(secret)
	}
}

func generateFromScratch(projectName string, flagAdvanced bool, flagForce bool) {
	// Create custom stack
	utils.Heading("Let's create a custom stack for you!")
	fmt.Println()

	services := make(map[string]model.Service)
AddService:
	service, serviceName, _ := AddService(services, flagAdvanced, flagForce, true)
	services[serviceName] = service
	fmt.Println()

	// Ask for another service
	if utils.YesNoQuestion("Generate another service?", true) {
		goto AddService
	}
	fmt.Println()

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
	fmt.Println("Generating compose file ...")
	composeFile := model.ComposeFile{}
	composeFile.Version = "3.8"
	composeFile.Services = services
	color.Green(" done")

	// Write to file
	fmt.Print("Saving compose file ...")
	output, err1 := yaml.Marshal(&composeFile)
	err2 := ioutil.WriteFile("./docker-compose.yml", output, 0777)
	if err1 != nil || err2 != nil {
		utils.Error("Could not write yaml to compose file.", true)
	}
	color.Green(" done")
}
