package commands

import (
	"compose-generator/model"
	"compose-generator/utils"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	yaml "gopkg.in/yaml.v3"
)

func Remove(flag_run bool, flag_demonized bool) {
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

	// Ask for service
	var items []string
	for k := range compose_file.Services {
		items = append(items, k)
	}
	_, service_name := utils.MenuQuestion("Which service do you want to remove?", items)
	fmt.Println()

	// Remove service
	fmt.Print("Removing service ...")
	delete(compose_file.Services, service_name) // Remove service itself
	for k, s := range compose_file.Services {
		s.DependsOn = utils.RemoveStringFromSlice(s.DependsOn, service_name) // Remove dependencies on service
		s.Links = utils.RemoveStringFromSlice(s.Links, service_name)         // Remove links on service
		compose_file.Services[k] = s
	}
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
