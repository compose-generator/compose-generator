package project

import (
	"compose-generator/util"
	"io/ioutil"
	"os"

	"github.com/compose-spec/compose-go/loader"
)

var printWarning = util.Warning
var printError = util.Error
var remove = os.Remove
var removeAll = os.RemoveAll
var normalizePaths = util.NormalizePaths
var fileExists = util.FileExists
var readFile = ioutil.ReadFile
var loadComposition = loader.Load
var parseCompositionYAML = loader.ParseYAML
