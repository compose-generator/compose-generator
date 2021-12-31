/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	commonPass "compose-generator/pass/common"
	"compose-generator/project"
	"compose-generator/util"

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

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Add a service to an existing compose file
func Add(c *cli.Context) error {
	infoLogger.Println("Add command executed")

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

	// Enrich project
	EnrichProjectWithServices(proj, nil)

	// Save project
	spinner = startProcess("Saving project ...")
	project.SaveProject(proj)
	stopProcess(spinner)
	pel()

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached, proj.ProductionReady)
	}

	return nil
}
