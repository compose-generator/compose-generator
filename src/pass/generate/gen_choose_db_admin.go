/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseDbAdmins lets the user choose predefined db admin service templates
func GenerateChooseDbAdmins(
	project *model.CGProject,
	available *model.AvailableTemplates,
	selected *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	if config != nil && config.FromFile {
		// Generate from config file
		selectedServiceConfigs := getServiceConfigurationsByType(config, model.TemplateTypeDbAdmin)
		if project.Vars == nil {
			project.Vars = make(map[string]string)
		}
		for _, template := range available.DbAdminServices {
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
					selected.DbAdminServices = append(selected.DbAdminServices, template)
					break
				}
			}
		}
	} else {
		// Generate from user input
		items := templateListToLabelList(available.DbAdminServices)
		items = append(items, "Custom db admin service")
		itemsPreselected := templateListToPreselectedLabelList(available.DbAdminServices, selected)
		templateSelections := multiSelectMenuQuestionIndex("Which db admin services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			pel()
			if index == len(available.DbAdminServices) { // Custom service was selected
				GenerateAddCustomService(project, model.TemplateTypeDbAdmin)
			} else { // Predefined service was selected
				// Get selected template config
				selectedConfig := available.DbAdminServices[index]
				// Ask questions to the user
				askTemplateQuestions(project, &selectedConfig)
				// Ask proxy questions to the user
				askTemplateProxyQuestions(project, &selectedConfig, selected)
				// Ask volume questions to the user
				askForCustomVolumePaths(project, &selectedConfig)
				// Save template to the selected templates
				selected.DbAdminServices = append(selected.DbAdminServices, selectedConfig)
			}
		}
	}
}
