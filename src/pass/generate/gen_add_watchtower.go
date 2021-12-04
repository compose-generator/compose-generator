package pass

import (
	"compose-generator/model"
	"fmt"
)

func GenerateAddWatchtower(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	// Ask the user if watchtower should be added to the project
	addWatchtower := yesNoQuestion("Do you want to add Watchtower to check for new image versions?", false)
	if addWatchtower {
		// Ask which services should be equipped with image update detection
		selectedLabels := multiSelectMenuQuestion("For which services do you want to add Watchtower?", selectedTemplates.GetAllLabels())
		for _, label := range selectedLabels {
			fmt.Println(label)
		}
	}
}
