package commands

import (
	"compose-generator/model"
	"compose-generator/utils"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Remove services from an existing compose file
func Remove(serviceNames []string, flagRun bool, flagDetached bool, flagWithVolumes bool, flagForce bool, flagAdvanced bool) {
	utils.ClearScreen()

	// Ask for custom YAML file
	path := "./docker-compose.yml"
	if flagAdvanced {
		path = utils.TextQuestionWithDefault("From which compose file do you want to remove a service?", "./docker-compose.yml")
	}

	fmt.Print("Parsing compose file ... ")
	// Load compose file
	jsonFile, err := os.Open(path)
	if err != nil {
		utils.Error("Internal error - unable to load compose file", true)
	}
	bytes, _ := ioutil.ReadAll(jsonFile)

	// Parse YAML
	composeFile := model.ComposeFile{}
	if err = yaml.Unmarshal(bytes, &composeFile); err != nil {
		utils.Error("Internal error - unable to parse compose file", true)
	}
	utils.Done()

	// Ask for service
	if len(serviceNames) == 0 {
		var items []string
		for k := range composeFile.Services {
			items = append(items, k)
		}
		serviceNames = utils.MultiSelectMenuQuestion("Which services do you want to remove?", items)
		utils.Pel()
	}

	for _, serviceName := range serviceNames {
		// Remove volumes
		if flagWithVolumes {
			reallyDeleteVolumes := true
			if !flagForce {
				reallyDeleteVolumes = utils.YesNoQuestion("Do you really want to delete all attached volumes. All data will be lost.", false)
			}
			if reallyDeleteVolumes {
				fmt.Print("Removing volumes of '" + serviceName + "' ... ")
				volumes := composeFile.Services[serviceName].Volumes
				for _, paths := range volumes {
					path := paths
					if strings.Contains(path, ":") {
						path = path[:strings.IndexByte(path, ':')]
					}
					// Check if volume is used by another container
					canBeDeleted := true
				out:
					for k, s := range composeFile.Services {
						if k != serviceName {
							for _, pathsInner := range s.Volumes {
								pathInner := pathsInner
								if strings.Contains(pathInner, ":") {
									pathInner = pathInner[:strings.IndexByte(pathInner, ':')]
								}
								if pathInner == path {
									canBeDeleted = false
									break out
								}
							}
						}
					}
					if canBeDeleted && utils.FileExists(path) {
						os.RemoveAll(path)
					}
				}
				utils.Done()
			}
		}

		// Remove service
		fmt.Print("Removing service '" + serviceName + "' ... ")
		delete(composeFile.Services, serviceName) // Remove service itself
		for k, s := range composeFile.Services {
			s.DependsOn = utils.RemoveStringFromSlice(s.DependsOn, serviceName) // Remove dependencies on service
			s.Links = utils.RemoveStringFromSlice(s.Links, serviceName)         // Remove links on service
			composeFile.Services[k] = s
		}
		utils.Done()
	}

	// Write to file
	fmt.Print("Saving compose file ... ")
	output, err1 := yaml.Marshal(&composeFile)
	err2 := ioutil.WriteFile(path, output, 0777)
	if err1 != nil || err2 != nil {
		utils.Error("Could not write yaml to compose file.", true)
	}
	utils.Done()

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		utils.DockerComposeUp(flagDetached)
	}
}
