/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
)

var checkForDependencyCycleMockable = checkForDependencyCycle

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddDependants asks the user if he/she wants to let other services depend on a service
func AddDependants(service *spec.ServiceConfig, project *model.CGProject) {
	if len(project.Composition.Services) > 0 && yesNoQuestion("Do you want other services depend on the new one?", false) {
		pel()
		selectedServices := multiSelectMenuQuestion("Which ones?", project.Composition.ServiceNames())
		infoLogger.Println("Selected dependants: " + strings.Join(selectedServices, ", "))
		// Add service dependencies
		for _, name := range selectedServices {
			otherService, err := project.Composition.GetService(name)
			if err != nil {
				continue
			}
			// Check if the dependency would produce a cycle
			if checkForDependencyCycleMockable(service, otherService.Name, project.Composition) {
				warningLogger.Println("Could not add dependency from '" + otherService.Name + "' to '" + service.Name + "' because it would cause a cycle")
				logWarning("Could not add dependency from '" + otherService.Name + "' to '" + service.Name + "' because it would cause a cycle")
				continue
			}
			// Create map if not exists
			if otherService.DependsOn == nil {
				otherService.DependsOn = make(spec.DependsOnConfig)
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
func checkForDependencyCycle(currentService *spec.ServiceConfig, otherServiceName string, project *spec.Project) bool {
	for dependency := range currentService.DependsOn {
		visitedServices := []string{currentService.Name, otherServiceName}
		if visitServiceDependencies(project, dependency, &visitedServices) {
			return true
		}
	}
	return false
}
