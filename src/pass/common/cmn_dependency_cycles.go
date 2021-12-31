/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
)

var hasDependencyCyclesMockable = hasDependencyCycles
var visitServiceDependenciesMockable = VisitServiceDependencies

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// CommonCheckForDependencyCycles ensures that the project contains no dependency cycles
func CommonCheckForDependencyCycles(project *model.CGProject) {
	infoLogger.Println("Checking for dependency cycles ...")
	if hasDependencyCyclesMockable(project.Composition) {
		errorLogger.Println("Configuration contains dependency cycles")
		logError("Configuration contains dependency cycles", true)
	}
	infoLogger.Println("Checking for dependency cycles (done)")
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
		if sliceContainsString(*visitedServices, dependency) {
			return true
		}
		if VisitServiceDependencies(p, dependency, visitedServices) {
			return true
		}
	}
	return false
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func hasDependencyCycles(project *spec.Project) bool {
	for _, service := range project.Services {
		visitedServices := []string{service.Name}
		if visitServiceDependenciesMockable(project, service.Name, &visitedServices) {
			return true
		}
	}
	return false
}
