/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/util"
	"os"
	"path/filepath"
)

// Function list for mocking
var yesNoQuestion = util.YesNoQuestion
var logWarning = util.LogWarning
var fileExists = util.FileExists
var removeAll = os.RemoveAll
var abs = filepath.Abs
