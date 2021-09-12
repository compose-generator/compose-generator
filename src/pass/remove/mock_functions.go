package pass

import (
	"compose-generator/util"
	"os"
)

// Function list for mocking
var yesNoQuestion = util.YesNoQuestion
var printWarning = util.Warning
var fileExists = util.FileExists
var removeAll = os.RemoveAll