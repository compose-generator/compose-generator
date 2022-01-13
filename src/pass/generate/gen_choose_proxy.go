/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseProxies lets the user choose predefined proxy service templates
func GenerateChooseProxies(
	project *model.CGProject,
	available *model.AvailableTemplates,
	selected *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	if config != nil && config.FromFile {
		// Generate from config file
		infoLogger.Println("Generating proxies from config file ...")
		selectedServiceConfigs := getServiceConfigurationsByType(config, model.TemplateTypeProxy)
		if project.Vars == nil {
			project.Vars = make(map[string]string)
		}
		for _, template := range available.ProxyService {
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
					for _, question := range template.Volumes {
						if value, ok := selectedConfig.Params[question.Variable]; ok {
							project.Vars[question.Variable] = value
						} else {
							project.Vars[question.Variable] = question.DefaultValue
						}
					}
					// Add template to selected templates
					selected.ProxyServices = append(selected.ProxyServices, template)
					break
				}
			}
		}
		infoLogger.Println("Generating proxies from config file (done)")
	} else {
		// Generate from user input
		infoLogger.Println("Generating proxies from user input ...")
		items := templateListToLabelList(available.ProxyService)
		itemsPreselected := templateListToPreselectedLabelList(available.ProxyService, selected)
		templateSelections := multiSelectMenuQuestionIndex("Which proxy services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			pel()
			// Get selected template config
			selectedConfig := available.ProxyService[index]
			// Ask questions to the user
			askTemplateQuestions(project, &selectedConfig)
			// Ask volume questions to the user
			askForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.ProxyServices = append(selected.ProxyServices, selectedConfig)
		}
		infoLogger.Println("Generating proxies from user input (done)")
	}
}
