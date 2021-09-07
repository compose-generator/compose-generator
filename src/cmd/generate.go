package cmd

import (
	"compose-generator/model"
	"compose-generator/parser"
	"compose-generator/pass"
	"compose-generator/project"
	"compose-generator/util"

	"github.com/urfave/cli/v2"
)

// Cli flags for the generate command
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
		Usage:   "Run docker-compose detached after creating the compose file",
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
		Usage:   "Run docker-compose after creating the compose file",
		Value:   false,
	},
}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate a docker compose configuration
func Generate(c *cli.Context) error {
	// Extract flags
	configPath := c.Path("config")
	flagAdvanced := c.Bool("advanced")
	flagRun := c.Bool("run")
	flagDetached := c.Bool("detached")
	flagForce := c.Bool("force")
	flagWithInstructions := c.Bool("with-instructions")

	// Check if CCom is installed
	util.EnsureCComIsInstalled()

	// Clear screen if in interactive mode
	if configPath == "" {
		util.ClearScreen()
	}

	// Check for predefined service templates updates
	util.CheckForServiceTemplateUpdate()

	// Create instances of project and generate config
	proj := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: flagAdvanced,
			WithGitignore:  true,
			WithReadme:     flagWithInstructions,
		},
		ForceConfig: flagForce,
		Vars:        make(map[string]string),
		Secrets:     []model.ProjectSecret{},
	}
	config := &model.GenerateConfig{}

	// Run passes
	pass.LoadGenerateConfig(proj, config, configPath)

	// Enrich project with information
	generateProject(proj, config)

	// Save project
	util.P("Saving project ... ")
	project.SaveProject(proj)
	util.Done()

	// Print generated secrets
	pass.GeneratePrintSecrets(proj)

	// Run if the corresponding flag is set. Otherwise, print success message
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached)
	} else {
		util.Pel()
		util.Success("🎉 Done! You now can execute \"$ docker-compose up\" to launch your app! 🎉")
	}
	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func generateProject(project *model.CGProject, config *model.GenerateConfig) {
	// Clear screen
	if !config.FromFile {
		util.ClearScreen()
	}

	// Parse available service templates
	util.P("Loading predefined service templates ... ")
	availableTemplates := parser.GetAvailablePredefinedTemplates()
	util.Done()

	// Generate composition
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{},
		BackendServices:  []model.PredefinedTemplateConfig{},
		DatabaseServices: []model.PredefinedTemplateConfig{},
		DbAdminServices:  []model.PredefinedTemplateConfig{},
		ProxyService:     []model.PredefinedTemplateConfig{},
		TlsHelperService: []model.PredefinedTemplateConfig{},
	}
	pass.GenerateChooseFrontends(project, availableTemplates, selectedTemplates, config)
	pass.GenerateChooseBackends(project, availableTemplates, selectedTemplates, config)
	pass.GenerateChooseDatabases(project, availableTemplates, selectedTemplates, config)
	pass.GenerateChooseDbAdmins(project, availableTemplates, selectedTemplates, config)
	if project.ProductionReady {
		pass.GenerateChooseProxies(project, availableTemplates, selectedTemplates, config)
		pass.GenerateChooseTlsHelpers(project, availableTemplates, selectedTemplates, config)
	}

	// Execute passes
	pass.Generate(project, selectedTemplates)
	pass.GenerateResolveDependencyGroups(project, selectedTemplates)
	pass.GenerateSecrets(project, selectedTemplates)
	pass.GenerateCopyVolumes(project)
	pass.GenerateExecServiceInitCommands(project, selectedTemplates)
	pass.GenerateExecDemoAppInitCommands(project, selectedTemplates)
}
