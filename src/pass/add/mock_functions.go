/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package pass

import (
	commonPass "compose-generator/pass/common"
	"compose-generator/util"

	"github.com/compose-generator/diu"
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
var multiSelectMenuQuestion = util.MultiSelectMenuQuestion
var printError = util.Error
var printWarning = util.Warning
var getImageManifest = diu.GetImageManifest
var p = util.P
var pel = util.Pel
var success = util.Success
var fileExists = util.FileExists
var isDir = util.IsDir
var visitServiceDependencies = commonPass.VisitServiceDependencies
