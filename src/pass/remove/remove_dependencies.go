/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// RemoveDependencies removes all dependencies on a service from all other services of the configuration
func RemoveDependencies(service *spec.ServiceConfig, project *model.CGProject) {
	util.InfoLogger.Println("Removing dependencies ...")
	for _, otherService := range project.Composition.Services {
		for dependency := range otherService.DependsOn {
			if dependency == service.Name {
				delete(otherService.DependsOn, service.Name)
				break
			}
		}
	}
	util.InfoLogger.Println("Removing dependencies done")
}
