/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package project

import (
	"compose-generator/util"
	"io/ioutil"
	"os"

	"github.com/compose-spec/compose-go/loader"
)

var logWarning = util.LogWarning
var logError = util.LogError
var warningLogger = util.WarningLogger
var errorLogger = util.ErrorLogger
var remove = os.Remove
var removeAll = os.RemoveAll
var normalizePaths = util.NormalizePaths
var fileExists = util.FileExists
var readFile = ioutil.ReadFile
var loadComposition = loader.Load
var parseCompositionYAML = loader.ParseYAML
