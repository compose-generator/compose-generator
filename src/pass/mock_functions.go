package pass

import (
	"compose-generator/model"
	"compose-generator/project"
	"compose-generator/util"
	"io/ioutil"
	"os"

	"github.com/compose-generator/diu"
	"github.com/fatih/color"
	"github.com/otiai10/copy"
	"github.com/sethvargo/go-password/password"
	"gopkg.in/yaml.v2"
)

// Function list for mocking
var textQuestion = util.TextQuestion
var textQuestionWithDefault = util.TextQuestionWithDefault
var textQuestionWithSuggestions = util.TextQuestionWithSuggestions
var textQuestionWithDefaultAndSuggestions = util.TextQuestionWithDefaultAndSuggestions
var textQuestionWithValidator = util.TextQuestionWithValidator
var yesNoQuestion = util.YesNoQuestion
var menuQuestion = util.MenuQuestion
var menuQuestionIndex = util.MenuQuestionIndex
var multiSelectMenuQuestionIndex = util.MultiSelectMenuQuestionIndex
var multiSelectMenuQuestion = util.MultiSelectMenuQuestion
var printError = util.Error
var printWarning = util.Warning
var getImageManifest = diu.GetImageManifest
var heading = util.Heading
var p = util.P
var pl = util.Pl
var pel = util.Pel
var startProcess = util.StartProcess
var stopProcess = util.StopProcess
var success = util.Success
var printSecretValue = color.Yellow
var fileExists = util.FileExists
var isDir = util.IsDir
var getPredefinedServicesPath = util.GetPredefinedServicesPath
var mkdirAll = os.MkdirAll
var removeAll = os.RemoveAll
var executeOnLinux = util.ExecuteOnLinux
var isPrivileged = util.IsPrivileged
var executeAndWait = util.ExecuteAndWait
var executeWithOutput = util.ExecuteWithOutput
var copyFile = copy.Copy
var downloadFile = util.DownloadFile
var templateListToLabelList = util.TemplateListToLabelList
var templateListToPreselectedLabelList = util.TemplateListToPreselectedLabelList
var askTemplateQuestions = util.AskTemplateQuestions
var askForCustomVolumePaths = util.AskForCustomVolumePaths
var unmarshalYaml = yaml.Unmarshal
var openFile = os.Open
var readAllFromFile = ioutil.ReadAll
var generatePassword = password.Generate
var loadTemplateService = project.LoadTemplateService
var sliceContainsString = util.SliceContainsString
var getServiceConfigurationsByName = func(config *model.GenerateConfig, templateType string) []model.ServiceConfig {
	return config.GetServiceConfigurationsByName(templateType)
}
