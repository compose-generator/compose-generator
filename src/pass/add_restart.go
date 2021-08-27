package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// AddRestart asks the user if he/she wants to add the restart attribute to the configuration
func AddRestart(service *spec.ServiceConfig, project *model.CGProject) {
	if project.AdvancedConfig {
		util.Pel()
		items := []string{"always", "on-failure", "unless-stopped", "no"}
		service.Restart = util.MenuQuestion("When should the service get restarted?", items)
		util.Pel()
	}
}
