/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	addPass "compose-generator/pass/add"
	"compose-generator/project"
	"compose-generator/util"
	"io/ioutil"
	"os"

	"github.com/docker/docker/client"
	"github.com/fatih/color"
	"github.com/otiai10/copy"
	"github.com/sethvargo/go-password/password"
	"gopkg.in/yaml.v2"
)

// Logging
var logError = util.LogError
var logWarning = util.LogWarning
var infoLogger = util.InfoLogger
var warningLogger = util.WarningLogger
var errorLogger = util.ErrorLogger

// Text output
var p = util.P
var pl = util.Pl
var pel = util.Pel
var heading = util.Heading
var textQuestion = util.TextQuestion
var yesNoQuestion = util.YesNoQuestion
var multiSelectMenuQuestionIndex = util.MultiSelectMenuQuestionIndex
var startProcess = util.StartProcess
var stopProcess = util.StopProcess
var printSecretValue = color.Yellow
var askTemplateQuestions = util.AskTemplateQuestions
var askTemplateProxyQuestions = util.AskTemplateProxyQuestions
var askForCustomVolumePaths = util.AskForCustomVolumePaths
var evaluateConditionalSections = util.EvaluateConditionalSections

// File operations
var fileExists = util.FileExists
var getPredefinedServicesPath = util.GetPredefinedServicesPath
var mkdirAll = os.MkdirAll
var copyFile = copy.Copy
var openFile = os.Open
var readAllFromFile = ioutil.ReadAll
var readFile = ioutil.ReadFile
var writeFile = ioutil.WriteFile
var unmarshalYaml = yaml.Unmarshal
var loadTemplateService = project.LoadTemplateService

// Passes
var addBuildOrImagePass = addPass.AddBuildOrImage
var addNamePass = addPass.AddName
var addContainerNamePass = addPass.AddContainerName
var addVolumesPass = addPass.AddVolumes
var addNetworksPass = addPass.AddNetworks
var addPortsPass = addPass.AddPorts
var addEnvVarsPass = addPass.AddEnvVars
var addEnvFilesPass = addPass.AddEnvFiles
var addRestartPass = addPass.AddRestart
var addDependsPass = addPass.AddDepends
var addDependantsPass = addPass.AddDependants

// Other
var newClientWithOpts = client.NewClientWithOpts
var executeOnToolbox = util.ExecuteOnToolbox
var templateListToLabelList = util.TemplateListToLabelList
var templateListToPreselectedLabelList = util.TemplateListToPreselectedLabelList
var generatePassword = password.Generate
var sliceContainsString = util.SliceContainsString
var getServiceConfigurationsByType = func(config *model.GenerateConfig, templateType string) []model.ServiceConfig {
	return config.GetServiceConfigurationsByType(templateType)
}
