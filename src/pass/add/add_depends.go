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

// AddDepends asks the user if he/she wants to let a service depend on other services of the configuration
func AddDepends(service *spec.ServiceConfig, project *model.CGProject) {
	if len(project.Composition.Services) > 0 && yesNoQuestion("Do you want your service to depend on other services?", false) {
		pel()
		// Ask for services
		selectedServices := multiSelectMenuQuestion("Which ones?", project.Composition.ServiceNames())
		// Create map if not exists
		if service.DependsOn == nil {
			service.DependsOn = make(spec.DependsOnConfig)
		}
		// Add service dependencies
		for _, name := range selectedServices {
			service.DependsOn[name] = spec.ServiceDependency{
				Condition: spec.ServiceConditionStarted,
			}
		}
		pel()
	}
}
