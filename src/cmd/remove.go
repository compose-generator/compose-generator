package cmd

import (
	"compose-generator/utils"
	"os"
	"strings"

	dcu "github.com/compose-generator/dcu"
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

	utils.P("Parsing compose file ... ")
	// Load compose file
	composeFile, err := dcu.DeserializeFromFile(path)
	if err != nil {
		utils.Error("Internal error - unable to parse compose file", err, true)
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
				utils.P("Removing volumes of '" + serviceName + "' ... ")
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
		utils.P("Removing service '" + serviceName + "' ... ")
		delete(composeFile.Services, serviceName) // Remove service itself
		for k, s := range composeFile.Services {
			s.DependsOn = utils.RemoveStringFromSlice(s.DependsOn, serviceName) // Remove dependencies on service
			s.Links = utils.RemoveStringFromSlice(s.Links, serviceName)         // Remove links on service
			composeFile.Services[k] = s
		}
		utils.Done()
	}

	// Write to file
	utils.P("Saving compose file ... ")
	if err := dcu.SerializeToFile(composeFile, "./docker-compose.yml"); err != nil {
		utils.Error("Could not write yaml to compose file", err, true)
	}
	utils.Done()

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		utils.DockerComposeUp(flagDetached)
	}
}
