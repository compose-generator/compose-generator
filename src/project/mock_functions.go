package project

import (
	"compose-generator/util"
	"os"
)

var printWarning = util.Warning
var remove = os.Remove
var removeAll = os.RemoveAll
var normalizePaths = util.NormalizePaths
var fileExists = util.FileExists
