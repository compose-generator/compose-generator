package pass

import "compose-generator/util"

// Function list for mocking
var printError = util.Error
var startProcess = util.StartProcess
var stopProcess = util.StopProcess
var isPrivileged = util.IsPrivileged
var executeAndWait = util.ExecuteAndWait
var executeWithOutput = util.ExecuteWithOutput
var downloadFile = util.DownloadFile
