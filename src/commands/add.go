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

func Add(flag_advanced bool, flag_run bool, flag_demonized bool) {
	// Ask for custom YAML file
	path := utils.TextQuestionWithDefault("Which compose file do you want to add the service to?", "./docker-compose.yml")

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
	if build {
		// Ask for build path
		build_path = utils.TextQuestionWithDefault("Where is your Dockerfile located?", ".")
		// Check if Dockerfile exists
		if !utils.FileExists(build_path+"/Dockerfile") && !utils.FileExists(build_path+"Dockerfile") {
			utils.Error("Aborting. The Dockerfile cannot be found.", true)
		}
	}

	// Ask for registry
	registry := utils.TextQuestionWithDefault("From which registry do you want to pick?", "docker.io")

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
	fmt.Println(service_name)
	fmt.Println(registry)

	// Add service
	fmt.Print("Adding service ...")
	service := model.Service{
		Build:         build_path,
		Image:         image,
		ContainerName: service_name,
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
