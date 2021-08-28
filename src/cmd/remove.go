package cmd

import (
	"compose-generator/model"
	"compose-generator/pass"
	"compose-generator/project"
	"compose-generator/util"
	"errors"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Remove services from an existing compose file
func Remove(
	serviceNames []string,
	flagRun bool,
	flagDetached bool,
	flagWithVolumes bool,
	flagForce bool,
	flagAdvanced bool,
) {
	// Clear the screen for CG output
	util.ClearScreen()

	// Ask for custom compose file
	composeFilePath := "docker-compose.yml"
	if flagAdvanced {
		composeFilePath = util.TextQuestionWithDefault("From with compose file do you want to load?", "./docker-compose.yml")
	}

	// Load project
	util.P("Loading project ... ")
	options := project.LoadOptions{ComposeFileName: composeFilePath}
	proj := project.LoadProject(options)
	proj.AdvancedConfig = flagAdvanced
	proj.ForceConfig = flagForce
	proj.WithVolumesConfig = flagWithVolumes
	util.Done()
	util.Pel()

	// Ask for services to remove
	if len(serviceNames) == 0 {
		serviceNames = proj.Project.ServiceNames()
		if len(serviceNames) == 0 {
			util.Error("No services found", nil, true)
		}
		serviceNames = util.MultiSelectMenuQuestion("Which services do you want to remove?", serviceNames)
	}

	// Remove selected services
	for _, serviceName := range serviceNames {
		removeService(proj, serviceName, flagWithVolumes)
	}

	// Save project
	util.P("Saving project ... ")
	project.SaveProject(proj)
	util.Done()
	util.Pel()

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func removeService(project *model.CGProject, serviceName string, withVolumes bool) {
	// Get service and its index by its name
	service, index, err := findServiceWithIndex(project, serviceName)
	if err != nil {
		util.Error("Service not found", err, false)
		return
	}

	// Display warning to the user
	if !project.ForceConfig {
		if !util.YesNoQuestion("Do you really want to remove service '"+serviceName+"'?", false) {
			return
		}
	}

	// Execute passes on the service
	pass.RemoveVolumes(&service, project)
	pass.RemoveNetworks(&service, project)
	pass.RemoveDependencies(&service, project)

	// Remove service from the project
	/*project.Project.Services = */
	removeServiceFromProject(project.Project.Services, index)
}

// ---------------------------------------------------------------- Helper functions ---------------------------------------------------------------

func findServiceWithIndex(project *model.CGProject, serviceName string) (spec.ServiceConfig, int, error) {
	for index, service := range project.Project.Services {
		if service.Name == serviceName {
			return service, index, nil
		}
	}
	return spec.ServiceConfig{}, -1, errors.New("Service '" + serviceName + "' not found")
}

func removeServiceFromProject(services []spec.ServiceConfig, index int) []spec.ServiceConfig {
	services[index] = services[len(services)-1]
	return services[:len(services)-1]
}

/*func removeFromFile(filePath string, serviceNames []string, flagWithVolumes bool, flagForce bool, flagAdvanced bool) {
	util.P("Parsing compose file ... ")
	// Load compose file
	composeFile, err := dcu.DeserializeFromFile(filePath)
	if err != nil {
		util.Error("Internal error - unable to parse compose file", err, true)
	}
	util.Done()

	// Ask for service(s)
	if len(serviceNames) == 0 {
		var items []string
		for k := range composeFile.Services {
			items = append(items, k)
		}
		serviceNames = util.MultiSelectMenuQuestion("Which services do you want to remove?", items)
		util.Pel()
	}

	for _, serviceName := range serviceNames {
		if !flagForce {
			reallyRemove := util.YesNoQuestion("Do you really want to remove service '"+serviceName+"'?", false)
			if !reallyRemove {
				continue;
			}
		}

		// Remove volumes
		if flagWithVolumes {
			removeVolumesForService(composeFile, serviceName, flagForce)
		}

		// Remove service
		util.P("Removing service '" + serviceName + "' ... ")
		var networkCount = make(map[string]int32)
		delete(composeFile.Services, serviceName) // Remove service itself
		for k, s := range composeFile.Services {
			s.DependsOn = util.RemoveStringFromSlice(s.DependsOn, serviceName) // Remove dependencies on service
			s.Links = util.RemoveStringFromSlice(s.Links, serviceName)         // Remove links on service
			for networkName := range composeFile.Networks {                    // Collect count of every network
				if util.SliceContainsString(s.Networks, networkName) {
					networkCount[networkName]++
				}
			}
			composeFile.Services[k] = s
		}

		// Remove unused networks
		for networkName := range composeFile.Networks {
			if networkCount[networkName] < 2 {
				delete(composeFile.Networks, networkName) // Delete network itself
				for k, s := range composeFile.Services {  // Delete references on service
					s.Networks = util.RemoveStringFromSlice(s.Networks, networkName)
					composeFile.Services[k] = s
				}
			}
		}
		util.Done()
	}

	// Write to file
	util.P("Saving compose file ... ")
	if err := dcu.SerializeToFile(composeFile, "./docker-compose.yml"); err != nil {
		util.Error("Could not write yaml to compose file", err, true)
	}
	util.Done()
}

func removeVolumesForService(composeFile dcu_model.ComposeFile, serviceName string, flagForce bool) {
	reallyDeleteVolumes := true
	if !flagForce {
		reallyDeleteVolumes = util.YesNoQuestion("Do you really want to delete all attached volumes of service '"+serviceName+"'? All data will be lost.", false)
	}
	if reallyDeleteVolumes {
		util.P("Removing volumes of '" + serviceName + "' ... ")
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
			if canBeDeleted && util.FileExists(path) {
				os.RemoveAll(path)
			}
		}
		util.Done()
	}
}*/
