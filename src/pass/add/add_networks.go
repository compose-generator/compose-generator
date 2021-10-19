/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"context"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// CreateDockerNetwork calls the Docker client to create a new network
var CreateDockerNetwork = func(client *client.Client, networkName string) error {
	_, err := client.NetworkCreate(context.Background(), networkName, types.NetworkCreate{
		Internal: false,
	})
	return err
}

// ListDockerNetworks calls the Docker client to list all available networks
var ListDockerNetworks = func(client *client.Client) ([]types.NetworkResource, error) {
	return client.NetworkList(context.Background(), types.NetworkListOptions{})
}

var askForExternalNetworkMockable = askForExternalNetwork
var askForNewNetworkMockable = askForNewNetwork

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddNetworks asks the user if he/she wants to add some networks to the configuration
func AddNetworks(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if yesNoQuestion("Do you want to add networks to your service?", false) {
		pel()
		for ok := true; ok; ok = yesNoQuestion("Add another network?", true) {
			globalNetwork := yesNoQuestion("Do you want to add an external network (y) or create a new one (N)?", false)
			if globalNetwork {
				askForExternalNetworkMockable(service, project, client)
			} else {
				askForNewNetworkMockable(service, project, client)
			}
		}
	}
}

// ---------------------------------------------------------------- Private functions ---------------------------------------------------------------

func askForExternalNetwork(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	// Search for external networks
	externalNetworks, err := ListDockerNetworks(client)
	if err != nil {
		errorLogger.Println("Error parsing external networks: " + err.Error())
		logError("Error parsing external networks", false)
		return
	}
	if len(externalNetworks) == 0 {
		errorLogger.Println("There is no external network existing")
		logError("There is no external network existing", false)
		return
	}
	// Let the user choose one
	menuItems := []string{}
	for _, network := range externalNetworks {
		menuItems = append(menuItems, network.Name)
	}
	index := menuQuestionIndex("Which one?", menuItems)
	selectedNetwork := externalNetworks[index]

	// Ask for a custom name within the compose file
	networkName := textQuestionWithDefault("How do you want to call the network internally?", selectedNetwork.Name)

	// Create maps if not exists
	if service.Networks == nil {
		service.Networks = make(map[string]*spec.ServiceNetworkConfig)
	}
	if project.Composition == nil {
		project.Composition = &spec.Project{}
	}
	if project.Composition.Networks == nil {
		project.Composition.Networks = make(spec.Networks)
	}
	// Add network to the service
	service.Networks[networkName] = nil
	// Add network to project-wide network section
	if project.Composition.Networks == nil {
		project.Composition.Networks = make(map[string]spec.NetworkConfig)
	}
	project.Composition.Networks[networkName] = spec.NetworkConfig{
		Name: networkName,
		External: spec.External{
			Name:     selectedNetwork.Name,
			External: true,
		},
	}
	infoLogger.Println("Added external network '" + networkName + "' to the new service")
}

func askForNewNetwork(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	// Ask user to add a new network
	networkName := textQuestion("How do you want to call the new network?")
	external := yesNoQuestion("Do you want to create it as an external network and link it in?", false)
	externalConfig := spec.External{}
	if external {
		// Create external network
		err := CreateDockerNetwork(client, networkName)
		if err != nil {
			errorLogger.Println("External network could not be created: " + err.Error())
			logError("External network could not be created", false)
			return
		}
		externalConfig = spec.External{
			External: true,
			Name:     networkName,
		}
	}
	// Create maps if not exists
	if service.Networks == nil {
		service.Networks = make(map[string]*spec.ServiceNetworkConfig)
	}
	if project.Composition == nil {
		project.Composition = &spec.Project{}
	}
	if project.Composition.Networks == nil {
		project.Composition.Networks = make(spec.Networks)
	}
	// Add network to the service
	service.Networks[networkName] = &spec.ServiceNetworkConfig{}
	// Add network to project-wide network section
	project.Composition.Networks[networkName] = spec.NetworkConfig{
		Name:     networkName,
		External: externalConfig,
	}
	infoLogger.Println("Added new network '" + networkName + "' to the new service")
}
