/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

var getServicesWhichUseNetworkMockable = getServicesWhichUseNetwork

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// RemoveNetworks removes all networks from a service
func RemoveNetworks(service *spec.ServiceConfig, project *model.CGProject) {
	infoLogger.Println("Removing networks ...")
	for networkName := range service.Networks {
		// Get project-wide network config by the name of the network
		networkConfig := project.Composition.Networks[networkName]
		otherServices := getServicesWhichUseNetworkMockable(networkName, service, project)
		canBeRemoved := networkConfig.External.External && len(otherServices) == 0 || !networkConfig.External.External && len(otherServices) <= 1
		// Go to next if condition is not fulfilled
		if !canBeRemoved {
			continue
		}
		// Remove network from every other service, which references it
		for _, otherService := range otherServices {
			delete(otherService.Networks, networkName)
		}
		// Remove network from project-wide network section
		delete(project.Composition.Networks, networkName)
	}
	infoLogger.Println("Removing networks (done)")
}

// ---------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getServicesWhichUseNetwork(networkName string, service *spec.ServiceConfig, project *model.CGProject) []spec.ServiceConfig {
	var services []spec.ServiceConfig
	for _, otherService := range project.Composition.Services {
		// Skip the service, we are currently editing
		if otherService.Name == service.Name {
			continue
		}
		// Search through networks of all other services
		for otherNetworkName := range otherService.Networks {
			if networkName == otherNetworkName {
				// Another service needs the same network
				services = append(services, otherService)
				break
			}
		}
	}
	return services
}
