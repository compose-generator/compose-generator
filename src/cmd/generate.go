package cmd

import (
	"compose-generator/model"
	genPass "compose-generator/pass/generate"
	"compose-generator/project"
	"compose-generator/util"
	"time"

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
		ForceConfig: flagForce,
		Vars:        make(map[string]string),
		Secrets:     []model.ProjectSecret{},
	}
	config := &model.GenerateConfig{}

	// Run passes
	genPass.LoadGenerateConfig(proj, config, configPath)

	// Enrich project with information
	generateProject(proj, config)

	// Save project
	spinner := startProcess("Saving project ...")
	project.SaveProject(proj)
	stopProcess(spinner)

	// Print generated secrets
	genPass.GeneratePrintSecrets(proj)

	// Run if the corresponding flag is set. Otherwise, print success message
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached)
	} else {
		pel()
		printSuccess("🎉 Done! You now can execute \"$ docker-compose up\" to launch your app! 🎉")
		pel()
	}
	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func generateProject(project *model.CGProject, config *model.GenerateConfig) {
	// Clear screen
	if !config.FromFile {
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
	generateChooseFrontendsPass(project, availableTemplates, selectedTemplates, config)
	generateChooseBackendsPass(project, availableTemplates, selectedTemplates, config)
	generateChooseDatabasesPass(project, availableTemplates, selectedTemplates, config)
	generateChooseDbAdminsPass(project, availableTemplates, selectedTemplates, config)
	if project.ProductionReady {
		generateChooseProxiesPass(project, availableTemplates, selectedTemplates, config)
		generateChooseTlsHelpersPass(project, availableTemplates, selectedTemplates, config)
	}

	// Execute passes
	generatePass(project, selectedTemplates)
	generateResolveDependencyGroupsPass(project, selectedTemplates)
	generateSecretsPass(project, selectedTemplates)
	genAddProfilesPass(project)
	generateCopyVolumesPass(project)
	generateReplaceVarsInConfigFilesPass(project, selectedTemplates)
	generateExecServiceInitCommandsPass(project, selectedTemplates)
	generateExecDemoAppInitCommandsPass(project, selectedTemplates)
}
