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
	if config.FromFile {
		// Generate from config file
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
					// Add template to selected templates
					selected.ProxyService = append(selected.ProxyService, template)
					break
				}
			}
		}
	} else {
		// Generate from user input
		availableProxies := available.ProxyService
		items := templateListToLabelList(availableProxies)
		itemsPreselected := templateListToPreselectedLabelList(availableProxies, selected)
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
			selected.ProxyService = append(selected.ProxyService, selectedConfig)
		}
	}
}
