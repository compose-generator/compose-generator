package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddName asks the user if he/she wants to set a name for a service
func AddName(service *spec.ServiceConfig, project *model.CGProject) {
	chooseAgain := true
	for chooseAgain {
		name := util.TextQuestionWithDefault("How do you want to call your service (best practice: lower, kebab cased):", service.Name)
		if util.SliceContainsString(project.Composition.ServiceNames(), name) {
			util.Error("This service name already exists. Please choose a different one", nil, false)
		} else {
			chooseAgain = false
		}
	}
}
