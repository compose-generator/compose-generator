package cmd

import (
	"compose-generator/model"
	pass "compose-generator/pass/add"
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
		Usage:   "Run docker-compose after creating the compose file",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "detached",
		Aliases: []string{"d"},
		Usage:   "Run docker-compose detached after creating the compose file",
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
	// Extract flags
	flagAdvanced := c.Bool("advanced")
	flagRun := c.Bool("run")
	flagDetached := c.Bool("detached")
	flagForce := c.Bool("force")

	// Check if CCom is installed
	util.EnsureCComIsInstalled()

	// Clear the screen for CG output
	util.ClearScreen()

	// Check for predefined service templates updates
	util.CheckForServiceTemplateUpdate()

	// Ask for custom compose file
	composeFilePath := "docker-compose.yml"
	if flagAdvanced {
		composeFilePath = util.TextQuestionWithDefault("From which compose file do you want to load?", "./docker-compose.yml")
	}

	// Load project
	spinner := util.StartProcess("Loading project ...")
	proj := project.LoadProject(
		project.LoadFromComposeFile(composeFilePath),
	)
	proj.AdvancedConfig = flagAdvanced
	proj.ForceConfig = flagForce
	util.StopProcess(spinner)
	util.Pel()

	// Add custom service
	AddCustomService(proj)

	// Save project
	spinner = util.StartProcess("Saving project ...")
	project.SaveProject(proj)
	util.StopProcess(spinner)
	util.Pel()

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
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		util.Error("Could not intanciate Docker client. Please check your Docker installation", err, true)
	}

	// Execute passes on the service
	pass.AddBuildOrImage(&newService, project)
	pass.AddName(&newService, project)
	pass.AddContainerName(&newService, project)
	pass.AddVolumes(&newService, project, client)
	pass.AddNetworks(&newService, project, client)
	pass.AddPorts(&newService, project)
	pass.AddEnvVars(&newService, project)
	pass.AddEnvFiles(&newService, project)
	pass.AddRestart(&newService, project)
	pass.AddDepends(&newService, project)
	pass.AddDependants(&newService, project)

	// Add the new service to the project
	project.Composition.Services = append(project.Composition.Services, newService)
}
