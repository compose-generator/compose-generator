/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/parser"
	genPass "compose-generator/pass/generate"
	installPass "compose-generator/pass/install"
	removePass "compose-generator/pass/remove"
	"compose-generator/project"
	"compose-generator/util"
	"io/ioutil"
	"path/filepath"

	"github.com/otiai10/copy"
)

var isDockerizedEnvironment = util.IsDockerizedEnvironment
var printError = util.Error
var printWarning = util.Warning
var printSuccess = util.Success
var printHeading = util.Heading
var printSuccessMessage = util.Success
var commandExists = util.CommandExists
var getDockerVersion = util.GetDockerVersion
var pl = util.Pl
var pel = util.Pel
var yesNoQuestion = util.YesNoQuestion
var textQuestionWithDefault = util.TextQuestionWithDefault
var menuQuestionIndex = util.MenuQuestionIndex
var clearScreen = util.ClearScreen
var startProcess = util.StartProcess
var stopProcess = util.StopProcess
var getAvailablePredefinedTemplates = parser.GetAvailablePredefinedTemplates
var getCustomTemplatesPath = util.GetCustomTemplatesPath
var readDir = ioutil.ReadDir
var loadProjectMetadata = project.LoadProjectMetadata
var abs = filepath.Abs
var rel = filepath.Rel
var copyDir = copy.Copy
var fileExists = util.FileExists
var normalizePaths = util.NormalizePaths
var installDockerPass = installPass.InstallDocker
var removeVolumesPass = removePass.RemoveVolumes
var removeNetworksPass = removePass.RemoveNetworks
var removeDependenciesPass = removePass.RemoveDependencies
var generateChooseProxiesPass = genPass.GenerateChooseProxies
var generateChooseTlsHelpersPass = genPass.GenerateChooseTlsHelpers
var generateChooseFrontendsPass = genPass.GenerateChooseFrontends
var generateChooseBackendsPass = genPass.GenerateChooseBackends
var generateChooseDatabasesPass = genPass.GenerateChooseDatabases
var generateChooseDbAdminsPass = genPass.GenerateChooseDbAdmins
var generatePass = genPass.Generate
var generateResolveDependencyGroupsPass = genPass.GenerateResolveDependencyGroups
var generateSecretsPass = genPass.GenerateSecrets
var generateAddProfilesPass = genPass.GenerateAddProfiles
var generateAddProxyNetworks = genPass.GenerateAddProxyNetworks
var generateCopyVolumesPass = genPass.GenerateCopyVolumes
var generateReplaceVarsInConfigFilesPass = genPass.GenerateReplacePlaceholdersInConfigFiles
var generateExecServiceInitCommandsPass = genPass.GenerateExecServiceInitCommands
var generateExecDemoAppInitCommandsPass = genPass.GenerateExecDemoAppInitCommands
