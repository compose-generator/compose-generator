/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateResolveDependencyGroups resolves group dependencies like 'database' or 'frontend' to concrete service dependencies
func GenerateResolveDependencyGroups(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	infoLogger.Println("Resolving group dependencies ...")
	spinner := startProcess("Resolving group dependencies ...")
	for i := range project.Composition.Services {
		service := &project.Composition.Services[i]
		// Search for frontend group dependency
		if _, ok := service.DependsOn[model.TemplateTypeFrontend]; ok {
			replaceGroupDependency(service, selectedTemplates.FrontendServices, model.TemplateTypeFrontend)
		}
		// Search for backend group dependency
		if _, ok := service.DependsOn[model.TemplateTypeBackend]; ok {
			replaceGroupDependency(service, selectedTemplates.BackendServices, model.TemplateTypeBackend)
		}
		// Search for database group dependency
		if _, ok := service.DependsOn[model.TemplateTypeDatabase]; ok {
			replaceGroupDependency(service, selectedTemplates.DatabaseServices, model.TemplateTypeDatabase)
		}
		// Search for dbadmin group dependency
		if _, ok := service.DependsOn[model.TemplateTypeDbAdmin]; ok {
			replaceGroupDependency(service, selectedTemplates.DbAdminServices, model.TemplateTypeDbAdmin)
		}
		// Search for proxy group dependency
		if _, ok := service.DependsOn[model.TemplateTypeProxy]; ok {
			replaceGroupDependency(service, selectedTemplates.ProxyServices, model.TemplateTypeProxy)
		}
		// Search for tlshelper group dependency
		if _, ok := service.DependsOn[model.TemplateTypeTlsHelper]; ok {
			replaceGroupDependency(service, selectedTemplates.TlsHelperServices, model.TemplateTypeTlsHelper)
		}
	}
	stopProcess(spinner)
	infoLogger.Println("Resolving group dependencies (done)")
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
