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

func Add(flag_advanced bool, flag_run bool, flag_demonized bool, flag_force bool) {
	// Ask for custom YAML file
	path := "./docker-compose.yml"
	if flag_advanced {
		path = utils.TextQuestionWithDefault("Which compose file do you want to add the service to?", "./docker-compose.yml")
	}

	fmt.Print("Parsing compose file ...")
	// Load compose file
	jsonFile, err := os.Open(path)
	if err != nil {
		utils.Error("Internal error - unable to load compose file", true)
	}
	bytes, _ := ioutil.ReadAll(jsonFile)

	// Parse YAML
	compose_file := model.ComposeFile{}
	err = yaml.Unmarshal(bytes, &compose_file)
	color.Green(" done")
	fmt.Println()

	// Ask if the image should be built from source
	build := utils.YesNoQuestion("Build from source?", false)
	var build_path string
	registry := ""
	if build {
		// Ask for build path
		build_path = utils.TextQuestionWithDefault("Where is your Dockerfile located?", ".")
		// Check if Dockerfile exists
		if !utils.FileExists(build_path+"/Dockerfile") && !utils.FileExists(build_path+"Dockerfile") {
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

	default_service_name := image
	i := strings.Index(default_service_name, "/")
	if i != -1 {
		default_service_name = default_service_name[i+1:]
	}
	i = strings.Index(default_service_name, ":")
	if i != -1 {
		default_service_name = default_service_name[:i]
	}

	// Ask for service name
	service_name := utils.TextQuestionWithDefault("How do you want to call your service (best practise: lower cased):", default_service_name)
	if _, ok := compose_file.Services[service_name]; ok {
		// Service name already existing
		overwrite_service := utils.YesNoQuestion("This service name alreay exists in the compose file. It will be overwritten if you continue. Continue?", false)
		if !overwrite_service {
			os.Exit(0)
		}
	}

	// Ask for container name
	container_name := service_name
	if flag_advanced {
		container_name = utils.TextQuestionWithDefault("How do you want to call your container (best practise: lower cased):", service_name)
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
	var restart_value string = ""
	if flag_advanced {
		items := []string{"always", "on-failure", "unless-stopped", "no"}
		_, restart_value = utils.MenuQuestion("When should the service get restarted?", items)
		fmt.Println()

	}

	// Add service
	fmt.Print("Adding service ...")
	service := model.Service{
		Build:         build_path,
		Image:         registry + image,
		ContainerName: container_name,
		Restart:       restart_value,
	}
	compose_file.Services[service_name] = service
	color.Green(" done")

	// Write to file
	fmt.Print("Saving compose file ...")
	output, err1 := yaml.Marshal(&compose_file)
	err2 := ioutil.WriteFile(path, output, 0777)
	if err1 != nil || err2 != nil {
		utils.Error("Could not write yaml to compose file.", true)
	}
	color.Green(" done")

	// Run if the corresponding flag is set
	if flag_run || flag_demonized {
		utils.DockerComposeUp(flag_demonized)
	}
}
