package pass

import (
	"compose-generator/model"
	"compose-generator/project"
	"compose-generator/util"
	"path/filepath"
	"strconv"

	"github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate transforms the selected templates list and the enriched project to a composition
func Generate(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	util.Pel()
	templateCount := selectedTemplates.GetTotal()
	if templateCount > 0 {
		util.P("Generating configuration from " + strconv.Itoa(templateCount) + " templates ... ")
		// Prepare
		project.Composition = &types.Project{
			WorkingDir: "./",
			Services:   types.Services{},
		}
		if project.WithReadme {
			instructionsHeaderPath := util.GetPredefinedServicesPath() + "/INSTRUCTIONS_HEADER.md"
			project.ReadmeChildPaths = append(project.ReadmeChildPaths, instructionsHeaderPath)
		}

		// Generate services from selected templates
		// Generate frontends
		for _, template := range selectedTemplates.FrontendServices {
			generateService(project, selectedTemplates, template, model.TemplateTypeFrontend, template.Name)
		}
		// Generate backends
		for _, template := range selectedTemplates.BackendServices {
			generateService(project, selectedTemplates, template, model.TemplateTypeBackend, template.Name)
		}
		// Generate databases
		for _, template := range selectedTemplates.DatabaseServices {
			generateService(project, selectedTemplates, template, model.TemplateTypeDatabase, template.Name)
		}
		// Generate db admins
		for _, template := range selectedTemplates.DbAdminServices {
			generateService(project, selectedTemplates, template, model.TemplateTypeDbAdmin, template.Name)
		}
		// Generate proxies
		for _, template := range selectedTemplates.ProxyService {
			generateService(project, selectedTemplates, template, model.TemplateTypeProxy, template.Name)
		}
		// Generate tls helpers
		for _, template := range selectedTemplates.TlsHelperService {
			generateService(project, selectedTemplates, template, model.TemplateTypeTlsHelper, template.Name)
		}
		util.Done()
	} else {
		util.Error("No templates selected. Aborting ...", nil, true)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func generateService(
	proj *model.CGProject,
	selectedTemplates *model.SelectedTemplates,
	template model.PredefinedTemplateConfig,
	templateType string,
	serviceName string,
) {
	// Load service configuration
	service := project.LoadTemplateService(
		proj,
		selectedTemplates,
		templateType,
		serviceName,
		project.LoadFromDir(template.Dir),
		project.LoadFromComposeFile("service.yml"),
	)
	// Change to build context path to contain more information
	if service.Build != nil && service.Build.Context != "" {
		service.Build.Context = template.Dir + "/" + service.Build.Context
	}
	// Add service to the project
	proj.Composition.Services = append(proj.Composition.Services, *service)
	// Add child readme files
	for _, readmePath := range template.GetFilePathsByType("docs") {
		proj.ReadmeChildPaths = append(proj.ReadmeChildPaths, filepath.Join(template.Dir, readmePath))
	}
	// Add gitignore patterns
	for _, envFilePath := range template.GetFilePathsByType("env") {
		if !util.SliceContainsString(proj.GitignorePatterns, envFilePath) {
			proj.GitignorePatterns = append(proj.GitignorePatterns, envFilePath)
		}
	}
}
