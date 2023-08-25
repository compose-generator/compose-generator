/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	commonPass "compose-generator/pass/common"
	"compose-generator/util"

	"github.com/compose-generator/diu"
)

// Logging
var logError = util.LogError
var logWarning = util.LogWarning
var infoLogger = util.InfoLogger
var warningLogger = util.WarningLogger
var errorLogger = util.ErrorLogger

// Text output
var p = util.P
var pel = util.Pel
var success = util.Success
var textQuestion = util.TextQuestion
var textQuestionWithDefault = util.TextQuestionWithDefault
var textQuestionWithSuggestions = util.TextQuestionWithSuggestions
var textQuestionWithDefaultAndSuggestions = util.TextQuestionWithDefaultAndSuggestions
var textQuestionWithValidator = util.TextQuestionWithValidator
var yesNoQuestion = util.YesNoQuestion
var menuQuestion = util.MenuQuestion
var menuQuestionIndex = util.MenuQuestionIndex
var multiSelectMenuQuestion = util.MultiSelectMenuQuestion

// File operations
var fileExists = util.FileExists
var isDir = util.IsDir

// Other
var getImageManifest = diu.GetImageManifest
var visitServiceDependencies = commonPass.VisitServiceDependencies
