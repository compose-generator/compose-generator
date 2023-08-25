/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddRestart asks the user if he/she wants to add the restart attribute to the configuration
func AddRestart(service *spec.ServiceConfig, project *model.CGProject) {
	if project.AdvancedConfig {
		pel()
		items := []string{"always", "on-failure", "unless-stopped", "no"}
		service.Restart = menuQuestion("When should the service get restarted?", items)
		infoLogger.Println("Restart service: " + service.Restart)
		pel()
	}
}
