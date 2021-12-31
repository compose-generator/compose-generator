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

	"github.com/compose-spec/compose-go/types"
)

func GenerateAddWatchtower(
	project *model.CGProject,
	selectedTemplates *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	// Ask the user if watchtower should be added to the project
	pel()
	templates := selectedTemplates.GetAllRef()
	addWatchtower := false
	if config == nil || !config.FromFile {
		addWatchtower = yesNoQuestion("Do you want to add Watchtower to check for new image versions?", false)
	}
	if addWatchtower {
		infoLogger.Println("Adding Watchtower to the project ...")
		// Ask which services should be equipped with image update detection
		allLabels := []string{}
		preselectedLabels := []string{}
		for _, template := range templates {
			allLabels = append(allLabels, template.Label)
			if template.AutoUpdated {
				preselectedLabels = append(preselectedLabels, template.Label)
				template.AutoUpdated = false
			}
		}
		// Ask the user which of the services should be auto-updated
		selectedIndices := multiSelectMenuQuestionIndex("For which services do you want to add Watchtower?", allLabels, preselectedLabels)
		for _, i := range selectedIndices {
			template := templates[i]
			template.AutoUpdated = true
		}

		// Add watchtower service
		project.Composition.Services = append(project.Composition.Services, types.ServiceConfig{
			Name:    "companion-watchtower",
			Image:   "containrrr/watchtower:latest",
			Restart: types.RestartPolicyUnlessStopped,
			Volumes: []types.ServiceVolumeConfig{
				{
					Type:   types.VolumeTypeBind,
					Source: "/var/run/docker.sock",
					Target: "/var/run/docker.sock",
				},
			},
			DependsOn: types.DependsOnConfig{
				model.TemplateTypeFrontend:  types.ServiceDependency{},
				model.TemplateTypeBackend:   types.ServiceDependency{},
				model.TemplateTypeDatabase:  types.ServiceDependency{},
				model.TemplateTypeDbAdmin:   types.ServiceDependency{},
				model.TemplateTypeProxy:     types.ServiceDependency{},
				model.TemplateTypeTlsHelper: types.ServiceDependency{},
			},
			Command: types.ShellCommand{"--interval", "30"},
		})
		infoLogger.Println("Adding Watchtower to the project .. (done)")
	} else {
		// Remove all auto-updated flags
		for _, template := range templates {
			template.AutoUpdated = false
		}
	}
}
