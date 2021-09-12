package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddDependants asks the user if he/she wants to let other services depend on a service
func AddDependants(service *spec.ServiceConfig, project *model.CGProject) {
	if yesNoQuestion("Do you want other services depend on the new one?", false) {
		pel()
		selectedServices := multiSelectMenuQuestion("Which ones?", project.Composition.ServiceNames())
		// Add service dependencies
		for _, name := range selectedServices {
			otherService, err := project.Composition.GetService(name)
			if err != nil {
				continue
			}
			// Create map if not exists
			if otherService.DependsOn == nil {
				service.DependsOn = make(spec.DependsOnConfig)
			}
			// Add service dependency
			otherService.DependsOn[service.Name] = spec.ServiceDependency{
				Condition: spec.ServiceConditionStarted,
			}
		}
		pel()
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

/*
Idea: when it is possible to get from the current service via a dependency path to the service with the name 'otherServiceName', we can't add a
dependency from the other service to the current one, because our directed graph would coutain a cycle. This algorithm only works when the original
graph is acyclic. This is given as we check that in the beginning.
*/
func checkForDependencyCycle(otherServiceName string, project *model.CGProject) bool {

	return false
}
