package pass

import (
	"compose-generator/model"
	"compose-generator/project"
	"path/filepath"
	"strconv"

	"github.com/compose-spec/compose-go/types"
)

var generateServiceMockable = generateService

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate transforms the selected templates list and the enriched project to a composition
func Generate(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	pel()
	templateCount := selectedTemplates.GetTotal()
	if templateCount > 0 {
		spinner := startProcess("Generating configuration from " + strconv.Itoa(templateCount) + " template(s) ...")
		// Prepare
		project.Composition = &types.Project{
			WorkingDir: "./",
			Services:   types.Services{},
		}
		if project.WithReadme {
			instructionsHeaderPath := getPredefinedServicesPath() + "/INSTRUCTIONS_HEADER.md"
			project.ReadmeChildPaths = append(project.ReadmeChildPaths, instructionsHeaderPath)
		}

		// Generate services from selected templates
		// Generate frontends
		for _, template := range selectedTemplates.FrontendServices {
			generateServiceMockable(project, selectedTemplates, template, model.TemplateTypeFrontend, template.Name)
		}
		// Generate backends
		for _, template := range selectedTemplates.BackendServices {
			generateServiceMockable(project, selectedTemplates, template, model.TemplateTypeBackend, template.Name)
		}
		// Generate databases
		for _, template := range selectedTemplates.DatabaseServices {
			generateServiceMockable(project, selectedTemplates, template, model.TemplateTypeDatabase, template.Name)
		}
		// Generate db admins
		for _, template := range selectedTemplates.DbAdminServices {
			generateServiceMockable(project, selectedTemplates, template, model.TemplateTypeDbAdmin, template.Name)
		}
		// Generate proxies
		for _, template := range selectedTemplates.ProxyService {
			generateServiceMockable(project, selectedTemplates, template, model.TemplateTypeProxy, template.Name)
		}
		// Generate tls helpers
		for _, template := range selectedTemplates.TlsHelperService {
			generateServiceMockable(project, selectedTemplates, template, model.TemplateTypeTlsHelper, template.Name)
		}
		stopProcess(spinner)
	} else {
		printError("No templates selected. Aborting ...", nil, true)
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
	service := loadTemplateService(
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
		if !sliceContainsString(proj.GitignorePatterns, envFilePath) {
			proj.GitignorePatterns = append(proj.GitignorePatterns, envFilePath)
		}
	}
}
