package cmd

import (
	install_pass "compose-generator/pass/install"
	remove_pass "compose-generator/pass/remove"
	"compose-generator/util"
)

var isDockerizedEnvironment = util.IsDockerizedEnvironment
var printError = util.Error
var printSuccessMessage = util.Success
var commandExists = util.CommandExists
var getDockerVersion = util.GetDockerVersion
var pel = util.Pel
var yesNoQuestion = util.YesNoQuestion
var installDockerPass = install_pass.InstallDocker
var removeVolumesPass = remove_pass.RemoveVolumes
var removeNetworksPass = remove_pass.RemoveNetworks
var removeDependenciesPass = remove_pass.RemoveDependencies
