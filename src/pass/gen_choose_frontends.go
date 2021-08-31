package pass

import (
	"compose-generator/model"
	"compose-generator/util"
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
		selectedServiceConfigs := config.GetServiceConfigurationsByName(model.TemplateTypeFrontend)
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
		items := util.TemplateListToLabelList(availableFrontends)
		itemsPreselected := util.TemplateListToPreselectedLabelList(availableFrontends, selected)
		templateSelections := util.MultiSelectMenuQuestionIndex("Which frontend services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			util.Pel()
			// Get selected template config
			selectedConfig := available.FrontendServices[index]
			// Ask questions to the user
			util.AskTemplateQuestions(project, &selectedConfig)
			// Ask volume questions to the user
			util.AskForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.FrontendServices = append(selected.FrontendServices, selectedConfig)
		}
	}
}
