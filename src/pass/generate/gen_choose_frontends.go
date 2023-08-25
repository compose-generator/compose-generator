/*
Copyright © 2021-2023 Compose Generator Contributors
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
	if config != nil && config.FromFile {
		// Generate from config file
		infoLogger.Println("Generating frontends from config file ...")
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
					for _, question := range template.ProxyQuestions {
						if value, ok := selectedConfig.Params[question.Variable]; ok {
							project.Vars[question.Variable] = value
						} else {
							project.Vars[question.Variable] = question.DefaultValue
						}
					}
					for _, question := range template.Volumes {
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
		infoLogger.Println("Generating frontends from config file (done)")
	} else {
		// Generate from user input
		infoLogger.Println("Generating frontends from user input ...")
		items := templateListToLabelList(available.FrontendServices)
		items = append(items, "Custom frontend service")
		itemsPreselected := templateListToPreselectedLabelList(available.FrontendServices, selected)
		templateSelections := multiSelectMenuQuestionIndex("Which frontend services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			pel()
			if index == len(available.FrontendServices) { // Custom service was selected
				GenerateAddCustomService(project, model.TemplateTypeFrontend)
			} else { // Predefined service was selected
				// Get selected template config
				selectedConfig := available.FrontendServices[index]
				// Ask questions to the user
				askTemplateQuestions(project, &selectedConfig)
				// Ask proxy questions to the user
				askTemplateProxyQuestions(project, &selectedConfig, selected)
				// Ask customizable secrets
				askSecretQuestions(project, &selectedConfig)
				// Evaluate proxy labels
				evaluateProxyLabels(project, &selectedConfig, selected)
				// Ask volume questions to the user
				askForCustomVolumePaths(project, &selectedConfig)
				// Save template to the selected templates
				selected.FrontendServices = append(selected.FrontendServices, selectedConfig)
			}
		}
		infoLogger.Println("Generating frontends from user input (done)")
	}
}
