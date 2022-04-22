/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/model"
	commonPass "compose-generator/pass/common"
	"compose-generator/project"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/urfave/cli/v2"
)

// RemoveCliFlags are the cli flags for the remove command
var RemoveCliFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "advanced",
		Aliases: []string{"a"},
		Usage:   "Show questions for advanced customization",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "force",
		Aliases: []string{"f"},
		Usage:   "Skip safety checks",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run Docker Compose after creating the Compose file",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "detached",
		Aliases: []string{"d"},
		Usage:   "Run Docker Compose detached after creating the Compose file",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "with-volumes",
		Aliases: []string{"v"},
		Usage:   "Remove associated volumes",
		Value:   false,
	},
}

var removeServiceFromProjectMockable = removeServiceFromProject

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Remove services from an existing compose file
func Remove(c *cli.Context) error {
	infoLogger.Println("Remove command executed")

	// Extract flags
	serviceNames := c.Args().Slice()
	flagRun := c.Bool("run")
	flagDetached := c.Bool("detached")
	flagWithVolumes := c.Bool("with-volumes")
	flagForce := c.Bool("force")
	flagAdvanced := c.Bool("advanced")

	// Check if Docker is running
	util.EnsureDockerIsRunning()

	// Clear the screen for CG output
	util.ClearScreen()

	// Ask for custom compose file
	composeFilePath := "docker-compose.yml"
	if flagAdvanced {
		composeFilePath = util.TextQuestionWithDefault("From with compose file do you want to load?", "./docker-compose.yml")
	}

	// Load project
	spinner := util.StartProcess("Loading project ...")
	proj := project.LoadProject(
		project.LoadFromComposeFile(composeFilePath),
	)
	proj.AdvancedConfig = flagAdvanced
	proj.ForceConfig = flagForce
	proj.WithVolumesConfig = flagWithVolumes
	util.StopProcess(spinner)
	util.InfoLogger.Println("Loading project (done)")
	util.Pel()

	// Execute additional validation steps
	commonPass.CommonCheckForDependencyCycles(proj)

	// Ask for services to remove
	if len(serviceNames) == 0 {
		serviceNames = proj.Composition.ServiceNames()
		if len(serviceNames) == 0 {
			util.ErrorLogger.Println("Removal of 0 services. Therefore aborting")
			logError("No services found", true)
		}
		serviceNames = util.MultiSelectMenuQuestion("Which services do you want to remove?", serviceNames)
	}

	// Remove selected services
	for _, serviceName := range serviceNames {
		removeService(proj, serviceName, flagWithVolumes)
	}

	// Save project
	spinner = util.StartProcess("Saving project ...")
	project.SaveProject(proj)
	util.StopProcess(spinner)
	util.Pel()
	util.InfoLogger.Println("Saving project (done)")

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached, proj.ProductionReady)
	}
	util.InfoLogger.Println("Docker Compose command terminated")
	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func removeService(project *model.CGProject, serviceName string, withVolumes bool) {
	// Get service and its index by its name
	service, err := project.Composition.GetService(serviceName)
	if err != nil {
		util.WarningLogger.Println("Selected service was not found in composition: " + err.Error())
		logError("Service not found", false)
		return
	}

	// Display warning to the user
	if !project.ForceConfig {
		if !yesNoQuestion("Do you really want to remove service '"+serviceName+"'?", false) {
			return
		}
	}

	// Execute passes on the service
	removeVolumesPass(&service, project)
	removeNetworksPass(&service, project)
	removeDependenciesPass(&service, project)

	// Remove service from the project
	removeServiceFromProjectMockable(&project.Composition.Services, serviceName)
}

// ---------------------------------------------------------------- Helper functions ---------------------------------------------------------------

func removeServiceFromProject(services *spec.Services, serviceName string) {
	// Search service
	index := 0
	for i, service := range *services {
		if service.Name == serviceName {
			index = i
			break
		}
	}
	// Remove service
	(*services)[index] = (*services)[len(*services)-1]
	*services = (*services)[:len(*services)-1]
}
