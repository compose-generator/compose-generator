package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

var replaceGroupDependencyMockable = replaceGroupDependency

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateResolveDependencyGroups resolves group dependencies like 'database' or 'frontend' to concrete service dependencies
func GenerateResolveDependencyGroups(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	spinner := startProcess("Resolving group dependencies ...")
	for i := 0; i < len(project.Composition.Services); i++ {
		service := &project.Composition.Services[i]
		// Search for frontend group dependency
		if _, ok := service.DependsOn[model.TemplateTypeFrontend]; ok {
			replaceGroupDependencyMockable(service, selectedTemplates.FrontendServices, model.TemplateTypeFrontend)
		}
		// Search for backend group dependency
		if _, ok := service.DependsOn[model.TemplateTypeBackend]; ok {
			replaceGroupDependencyMockable(service, selectedTemplates.BackendServices, model.TemplateTypeBackend)
		}
		// Search for database group dependency
		if _, ok := service.DependsOn[model.TemplateTypeDatabase]; ok {
			replaceGroupDependencyMockable(service, selectedTemplates.DatabaseServices, model.TemplateTypeDatabase)
		}
		// Search for dbadmin group dependency
		if _, ok := service.DependsOn[model.TemplateTypeDbAdmin]; ok {
			replaceGroupDependencyMockable(service, selectedTemplates.DbAdminServices, model.TemplateTypeDbAdmin)
		}
		// Search for proxy group dependency
		if _, ok := service.DependsOn[model.TemplateTypeProxy]; ok {
			replaceGroupDependencyMockable(service, selectedTemplates.ProxyService, model.TemplateTypeProxy)
		}
		// Search for tlshelper group dependency
		if _, ok := service.DependsOn[model.TemplateTypeTlsHelper]; ok {
			replaceGroupDependencyMockable(service, selectedTemplates.TlsHelperService, model.TemplateTypeTlsHelper)
		}
	}
	stopProcess(spinner)
}

// ---------------------------------------------------------------- Private functions ---------------------------------------------------------------

func replaceGroupDependency(service *spec.ServiceConfig, templates []model.PredefinedTemplateConfig, groupName string) {
	// Add concrete service dependencies
	for _, otherService := range templates {
		if otherService.Name != service.Name {
			service.DependsOn[groupName+"-"+otherService.Name] = spec.ServiceDependency{
				Condition: spec.ServiceConditionStarted,
			}
		}
	}
	// Delete group dependency
	delete(service.DependsOn, groupName)
}
