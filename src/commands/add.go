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
	path := utils.TextQuestionWithDefault("From which compose file do you want to remove the service?", "./docker-compose.yml")

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

	service_name := utils.TextQuestionWithDefault("How do you want to call your service (best practise: lower cased):", default_service_name)
	fmt.Println(service_name)

	// Run if the corresponding flag is set
	if flag_run || flag_demonized {
		utils.DockerComposeUp(flag_demonized)
	}
}
