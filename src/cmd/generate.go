/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/model"
	genPass "compose-generator/pass/generate"
	"compose-generator/project"
	"compose-generator/util"
	"time"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/urfave/cli/v2"
)

// GenerateCliFlags are the cli flags for the generate command
var GenerateCliFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "advanced",
		Aliases: []string{"a"},
		Usage:   "Generate compose file in advanced mode",
		Value:   false,
	},
	&cli.PathFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "Pass a configuration as a `FILE` with predefined answers. Works good for CI",
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
	&cli.BoolFlag{
		Name:    "with-instructions",
		Aliases: []string{"i"},
		Usage:   "Generates a README.md file with instructions to use the template",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run docker compose after creating the compose file",
		Value:   false,
	},
}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate a docker compose configuration
func Generate(c *cli.Context) error {
	infoLogger.Println("Generate command executed")

	// Extract flags
	configPath := c.Path("config")
	flagAdvanced := c.Bool("advanced")
	flagRun := c.Bool("run")
	flagDetached := c.Bool("detached")
	flagForce := c.Bool("force")
	flagWithInstructions := c.Bool("with-instructions")

	// Check if CCom is installed and Docker is running
	util.EnsureCComIsInstalled()
	util.EnsureDockerIsRunning()

	// Clear screen if in interactive mode
	if configPath == "" {
		clearScreen()
	}

	// Check for predefined service templates updates
	util.CheckForServiceTemplateUpdate()

	// Create instances of project and generate config
	proj := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: flagAdvanced,
			WithGitignore:  true,
			WithReadme:     flagWithInstructions,
			CreatedBy:      util.GetUsername(),
			CreatedAt:      time.Now().UnixNano(),
		},
		Composition: &spec.Project{
			WorkingDir: "./",
			Services:   spec.Services{},
		},
		ForceConfig: flagForce,
		Vars:        make(model.Vars),
		ProxyVars:   make(map[string]model.Vars),
		Secrets:     []model.ProjectSecret{},
	}
	config := &model.GenerateConfig{}

	// Run passes
	genPass.LoadGenerateConfig(proj, config, configPath)

	// Enrich project with information
	EnrichProjectWithServices(proj, config)

	// Save project
	infoLogger.Println("Saving project ...")
	spinner := startProcess("Saving project ...")
	project.SaveProject(proj)
	stopProcess(spinner)
	infoLogger.Println("Saving project done")

	// Print generated secrets
	genPass.GeneratePrintSecrets(proj)

	// Run if the corresponding flag is set. Otherwise, print success message
	if flagRun || flagDetached {
		infoLogger.Println("Running Docker Compose ...")
		util.DockerComposeUp(flagDetached)
		infoLogger.Println("Running Docker Compose done")
	} else {
		pel()
		printSuccess("🎉 Done! You now can execute \"$ docker compose up\" to launch your app! 🎉")
		pel()
	}
	return nil
}

// EnrichProjectWithServices enriches a project with a custom selection of predefined services
func EnrichProjectWithServices(project *model.CGProject, config *model.GenerateConfig) {
	// Clear screen
	if config == nil || !config.FromFile {
		clearScreen()
	}

	// Parse available service templates
	spinner := startProcess("Loading predefined service templates ...")
	availableTemplates := getAvailablePredefinedTemplates()
	stopProcess(spinner)

	// Generate composition
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{},
		BackendServices:  []model.PredefinedTemplateConfig{},
		DatabaseServices: []model.PredefinedTemplateConfig{},
		DbAdminServices:  []model.PredefinedTemplateConfig{},
		ProxyService:     []model.PredefinedTemplateConfig{},
		TlsHelperService: []model.PredefinedTemplateConfig{},
	}
	if project.ProductionReady {
		generateChooseProxiesPass(project, availableTemplates, selectedTemplates, config)
		generateChooseTlsHelpersPass(project, availableTemplates, selectedTemplates, config)
	}
	generateChooseFrontendsPass(project, availableTemplates, selectedTemplates, config)
	generateChooseBackendsPass(project, availableTemplates, selectedTemplates, config)
	generateChooseDatabasesPass(project, availableTemplates, selectedTemplates, config)
	generateChooseDbAdminsPass(project, availableTemplates, selectedTemplates, config)

	// Execute passes
	generatePass(project, selectedTemplates)
	generateResolveDependencyGroupsPass(project, selectedTemplates)
	generateSecretsPass(project, selectedTemplates)
	generateAddProfilesPass(project)
	generateAddProxyNetworks(project, selectedTemplates)
	generateCopyVolumesPass(project)
	generateReplaceVarsInConfigFilesPass(project, selectedTemplates)
	generateExecServiceInitCommandsPass(project, selectedTemplates)
	generateExecDemoAppInitCommandsPass(project, selectedTemplates)
}
