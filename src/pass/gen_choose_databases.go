package pass

import (
	"compose-generator/model"
	"compose-generator/util"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseDatabases lets the user choose predefined database service templates
func GenerateChooseDatabases(
	project *model.CGProject,
	available *model.AvailableTemplates,
	selected *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	if config.FromFile {
		// Generate from config file
		selectedServiceConfigs := config.GetServiceConfigurationsByName(model.TemplateTypeDatabase)
		for _, template := range available.DatabaseServices {
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
					selected.DatabaseServices = append(selected.DatabaseServices, template)
					break
				}
			}
		}
	} else {
		// Generate from user input
		availableDatabases := available.DatabaseServices
		items := util.TemplateListToLabelList(availableDatabases)
		itemsPreselected := util.TemplateListToPreselectedLabelList(availableDatabases, selected)
		templateSelections := util.MultiSelectMenuQuestionIndex("Which database services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			util.Pel()
			// Get selected template config
			selectedConfig := available.DatabaseServices[index]
			// Ask questions to the user
			util.AskTemplateQuestions(project, &selectedConfig)
			// Ask volume questions to the user
			util.AskForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.DatabaseServices = append(selected.DatabaseServices, selectedConfig)
		}
	}
}
