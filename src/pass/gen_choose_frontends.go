package pass

import (
	"compose-generator/model"
	"compose-generator/util"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateChooseFrontends lets the user choose predefined frontend service templates
func GenerateChooseFrontends(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates) {
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
		// Save template to the selected templates
		selected.FrontendServices = append(selected.FrontendServices, selectedConfig)
	}
}

/*
func askForStackComponent(
	templateData *map[string][]model.ServiceTemplateConfig,
	selectedTemplateData *map[string][]model.ServiceTemplateConfig,
	varMap *map[string]string,
	volMap *map[string]string,
	usedPorts *[]int,
	usedVolumes *[]string,
	component string,
	multiSelect bool,
	question string,
	flagAdvanced bool,
	flagWithDockerfile bool,
) (componentCount int) {
	templates := (*templateData)[component]
	items := util.TemplateListToLabelList(templates)
	itemsPreselected := util.TemplateListToPreselectedLabelList(templates, selectedTemplateData)
	(*selectedTemplateData)[component] = []model.ServiceTemplateConfig{}
	if multiSelect {
		templateSelections := util.MultiSelectMenuQuestionIndex(question, items, itemsPreselected)
		for _, index := range templateSelections {
			util.Pel()
			(*selectedTemplateData)[component] = append((*selectedTemplateData)[component], templates[index])
			getVarMapFromQuestions(varMap, usedPorts, templates[index].Questions, flagAdvanced)
			getVolumeMapFromVolumes(varMap, volMap, usedVolumes, templates[index], flagAdvanced, flagWithDockerfile)
			componentCount++
		}
	} else {
		templateSelection := util.MenuQuestionIndex(question, items)
		(*selectedTemplateData)[component] = append((*selectedTemplateData)[component], templates[templateSelection])
		getVarMapFromQuestions(varMap, usedPorts, templates[templateSelection].Questions, flagAdvanced)
		getVolumeMapFromVolumes(varMap, volMap, usedVolumes, templates[templateSelection], flagAdvanced, flagWithDockerfile)
		componentCount = 1
	}
	util.Pel()
	return
}
*/
