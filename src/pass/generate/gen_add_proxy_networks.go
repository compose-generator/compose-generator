package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateAddProxyNetworks
func GenerateAddProxyNetworks(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	if project.ProductionReady {
		// Get reference of proxy service
		proxyService := getServiceRef(project.Composition, "proxy-"+selectedTemplates.ProxyService[0].Name)
		if proxyService == nil {
			printError("Proxy service cannot be found for network inserting", nil, true)
			return
		}
		if proxyService.Networks == nil {
			proxyService.Networks = make(map[string]*spec.ServiceNetworkConfig)
		}
		// Couple every proxied frontend, backend, database and dbadmin service with the proxy in a network
		for _, template := range selectedTemplates.GetAll() {
			if template.Type != model.TemplateTypeProxy && template.Type != model.TemplateTypeTlsHelper && template.Proxied {
				networkName := "proxy-" + template.Name
				// Get reference to current service
				service := getServiceRef(project.Composition, template.Type+"-"+template.Name)
				if service == nil {
					continue
				}
				// Add network to proxy and current service
				if service.Networks == nil {
					service.Networks = make(map[string]*spec.ServiceNetworkConfig)
				}
				service.Networks[networkName] = nil
				proxyService.Networks[networkName] = nil
				// Remove all exposed ports from the proxied service
				service.Ports = []spec.ServicePortConfig{}
			}
		}
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getServiceRef(project *spec.Project, serviceName string) *spec.ServiceConfig {
	for index := range project.Services {
		service := project.Services[index]
		if service.Name == serviceName {
			return &service
		}
	}
	return nil
}
