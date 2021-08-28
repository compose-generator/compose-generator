package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"context"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddNetworks asks the user if he/she wants to add some networks to the configuration
func AddNetworks(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if util.YesNoQuestion("Do you want to add networks to your service?", false) {
		util.Pel()
		for ok := true; ok; ok = util.YesNoQuestion("Add another network?", true) {
			globalNetwork := util.YesNoQuestion("Do you want to add an external network (y) or create a new one (N)?", false)
			if globalNetwork {
				askForExternalNetwork(service, project, client)
			} else {
				askForNewNetwork(service, project, client)
			}
		}
	}
}

// ---------------------------------------------------------------- Private functions ---------------------------------------------------------------

func askForExternalNetwork(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	// Search for external networks
	externalNetworks, err := client.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		util.Error("Error parsing external networks", err, false)
		return
	}
	if externalNetworks == nil || len(externalNetworks) == 0 {
		util.Error("There is no external network existing", nil, false)
		return
	}
	// Let the user choose one
	menuItems := []string{}
	for _, network := range externalNetworks {
		menuItems = append(menuItems, network.Name)
	}
	index := util.MenuQuestionIndex("Which one?", menuItems)
	selectedNetwork := externalNetworks[index]

	// Ask for a custom name within the compose file
	customName := util.TextQuestionWithDefault("How do you want to call the network internally?", selectedNetwork.Name)

	// Create maps if not exists
	if service.Networks == nil {
		service.Networks = make(map[string]*spec.ServiceNetworkConfig)
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
	networkName := util.TextQuestion("How do you want to call the new network?")
	external := util.YesNoQuestion("Do you want to create it as an external network and link it in?", false)
	externalConfig := spec.External{}
	if external {
		// Create external network
		_, err := client.NetworkCreate(context.Background(), networkName, types.NetworkCreate{
			Internal: false,
		})
		if err != nil {
			util.Error("External network could not be created", err, false)
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
