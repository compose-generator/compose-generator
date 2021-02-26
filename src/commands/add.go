package commands

import (
	"compose-generator/model"
	"compose-generator/utils"
	"fmt"
	"io/ioutil"
	"os"
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

	// Ask for env file
	if utils.YesNoQuestion("Do you want to provide an environment file to your service?", false) {

	}

	// Ask for depends on
	if utils.YesNoQuestion("Should your service depend on other services?", false) {

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
