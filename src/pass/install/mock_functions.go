/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package pass

import "compose-generator/util"

// Logging
var logError = util.LogError
var infoLogger = util.InfoLogger
var errorLogger = util.ErrorLogger

// Text output
var pl = util.Pl
var pel = util.Pel
var startProcess = util.StartProcess
var stopProcess = util.StopProcess

// File operations
var downloadFile = util.DownloadFile

// Environment
var isPrivileged = util.IsPrivileged

// Other
var executeAndWait = util.ExecuteAndWait
var executeWithOutput = util.ExecuteWithOutput
