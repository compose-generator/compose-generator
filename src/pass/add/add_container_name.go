package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddContainerName asks the user if he/she wants to set the container name of a service
func AddContainerName(service *spec.ServiceConfig, project *model.CGProject) {
	service.ContainerName = textQuestionWithDefault("How do you want to call your container (best practice: lower, kebab cased):", service.ContainerName)
}
