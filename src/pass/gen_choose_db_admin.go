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
	if config.FromFile {
		// Generate from config file
		selectedServiceConfigs := GetServiceConfigurationsByName(config, model.TemplateTypeDbAdmin)
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
					// Add template to selected templates
					selected.DbAdminServices = append(selected.DbAdminServices, template)
					break
				}
			}
		}
	} else {
		// Generate from user input
		availableDbAdmins := available.DbAdminServices
		items := TemplateListToLabelList(availableDbAdmins)
		itemsPreselected := TemplateListToPreselectedLabelList(availableDbAdmins, selected)
		templateSelections := MultiSelectMenuQuestionIndex("Which db admin services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			Pel()
			// Get selected template config
			selectedConfig := available.DbAdminServices[index]
			// Ask questions to the user
			AskTemplateQuestions(project, &selectedConfig)
			// Ask volume questions to the user
			AskForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.DbAdminServices = append(selected.DbAdminServices, selectedConfig)
		}
	}
}
