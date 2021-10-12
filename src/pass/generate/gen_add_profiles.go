/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import "compose-generator/model"

const (
	// ProfileDev is the name of the profile for development
	ProfileDev        = "dev"
	// ProfileProduction is the name of the profile for production
	ProfileProduction = "prod"
)

// GenerateAddProfiles adds two profiles to the project in case the production-ready variant was selected
func GenerateAddProfiles(project *model.CGProject) {
	if project.CGProjectMetadata.ProductionReady {
		spinner := startProcess("Adding dev and production profiles")
		for i := range project.Composition.Services {
			service := &project.Composition.Services[i]
			if len(service.Profiles) == 0 {
				service.Profiles = append(service.Profiles, ProfileDev, ProfileProduction)
			}
		}
		stopProcess(spinner)
	}
}
