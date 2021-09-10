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
