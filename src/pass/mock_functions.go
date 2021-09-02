package pass

import (
	"compose-generator/util"

	"github.com/compose-generator/diu"
)

// Function list for mocking
var TextQuestion = util.TextQuestion
var TextQuestionWithDefault = util.TextQuestionWithDefault
var TextQuestionWithDefaultAndSuggestions = util.TextQuestionWithDefaultAndSuggestions
var TextQuestionWithValidator = util.TextQuestionWithValidator
var YesNoQuestion = util.YesNoQuestion
var MenuQuestion = util.MenuQuestion
var MultiSelectMenuQuestion = util.MultiSelectMenuQuestion
var Error = util.Error
var GetImageManifest = diu.GetImageManifest
var Pel = util.Pel
var P = util.P
var Done = util.Done
var Success = util.Success
var FileExists = util.FileExists
var IsDir = util.IsDir
var ExecuteOnLinux = util.ExecuteOnLinux
