package cmd

import (
	"compose-generator/parser"
	add_pass "compose-generator/pass/add"
	gen_pass "compose-generator/pass/generate"
	install_pass "compose-generator/pass/install"
	remove_pass "compose-generator/pass/remove"
	"compose-generator/project"
	"compose-generator/util"
	"io/ioutil"
	"path/filepath"

	"github.com/otiai10/copy"

	"github.com/docker/docker/client"
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
var newClientWithOpts = client.NewClientWithOpts
var getCustomTemplatesPath = util.GetCustomTemplatesPath
var readDir = ioutil.ReadDir
var loadProjectMetadata = project.LoadProjectMetadata
var abs = filepath.Abs
var rel = filepath.Rel
var copyDir = copy.Copy

var addBuildOrImagePass = add_pass.AddBuildOrImage
var addNamePass = add_pass.AddName
var addContainerNamePass = add_pass.AddContainerName
var addVolumesPass = add_pass.AddVolumes
var addNetworksPass = add_pass.AddNetworks
var addPortsPass = add_pass.AddPorts
var addEnvVarsPass = add_pass.AddEnvVars
var addEnvFilesPass = add_pass.AddEnvFiles
var addRestartPass = add_pass.AddRestart
var addDependsPass = add_pass.AddDepends
var addDependantsPass = add_pass.AddDependants
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
