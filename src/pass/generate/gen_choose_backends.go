/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package pass

import "compose-generator/model"

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseBackends lets the user choose predefined backend service templates
func GenerateChooseBackends(
	project *model.CGProject,
	available *model.AvailableTemplates,
	selected *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	if config != nil && config.FromFile {
		// Generate from config file
		infoLogger.Println("Generating backends from config file ...")
		selectedServiceConfigs := getServiceConfigurationsByType(config, model.TemplateTypeBackend)
		if project.Vars == nil {
			project.Vars = make(map[string]string)
		}
		for _, template := range available.BackendServices {
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
					selected.BackendServices = append(selected.BackendServices, template)
					break
				}
			}
		}
		infoLogger.Println("Generating backends from config file (done)")
	} else {
		// Generate from user input
		infoLogger.Println("Generating backends from user input ...")
		items := templateListToLabelList(available.BackendServices)
		items = append(items, "Custom backend service")
		itemsPreselected := templateListToPreselectedLabelList(available.BackendServices, selected)
		templateSelections := multiSelectMenuQuestionIndex("Which backend services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			pel()
			if index == len(available.BackendServices) { // Custom service was selected
				GenerateAddCustomService(project, model.TemplateTypeBackend)
			} else { // Predefined service was selected
				// Get selected template config
				selectedConfig := available.BackendServices[index]
				infoLogger.Println("Selected backend service: " + selectedConfig.Label)
				// Ask questions to the user
				askTemplateQuestions(project, &selectedConfig)
				// Ask proxy questions to the user
				askTemplateProxyQuestions(project, &selectedConfig, selected)
				// Evaluate proxy labels
				evaluateProxyLabels(project, &selectedConfig, selected)
				// Ask customizable secrets
				askSecretQuestions(project, &selectedConfig)
				// Ask volume questions to the user
				askForCustomVolumePaths(project, &selectedConfig)
				// Save template to the selected templates
				selected.BackendServices = append(selected.BackendServices, selectedConfig)
			}
		}
		infoLogger.Println("Generating backends from user input (done)")
	}
}
