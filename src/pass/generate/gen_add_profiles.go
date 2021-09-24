package pass

import "compose-generator/model"

const (
	ProfileDev        = "dev"
	ProfileProduction = "production"
)

func GenAddProfiles(project *model.CGProject) {
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
