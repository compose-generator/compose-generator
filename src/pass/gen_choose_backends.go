package pass

import (
	"compose-generator/model"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseBackends lets the user choose predefined backend service templates
func GenerateChooseBackends(
	project *model.CGProject,
	available *model.AvailableTemplates,
	selected *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	if config.FromFile {
		// Generate from config file
		selectedServiceConfigs := GetServiceConfigurationsByName(config, model.TemplateTypeBackend)
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
					// Add template to selected templates
					selected.BackendServices = append(selected.BackendServices, template)
					break
				}
			}
		}
	} else {
		// Generate from user input
		availableBackends := available.BackendServices
		items := TemplateListToLabelList(availableBackends)
		itemsPreselected := TemplateListToPreselectedLabelList(availableBackends, selected)
		templateSelections := MultiSelectMenuQuestionIndex("Which backends services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			Pel()
			// Get selected template config
			selectedConfig := available.BackendServices[index]
			// Ask questions to the user
			AskTemplateQuestions(project, &selectedConfig)
			// Ask volume questions to the user
			AskForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.BackendServices = append(selected.BackendServices, selectedConfig)
		}
	}
}
