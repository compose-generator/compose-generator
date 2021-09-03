package pass

import (
	"compose-generator/model"
	"compose-generator/util"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseTlsHelpers lets the user choose predefined tls helper service templates
func GenerateChooseTlsHelpers(
	project *model.CGProject,
	available *model.AvailableTemplates,
	selected *model.SelectedTemplates,
	config *model.GenerateConfig,
) {
	if config.FromFile {
		// Generate from config file
		selectedServiceConfigs := config.GetServiceConfigurationsByName(model.TemplateTypeTlsHelper)
		for _, template := range available.TlsHelperService {
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
					selected.TlsHelperService = append(selected.TlsHelperService, template)
					break
				}
			}
		}
	} else {
		// Generate from user input
		availableTlsHelpers := available.TlsHelperService
		items := util.TemplateListToLabelList(availableTlsHelpers)
		itemsPreselected := util.TemplateListToPreselectedLabelList(availableTlsHelpers, selected)
		templateSelections := util.MultiSelectMenuQuestionIndex("Which tls helper services do you need?", items, itemsPreselected)
		for _, index := range templateSelections {
			util.Pel()
			// Get selected template config
			selectedConfig := available.TlsHelperService[index]
			// Ask questions to the user
			util.AskTemplateQuestions(project, &selectedConfig)
			// Ask volume questions to the user
			util.AskForCustomVolumePaths(project, &selectedConfig)
			// Save template to the selected templates
			selected.TlsHelperService = append(selected.TlsHelperService, selectedConfig)
		}
	}
}