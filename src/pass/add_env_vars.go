package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddEnvVars asks the user if he/she wants to add environment variables to the configuration
func AddEnvVars(service *spec.ServiceConfig, _ *model.CGProject) {
	if util.YesNoQuestion("Do you want to provide environment variables to your service?", false) {
		util.Pel()
		if service.Environment == nil {
			service.Environment = make(map[string]*string)
		}
		for another := true; another; another = util.YesNoQuestion("Expose another environment variable?", true) {
			// Ask for name and value
			variableName := util.TextQuestionWithValidator("Variable name (BEST_PRACTICE_IS_CAPS):", util.EnvVarNameValidator)
			variableValue := util.TextQuestion("Variable value:")
			// Add env var to service
			service.Environment[variableName] = &variableValue
		}
		util.Pel()
	}
}
