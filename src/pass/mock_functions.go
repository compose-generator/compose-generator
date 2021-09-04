package pass

import (
	"compose-generator/model"
	"compose-generator/util"

	"github.com/compose-generator/diu"
)

// Function list for mocking
var TextQuestion = util.TextQuestion
var TextQuestionWithDefault = util.TextQuestionWithDefault
var TextQuestionWithSuggestions = util.TextQuestionWithSuggestions
var TextQuestionWithDefaultAndSuggestions = util.TextQuestionWithDefaultAndSuggestions
var TextQuestionWithValidator = util.TextQuestionWithValidator
var YesNoQuestion = util.YesNoQuestion
var MenuQuestion = util.MenuQuestion
var MenuQuestionIndex = util.MenuQuestionIndex
var MultiSelectMenuQuestionIndex = util.MultiSelectMenuQuestionIndex
var MultiSelectMenuQuestion = util.MultiSelectMenuQuestion
var Error = util.Error
var GetImageManifest = diu.GetImageManifest
var P = util.P
var Pl = util.Pl
var Pel = util.Pel
var Done = util.Done
var Success = util.Success
var FileExists = util.FileExists
var IsDir = util.IsDir
var ExecuteOnLinux = util.ExecuteOnLinux
var IsPrivileged = util.IsPrivileged
var ExecuteAndWait = util.ExecuteAndWait
var ExecuteWithOutput = util.ExecuteWithOutput
var DownloadFile = util.DownloadFile
var TemplateListToLabelList = util.TemplateListToLabelList
var TemplateListToPreselectedLabelList = util.TemplateListToPreselectedLabelList
var AskTemplateQuestions = util.AskTemplateQuestions
var AskForCustomVolumePaths = util.AskForCustomVolumePaths
var GetServiceConfigurationsByName = func(config *model.GenerateConfig, templateType string) []model.ServiceConfig {
	return config.GetServiceConfigurationsByName(templateType)
}
