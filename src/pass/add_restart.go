package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddRestart asks the user if he/she wants to add the restart attribute to the configuration
func AddRestart(service *spec.ServiceConfig, project *model.CGProject) {
	if project.AdvancedConfig {
		Pel()
		items := []string{"always", "on-failure", "unless-stopped", "no"}
		service.Restart = MenuQuestion("When should the service get restarted?", items)
		Pel()
	}
}
