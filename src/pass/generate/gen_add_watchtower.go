package pass

import (
	"compose-generator/model"

	"github.com/compose-spec/compose-go/types"
)

func GenerateAddWatchtower(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	// Ask the user if watchtower should be added to the project
	addWatchtower := yesNoQuestion("Do you want to add Watchtower to check for new image versions?", false)
	if addWatchtower {
		// Ask which services should be equipped with image update detection
		templates := selectedTemplates.GetAll()
		selectedIndices := multiSelectMenuQuestionIndex("For which services do you want to add Watchtower?", selectedTemplates.GetAllLabels(), []string{})
		for _, i := range selectedIndices {
			template := templates[i]
			// Add label to service
			if service, err := project.Composition.GetService(template.Name); err != nil {
				service.Labels["com.centurylinklabs.watchtower.enable"] = "true"
			}
		}

		// Add watchtower service
		project.Composition.Services = append(project.Composition.Services, types.ServiceConfig{
			Name: "companion-watchtower",
			Volumes: []types.ServiceVolumeConfig{
				{
					Type:   types.VolumeTypeBind,
					Source: "/var/run/docker.sock",
					Target: "/var/run/docker.sock",
				},
			},
			Command: types.ShellCommand{"--interval", "30"},
		})
	}
}
