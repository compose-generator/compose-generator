package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

func CommonCheckForDependencyCycles(project *model.CGProject) {
	if hasDependencyCycles(project.Composition) {
		util.Error("Configuration contains dependency cycles", nil, true)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func hasDependencyCycles(project *spec.Project) bool {
	for _, service := range project.Services {
		visitedServices := []string{service.Name}
		if visitServiceDependencies(project, service.Name, &visitedServices) {
			return true
		}
	}
	return false
}

func visitServiceDependencies(p *spec.Project, currentServiceName string, visitedServices *[]string) bool {
	// Get service
	service, err := p.GetService(currentServiceName)
	if err != nil {
		return false
	}
	// Add current item to visited services list
	*visitedServices = append(*visitedServices, currentServiceName)
	// Visit dependencies
	for dependency := range service.DependsOn {
		// Check if the service was already visited
		for _, item := range *visitedServices {
			if item == dependency {
				return true
			}
		}
		return visitServiceDependencies(p, dependency, visitedServices)
	}
	return false
}
