package cmd

import (
	pass "compose-generator/pass/install"
	"compose-generator/util"
)

var isDockerizedEnvironment = util.IsDockerizedEnvironment
var printError = util.Error
var printSuccessMessage = util.Success
var commandExists = util.CommandExists
var getDockerVersion = util.GetDockerVersion
var pel = util.Pel
var installDockerPass = pass.InstallDocker
