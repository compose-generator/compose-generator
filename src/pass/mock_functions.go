package pass

import (
	"compose-generator/util"

	"github.com/compose-generator/diu"
)

// Function list for mocking
var TextQuestionWithDefault = util.TextQuestionWithDefault
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
