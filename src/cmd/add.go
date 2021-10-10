/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/model"
	commonPass "compose-generator/pass/common"
	"compose-generator/project"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/client"
	"github.com/urfave/cli/v2"
)

// AddCliFlags are the cli flags for the add command
var AddCliFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "advanced",
		Aliases: []string{"a"},
		Usage:   "Generate compose file in advanced mode",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run docker compose after creating the compose file",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "detached",
		Aliases: []string{"d"},
		Usage:   "Run docker compose detached after creating the compose file",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "force",
		Aliases: []string{"f"},
		Usage:   "Skip safety checks",
		Value:   false,
	},
}

var AddCustomServiceMockable = AddCustomService

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Add a service to an existing compose file
func Add(c *cli.Context) error {
	// Extract flags
	flagAdvanced := c.Bool("advanced")
	flagRun := c.Bool("run")
	flagDetached := c.Bool("detached")
	flagForce := c.Bool("force")

	// Check if CCom is installed and Docker is running
	util.EnsureCComIsInstalled()
	util.EnsureDockerIsRunning()

	// Clear the screen for CG output
	clearScreen()

	// Check for predefined service templates updates
	util.CheckForServiceTemplateUpdate()

	// Ask for custom compose file
	composeFilePath := "docker-compose.yml"
	if flagAdvanced {
		composeFilePath = textQuestionWithDefault("From which compose file do you want to load?", "./docker-compose.yml")
	}

	// Load project
	spinner := startProcess("Loading project ...")
	proj := project.LoadProject(
		project.LoadFromComposeFile(composeFilePath),
	)
	proj.AdvancedConfig = flagAdvanced
	proj.ForceConfig = flagForce
	stopProcess(spinner)
	pel()

	// Execute additional validation steps
	commonPass.CommonCheckForDependencyCycles(proj)

	// Add custom service
	AddCustomServiceMockable(proj)

	// Save project
	spinner = startProcess("Saving project ...")
	project.SaveProject(proj)
	stopProcess(spinner)
	pel()

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached)
	}

	return nil
}

// AddCustomService adds a fully customizable service to the project
func AddCustomService(project *model.CGProject) {
	newService := spec.ServiceConfig{}

	// Initialize Docker client
	client, err := newClientWithOpts(client.FromEnv)
	if err != nil {
		printError("Could not intanciate Docker client. Please check your Docker installation", err, true)
		return
	}

	// Execute passes on the service
	addBuildOrImagePass(&newService, project)
	addNamePass(&newService, project)
	addContainerNamePass(&newService, project)
	addVolumesPass(&newService, project, client)
	addNetworksPass(&newService, project, client)
	addPortsPass(&newService, project)
	addEnvVarsPass(&newService, project)
	addEnvFilesPass(&newService, project)
	addRestartPass(&newService, project)
	addDependsPass(&newService, project)
	addDependantsPass(&newService, project)

	// Add the new service to the project
	project.Composition.Services = append(project.Composition.Services, newService)
}
