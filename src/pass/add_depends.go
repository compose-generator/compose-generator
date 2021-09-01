package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddDepends asks the user if he/she wants to let a service depend on other services of the configuration
func AddDepends(service *spec.ServiceConfig, project *model.CGProject) {
	if util.YesNoQuestion("Do you want your service depend on other services?", false) {
		util.Pel()
		// Ask for services
		selectedServices := util.MultiSelectMenuQuestion("Which ones?", project.Composition.ServiceNames())
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
		util.Pel()
	}
}
