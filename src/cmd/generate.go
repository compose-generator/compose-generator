package cmd

import (
	"compose-generator/model"
	"compose-generator/parser"
	"compose-generator/pass"
	"compose-generator/project"
	"compose-generator/util"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate a docker compose configuration
func Generate(
	configPath string,
	flagAdvanced bool,
	flagRun bool,
	flagDetached bool,
	flagForce bool,
	flagWithInstructions bool,
) {
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
		Secrets:     make(map[string]string),
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

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		util.Pl("Running Docker Compose now ...")
		util.Pel()
		util.DockerComposeUp(flagDetached)
		return
	}
	// Print success message
	util.Pel()
	util.SuccessMessage("🎉 Done! You now can execute \"$ docker-compose up\" to launch your app! 🎉")
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
		DbAdminService:   []model.PredefinedTemplateConfig{},
		ProxyServices:    []model.PredefinedTemplateConfig{},
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
	pass.Generate(project, config, selectedTemplates)
	pass.GenerateSecrets(project, config)
	pass.GenerateExecServiceInitCommands(project, config)
	pass.GenerateExecDemoAppInitCommands(project, config)
}
