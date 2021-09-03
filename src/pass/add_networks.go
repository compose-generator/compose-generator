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
var askForNewNetworkMockable = askForExternalNetwork

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddNetworks asks the user if he/she wants to add some networks to the configuration
func AddNetworks(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if YesNoQuestion("Do you want to add networks to your service?", false) {
		Pel()
		for ok := true; ok; ok = YesNoQuestion("Add another network?", true) {
			globalNetwork := YesNoQuestion("Do you want to add an external network (y) or create a new one (N)?", false)
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
		Error("Error parsing external networks", err, false)
		return
	}
	if len(externalNetworks) == 0 {
		Error("There is no external network existing", nil, false)
		return
	}
	// Let the user choose one
	menuItems := []string{}
	for _, network := range externalNetworks {
		menuItems = append(menuItems, network.Name)
	}
	index := MenuQuestionIndex("Which one?", menuItems)
	selectedNetwork := externalNetworks[index]

	// Ask for a custom name within the compose file
	customName := TextQuestionWithDefault("How do you want to call the network internally?", selectedNetwork.Name)

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
	service.Networks[customName] = nil
	// Add network to project-wide network section
	if project.Composition.Networks == nil {
		project.Composition.Networks = make(map[string]spec.NetworkConfig)
	}
	project.Composition.Networks[customName] = spec.NetworkConfig{
		Name: customName,
		External: spec.External{
			Name:     selectedNetwork.Name,
			External: true,
		},
	}
}

func askForNewNetwork(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	// Ask user to add a new network
	networkName := TextQuestion("How do you want to call the new network?")
	external := YesNoQuestion("Do you want to create it as an external network and link it in?", false)
	externalConfig := spec.External{}
	if external {
		// Create external network
		err := CreateDockerNetwork(client, networkName)
		if err != nil {
			Error("External network could not be created", err, false)
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
}
