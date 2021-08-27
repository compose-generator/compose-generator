package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddDependants asks the user if he/she wants to let other services depend on a service
func AddDependants(service *spec.ServiceConfig, project *model.CGProject) {
	if util.YesNoQuestion("Do you want other services depend on the new one?", false) {
		util.Pel()
		selectedServices := util.MultiSelectMenuQuestion("Which ones?", project.Project.ServiceNames())
		// Add service dependencies
		for _, name := range selectedServices {
			otherService, err := project.Project.GetService(name)
			if err != nil {
				util.Error("Selected service '"+name+"' was not found", err, false)
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
		util.Pel()
	}
}
