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
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

// Logging
var logError = util.LogError
var logWarning = util.LogWarning
var infoLogger = util.InfoLogger
var warningLogger = util.WarningLogger
var errorLogger = util.ErrorLogger

// Text output
var printSuccess = util.Success
var printHeading = util.Heading
var pl = util.Pl
var pel = util.Pel
var yesNoQuestion = util.YesNoQuestion
var textQuestion = util.TextQuestion
var textQuestionWithDefault = util.TextQuestionWithDefault
var menuQuestion = util.MenuQuestion
var menuQuestionIndex = util.MenuQuestionIndex
var clearScreen = util.ClearScreen
var startProcess = util.StartProcess
var stopProcess = util.StopProcess
var marshalIndent = json.MarshalIndent

// File operations
var readDir = ioutil.ReadDir
var mkDir = os.Mkdir
var abs = filepath.Abs
var rel = filepath.Rel
var copyDir = copy.Copy
var fileExists = util.FileExists
var loadProjectMetadata = project.LoadProjectMetadata
var normalizePaths = util.NormalizePaths
var removeAll = os.RemoveAll
var writeFile = ioutil.WriteFile

// Environment
var isDockerizedEnvironment = util.IsDockerizedEnvironment
var isDevVersion = util.IsDevVersion
var commandExists = util.CommandExists
var getDockerVersion = util.GetDockerVersion
var getAvailablePredefinedTemplates = parser.GetAvailablePredefinedTemplates
var getPredefinedServicesPath = util.GetPredefinedServicesPath
var getCustomTemplatesPath = util.GetCustomTemplatesPath

// Passes
var generateChooseProxiesPass = genPass.GenerateChooseProxies
var generateChooseTlsHelpersPass = genPass.GenerateChooseTlsHelpers
var generateChooseFrontendsPass = genPass.GenerateChooseFrontends
var generateChooseBackendsPass = genPass.GenerateChooseBackends
var generateChooseDatabasesPass = genPass.GenerateChooseDatabases
var generateChooseDbAdminsPass = genPass.GenerateChooseDbAdmins
var generateAddWatchtowerPass = genPass.GenerateAddWatchtower
var generatePass = genPass.Generate
var generateResolveDependencyGroupsPass = genPass.GenerateResolveDependencyGroups
var generateSecretsPass = genPass.GenerateSecrets
var generateAddProfilesPass = genPass.GenerateAddProfiles
var generateAddProxyNetworks = genPass.GenerateAddProxyNetworks
var generateCopyVolumesPass = genPass.GenerateCopyVolumes
var generateReplaceVarsInConfigFilesPass = genPass.GenerateReplacePlaceholdersInConfigFiles
var generateExecServiceInitCommandsPass = genPass.GenerateExecServiceInitCommands
var generateExecDemoAppInitCommandsPass = genPass.GenerateExecDemoAppInitCommands
var removeVolumesPass = removePass.RemoveVolumes
var removeNetworksPass = removePass.RemoveNetworks
var removeDependenciesPass = removePass.RemoveDependencies
var installDockerPass = installPass.InstallDocker
