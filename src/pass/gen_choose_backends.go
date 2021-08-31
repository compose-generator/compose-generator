package pass

import (
	"compose-generator/model"
	"compose-generator/util"
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
		selectedServiceConfigs := config.GetServiceConfigurationsByName(model.TemplateTypeBackend)
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
		items := util.TemplateListToLabelList(availableBackends)
		itemsPreselected := util.TemplateListToPreselectedLabelList(availableBackends, selected)
		templateSelections := util.MultiSelectMenuQuestionIndex("Which backends services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			util.Pel()
			// Get selected template config
			selectedConfig := available.BackendServices[index]
			// Ask questions to the user
			util.AskTemplateQuestions(project, &selectedConfig)
			// Ask volume questions to the user
			util.AskForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.BackendServices = append(selected.BackendServices, selectedConfig)
		}
	}
}
