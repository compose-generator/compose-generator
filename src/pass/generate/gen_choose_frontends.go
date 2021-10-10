/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseFrontends lets the user choose predefined frontend service templates
func GenerateChooseFrontends(
	project *model.CGProject,
	available *model.AvailableTemplates,
	selected *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	if config.FromFile {
		// Generate from config file
		selectedServiceConfigs := getServiceConfigurationsByType(config, model.TemplateTypeFrontend)
		if project.Vars == nil {
			project.Vars = make(map[string]string)
		}
		for _, template := range available.FrontendServices {
			for _, selectedConfig := range selectedServiceConfigs {
				if template.Name == selectedConfig.Name {
					// Add vars to project
					for _, question := range template.Questions {
						if value, ok := selectedConfig.Params[question.Variable]; ok {
							project.Vars[question.Variable] = value
						} else {
							project.Vars[question.Variable] = question.DefaultValue
						}
					}
					// Add template to selected templates
					selected.FrontendServices = append(selected.FrontendServices, template)
					break
				}
			}
		}
	} else {
		// Generate from user input
		availableFrontends := available.FrontendServices
		items := templateListToLabelList(availableFrontends)
		itemsPreselected := templateListToPreselectedLabelList(availableFrontends, selected)
		templateSelections := multiSelectMenuQuestionIndex("Which frontend services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			pel()
			// Get selected template config
			selectedConfig := available.FrontendServices[index]
			// Ask questions to the user
			askTemplateQuestions(project, &selectedConfig)
			// Ask proxy questions to the user
			askTemplateProxyQuestions(project, &selectedConfig, selected)
			// Ask volume questions to the user
			askForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.FrontendServices = append(selected.FrontendServices, selectedConfig)
		}
	}
}
