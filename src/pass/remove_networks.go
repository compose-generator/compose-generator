package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"strconv"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

func RemoveNetworks(service *spec.ServiceConfig, project *model.CGProject) {
	for networkName := range service.Networks {
		// Get project-wide network config by the name of the network
		networkConfig := project.Project.Networks[networkName]
		otherServices := getServicesWhichUseNetwork(networkName, service, project)
		util.Pl(networkName + ": " + strconv.Itoa(len(otherServices)))
		canBeRemoved := networkConfig.External.External && len(otherServices) == 0 || !networkConfig.External.External && len(otherServices) == 1
		// Go to next if condition is not fulfilled
		if !canBeRemoved {
			continue
		}
		// Remove network from every other service, which references it
		for _, otherService := range otherServices {
			delete(otherService.Networks, networkName)
		}
		// Remove network from project-wide network section
		delete(project.Project.Networks, networkName)
	}
}

// ---------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getServicesWhichUseNetwork(networkName string, service *spec.ServiceConfig, project *model.CGProject) []spec.ServiceConfig {
	var services []spec.ServiceConfig
	for _, otherService := range project.Project.Services {
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
