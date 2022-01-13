/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/util"
	"os"
	"path/filepath"
)

// Logging
var logWarning = util.LogWarning
var infoLogger = util.InfoLogger

// Text output
var yesNoQuestion = util.YesNoQuestion

// File operations
var fileExists = util.FileExists
var removeAll = os.RemoveAll
var abs = filepath.Abs
