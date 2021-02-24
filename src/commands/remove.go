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

func Remove(flag_run bool, flag_demonized bool, flag_with_volumes bool, flag_force bool, flag_advanced bool) {
	// Ask for custom YAML file
	path := "./docker-compose.yml"
	if flag_advanced {
		path = utils.TextQuestionWithDefault("From which compose file do you want to remove a service?", "./docker-compose.yml")
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

	// Ask for service
	var items []string
	for k := range compose_file.Services {
		items = append(items, k)
	}
	_, service_name := utils.MenuQuestion("Which service do you want to remove?", items)
	fmt.Println()

	// Remove volumes
	if flag_with_volumes {
		really_delete_volumes := true
		if !flag_force {
			really_delete_volumes = utils.YesNoQuestion("Do you really want to delete all attached volumes. All data will be lost.", false)
		}
		if really_delete_volumes {
			fmt.Print("Removing volumes ...")
			volumes := compose_file.Services[service_name].Volumes
			for _, paths := range volumes {
				path := paths
				if strings.Contains(path, ":") {
					path = path[:strings.IndexByte(path, ':')]
				}
				// Check if volume is used by another container
				can_be_deleted := true
			out:
				for k, s := range compose_file.Services {
					if k != service_name {
						for _, paths_inner := range s.Volumes {
							path_inner := paths_inner
							if strings.Contains(path_inner, ":") {
								path_inner = path_inner[:strings.IndexByte(path_inner, ':')]
							}
							if path_inner == path {
								can_be_deleted = false
								break out
							}
						}
					}
				}
				if can_be_deleted && utils.FileExists(path) {
					os.RemoveAll(path)
				}
			}
			color.Green(" done")
		}
	}

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
