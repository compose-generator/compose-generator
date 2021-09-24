package cmd

import (
	"compose-generator/parser"
	gen_pass "compose-generator/pass/generate"
	install_pass "compose-generator/pass/install"
	remove_pass "compose-generator/pass/remove"
	"compose-generator/util"
)

var isDockerizedEnvironment = util.IsDockerizedEnvironment
var printError = util.Error
var printSuccess = util.Success
var printSuccessMessage = util.Success
var commandExists = util.CommandExists
var getDockerVersion = util.GetDockerVersion
var pel = util.Pel
var yesNoQuestion = util.YesNoQuestion
var clearScreen = util.ClearScreen
var startProcess = util.StartProcess
var stopProcess = util.StopProcess
var getAvailablePredefinedTemplates = parser.GetAvailablePredefinedTemplates

var installDockerPass = install_pass.InstallDocker
var removeVolumesPass = remove_pass.RemoveVolumes
var removeNetworksPass = remove_pass.RemoveNetworks
var removeDependenciesPass = remove_pass.RemoveDependencies
var generateChooseFrontendsPass = gen_pass.GenerateChooseFrontends
var generateChooseBackendsPass = gen_pass.GenerateChooseBackends
var generateChooseDatabasesPass = gen_pass.GenerateChooseDatabases
var generateChooseDbAdminsPass = gen_pass.GenerateChooseDbAdmins
var generateChooseProxiesPass = gen_pass.GenerateChooseProxies
var generateChooseTlsHelpersPass = gen_pass.GenerateChooseTlsHelpers
var generatePass = gen_pass.Generate
var generateResolveDependencyGroupsPass = gen_pass.GenerateResolveDependencyGroups
var generateSecretsPass = gen_pass.GenerateSecrets
var genAddProfilesPass = gen_pass.GenAddProfiles
var generateCopyVolumesPass = gen_pass.GenerateCopyVolumes
var generateReplaceVarsInConfigFilesPass = gen_pass.GenerateReplaceVarsInConfigFiles
var generateExecServiceInitCommandsPass = gen_pass.GenerateExecServiceInitCommands
var generateExecDemoAppInitCommandsPass = gen_pass.GenerateExecDemoAppInitCommands
