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
	proj := project.LoadProject(
		project.LoadFromComposeFile(composeFilePath),
	)
	proj.AdvancedConfig = flagAdvanced
	proj.ForceConfig = flagForce
	proj.WithVolumesConfig = flagWithVolumes
	util.Done()
	util.Pel()

	// Ask for services to remove
	if len(serviceNames) == 0 {
		serviceNames = proj.Composition.ServiceNames()
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
	project.Composition.Services = removeServiceFromProject(project.Composition.Services, index)
}

// ---------------------------------------------------------------- Helper functions ---------------------------------------------------------------

func findServiceWithIndex(project *model.CGProject, serviceName string) (spec.ServiceConfig, int, error) {
	for index, service := range project.Composition.Services {
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
