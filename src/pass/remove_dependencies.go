package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// RemoveDependencies removes all dependencies on a service from all other services of the configuration
func RemoveDependencies(service *spec.ServiceConfig, project *model.CGProject) {
	for _, otherService := range project.Composition.Services {
		for dependency := range otherService.DependsOn {
			if dependency == service.Name {
				delete(otherService.DependsOn, service.Name)
				break
			}
		}
	}
}
