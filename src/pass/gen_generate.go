package pass

import (
	"compose-generator/model"
	"fmt"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Generate transforms the selected templates list and the enriched project to a composition
func Generate(project *model.CGProject, config *model.GenerateConfig, selectedTemplates *model.SelectedTemplates) {
	// Generate frontends
	for _, template := range selectedTemplates.FrontendServices {
		generateService(project, config, selectedTemplates, template)
	}
	// Generate backends
	for _, template := range selectedTemplates.BackendServices {
		generateService(project, config, selectedTemplates, template)
	}
	// Generate databases
	for _, template := range selectedTemplates.DatabaseServices {
		generateService(project, config, selectedTemplates, template)
	}
	// Generate db admins
	for _, template := range selectedTemplates.DbAdminService {
		generateService(project, config, selectedTemplates, template)
	}
	// Generate proxies
	for _, template := range selectedTemplates.ProxyServices {
		generateService(project, config, selectedTemplates, template)
	}
	// Generate tls helpers
	for _, template := range selectedTemplates.TlsHelperService {
		generateService(project, config, selectedTemplates, template)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func generateService(
	proj *model.CGProject,
	config *model.GenerateConfig,
	selectedTemplates *model.SelectedTemplates,
	template model.PredefinedTemplateConfig,
) {
	fmt.Println(template.Dir)
	/*service := project.LoadTemplateService(
		project.LoadFromDir(template.Dir),
	)*/
}
