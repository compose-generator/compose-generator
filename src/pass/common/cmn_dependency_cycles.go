package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// CommonCheckForDependencyCycles ensures that the project contains no dependency cycles
func CommonCheckForDependencyCycles(project *model.CGProject) {
	if hasDependencyCycles(project.Composition) {
		util.Error("Configuration contains dependency cycles", nil, true)
	}
}

// VisitServiceDependencies checks a particular service for dependency cycles
func VisitServiceDependencies(p *spec.Project, currentServiceName string, visitedServices *[]string) bool {
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
		if util.SliceContainsString(*visitedServices, dependency) {
			return true
		}
		return VisitServiceDependencies(p, dependency, visitedServices)
	}
	return false
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func hasDependencyCycles(project *spec.Project) bool {
	for _, service := range project.Services {
		visitedServices := []string{service.Name}
		if VisitServiceDependencies(project, service.Name, &visitedServices) {
			return true
		}
	}
	return false
}
