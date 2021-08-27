package cmd

import (
	"compose-generator/model"
	"compose-generator/pass"
	"compose-generator/project"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/client"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Add a service to an existing compose file
func Add(
	flagAdvanced bool,
	flagRun bool,
	flagDetached bool,
	flagForce bool,
) {
	// Check if CCom is installed
	util.CheckIfCComIsInstalled()

	// Clear the screen for CG output
	util.ClearScreen()

	// Check for predefined service templates updates
	util.CheckForServiceTemplateUpdate()

	// Ask for custom YAML file
	composeFilePath := "docker-compose.yml"
	if flagAdvanced {
		composeFilePath = util.TextQuestionWithDefault("Which compose file do you want to add the service to?", "./docker-compose.yml")
	}

	// Load project
	util.P("Loading project ... ")
	options := project.LoadOptions{ComposeFileName: composeFilePath}
	proj := project.LoadProject(options)
	proj.AdvancedConfig = flagAdvanced
	util.Done()
	util.Pel()

	// Add custom service
	AddCustomService(proj)

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

// AddCustomService adds a fully customizable service to the project
func AddCustomService(project *model.CGProject) {
	newService := spec.ServiceConfig{}

	// Initialize Docker client
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		util.Error("Docker is not installed on your system", err, true)
	}

	// Ask questions
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
	project.Project.Services = append(project.Project.Services, newService)
}
