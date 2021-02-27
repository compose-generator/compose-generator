package commands

import (
	"compose-generator/model"
	"compose-generator/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	yaml "gopkg.in/yaml.v3"
)

// Add: Adds a service to an existing compose file
func Add(flagAdvanced bool, flagRun bool, flagDetached bool, flagForce bool) {
	// Ask for custom YAML file
	path := "./docker-compose.yml"
	if flagAdvanced {
		path = utils.TextQuestionWithDefault("Which compose file do you want to add the service to?", "./docker-compose.yml")
	}

	fmt.Print("Parsing compose file ...")
	// Load compose file
	jsonFile, err := os.Open(path)
	if err != nil {
		utils.Error("Internal error - unable to load compose file", true)
		err = nil
	}
	bytes, _ := ioutil.ReadAll(jsonFile)

	// Parse YAML
	composeFile := model.ComposeFile{}
	if err = yaml.Unmarshal(bytes, &composeFile); err != nil {
		utils.Error("Internal error - unable to parse compose file", true)
		err = nil
	}

	serviceNames := []string{}
	for name, _ := range composeFile.Services {
		serviceNames = append(serviceNames, name)
	}
	color.Green(" done")
	fmt.Println()

	// Ask if the image should be built from source
	build := utils.YesNoQuestion("Build from source?", false)
	var buildPath string
	registry := ""
	if build {
		// Ask for build path
		buildPath = utils.TextQuestionWithDefault("Where is your Dockerfile located?", ".")
		// Check if Dockerfile exists
		if !utils.FileExists(buildPath+"/Dockerfile") && !utils.FileExists(buildPath+"Dockerfile") {
			utils.Error("Aborting. The Dockerfile cannot be found.", true)
		}
	} else {
		// Ask for registry
		registry = utils.TextQuestionWithDefault("From which registry do you want to pick?", "docker.io")
		if registry == "docker.io" {
			registry = ""
		} else {
			registry = registry + "/"
		}
	}

	// Ask for image
	image := utils.TextQuestionWithDefault("Which base image do you want to pick?", "hello-world")

	// Check image
	if !build && !flagForce {
		fmt.Print("\nChecking image ...")
		layerCount := utils.CheckDockerImage(registry + image)
		if layerCount > -1 {
			color.Green(" found (" + strconv.Itoa(layerCount) + " layer(s))\n\n")
		} else {
			color.Red(" not found or no access\n\n")
		}
	}

	defaultServiceName := image
	i := strings.Index(defaultServiceName, "/")
	if i != -1 {
		defaultServiceName = defaultServiceName[i+1:]
	}
	i = strings.Index(defaultServiceName, ":")
	if i != -1 {
		defaultServiceName = defaultServiceName[:i]
	}

	// Ask for service name
	serviceName := utils.TextQuestionWithDefault("How do you want to call your service (best practise: lower cased):", defaultServiceName)
	if _, ok := composeFile.Services[serviceName]; ok {
		// Service name already existing
		if !utils.YesNoQuestion("This service name alreay exists in the compose file. It will be overwritten if you continue. Continue?", false) {
			os.Exit(0)
		}
	}

	// Ask for container name
	containerName := serviceName
	if flagAdvanced {
		containerName = utils.TextQuestionWithDefault("How do you want to call your container (best practise: lower cased):", serviceName)
	}

	// Ask for volumes
	if utils.YesNoQuestion("Do you want to add volumes to your service?", false) {

	}

	// Ask for networks
	if utils.YesNoQuestion("Do you want to add networks to your service?", false) {

	}

	// Ask for ports
	if utils.YesNoQuestion("Do you want to expose ports of your service?", false) {

	}

	// Ask for env variables
	if utils.YesNoQuestion("Do you want to provide environment variables to your service?", false) {

	}

	// Ask for env files
	envFile := ""
	if utils.YesNoQuestion("Do you want to provide an environment file to your service?", false) {
	EnvFile:
		fmt.Println()
		envFile = utils.TextQuestionWithSuggestions("Where is your env file located?", "environment.env", func(toComplete string) (files []string) {
			files, _ = filepath.Glob(toComplete + "*")
			return
		})
		if !utils.FileExists(envFile) || utils.IsDirectory(envFile) {
			utils.Error("File is not valid. Please select another file", false)
			goto EnvFile
		}
		fmt.Println()
	}

	// Ask for depends on
	dependsServices := []string{}
	if utils.YesNoQuestion("Should your service depend on other services?", false) {
		fmt.Println()
		dependsServices = utils.MultiSelectMenuQuestion("From which services should your service depend?", serviceNames)
		fmt.Println()
	}

	// Ask for restart mode
	var restartValue string = ""
	if flagAdvanced {
		items := []string{"always", "on-failure", "unless-stopped", "no"}
		restartValue = utils.MenuQuestion("When should the service get restarted?", items)
		fmt.Println()

	}

	// Add service
	fmt.Print("Adding service ...")
	service := model.Service{
		Build:         buildPath,
		Image:         registry + image,
		ContainerName: containerName,
		Restart:       restartValue,
		DependsOn:     dependsServices,
		EnvFile:       []string{envFile},
	}
	composeFile.Services[serviceName] = service
	color.Green(" done")

	// Write to file
	fmt.Print("Saving compose file ...")
	output, err1 := yaml.Marshal(&composeFile)
	err2 := ioutil.WriteFile(path, output, 0777)
	if err1 != nil || err2 != nil {
		utils.Error("Could not write yaml to compose file.", true)
	}
	color.Green(" done")

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		utils.DockerComposeUp(flagDetached)
	}
}
