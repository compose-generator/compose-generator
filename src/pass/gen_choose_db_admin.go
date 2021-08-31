package pass

import (
	"compose-generator/model"
	"compose-generator/util"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseDbAdmins lets the user choose predefined db admin service templates
func GenerateChooseDbAdmins(
	project *model.CGProject,
	available *model.AvailableTemplates,
	selected *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	if config.FromFile {
		// Generate from config file
		selectedServiceConfigs := config.GetServiceConfigurationsByName(model.TemplateTypeDbAdmin)
		for _, template := range available.DbAdminService {
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
					selected.DbAdminService = append(selected.DbAdminService, template)
					break
				}
			}
		}
	} else {
		// Generate from user input
		availableDbAdmins := available.DbAdminService
		items := util.TemplateListToLabelList(availableDbAdmins)
		itemsPreselected := util.TemplateListToPreselectedLabelList(availableDbAdmins, selected)
		templateSelections := util.MultiSelectMenuQuestionIndex("Which db admin services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			util.Pel()
			// Get selected template config
			selectedConfig := available.DbAdminService[index]
			// Ask questions to the user
			util.AskTemplateQuestions(project, &selectedConfig)
			// Ask volume questions to the user
			util.AskForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.DbAdminService = append(selected.DbAdminService, selectedConfig)
		}
	}
}
