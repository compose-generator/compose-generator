package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddContainerName asks the user if he/she wants to set the container name of a service
func AddContainerName(service *spec.ServiceConfig, project *model.CGProject) {
	service.ContainerName = util.TextQuestionWithDefault("How do you want to call your container (best practice: lower, kebab cased):", service.ContainerName)
}
