/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"compose-generator/model"
	"encoding/json"
	"os/exec"
	"strings"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// EvaluateConditionalSections evaluates conditional sections in template data
func EvaluateConditionalSections(
	filePath string,
	selected *model.SelectedTemplates,
	varMap model.Vars,
) {
	dataString := prepareInputData(selected, varMap)
	// Execute CCom
	// #nosec G204
	cmd := exec.Command("ccom", "-s", "-f", "-o", filePath, "-d", dataString, filePath)
	if output, err := cmd.CombinedOutput(); err != nil {
		if strings.TrimRight(string(output), "\r\n") != "Unknown lang" {
			ErrorLogger.Println("Could not execute CCom: " + string(output) + ": " + err.Error())
			logError("Could not execute CCom: "+string(output), true)
		}
	}
}

// EvaluateConditionalSectionsToString evaluates conditional sections in template data
func EvaluateConditionalSectionsToString(
	input string,
	selected *model.SelectedTemplates,
	varMap model.Vars,
) string {
	dataString := prepareInputData(selected, varMap)
	// Execute CCom
	ccomcPath := getCComCompilerPath()
	//cmd := exec.Command("ccom", "-l", "yml", "-d", dataString, "-s", input)
	// #nosec G204
	cmd := exec.Command(ccomcPath, "false", input, dataString, "#", "", "")
	output, err := cmd.CombinedOutput()
	if err != nil {
		ErrorLogger.Println("Could not execute CCom: " + string(output) + ": " + err.Error())
		logError("Could not execute CCom: "+string(output), true)
	}
	return strings.TrimRight(string(output), "\r\n")
}

// EvaluateCondition evaluates the given condition to a boolean result
func EvaluateCondition(
	condition string,
	selected *model.SelectedTemplates,
	varMap model.Vars,
) bool {
	// Cancel if condition is 'true' or 'false'
	if condition == "true" || condition == "false" {
		return condition == "true"
	}
	// Prepare data input for CCom
	dataString := prepareInputData(selected, varMap)
	// Execute CCom
	ccomcPath := getCComCompilerPath()
	//cmd := exec.Command("ccom", "-m", "-s", "-d", dataString, condition)
	// #nosec G204
	cmd := exec.Command(ccomcPath, "true", condition, dataString, "", "", "")
	output, err := cmd.CombinedOutput()
	if err != nil {
		WarningLogger.Println("CCom returned with an error: " + string(output) + ": " + err.Error())
		logWarning("CCom returned with an error: " + string(output))
	}
	return strings.TrimRight(string(output), "\r\n") == "true"
}

// EnsureCComIsInstalled checks if CCom is present on the current machine
func EnsureCComIsInstalled() {
	if !commandExists("ccom") {
		ErrorLogger.Println("CCom installation could not be found")
		logError("CCom could not be found on your system. Please go to https://github.com/compose-generator/compose-generator/releases/latest to download the latest version.", true)
	}
}

// EnsureDockerIsRunning checks if Docker is running
func EnsureDockerIsRunning() {
	if !isDockerRunning() {
		ErrorLogger.Println("Docker engine seems to be down")
		logError("Docker engine is not running. Please start it and execute Compose Generator again.", true)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func prepareInputData(
	selected *model.SelectedTemplates,
	varMap model.Vars,
) string {
	// Create data object
	data := model.CComDataInput{
		Services: *selected,
		Var:      varMap,
	}
	// Marshal to json
	dataJson, err := json.Marshal(data)
	if err != nil {
		ErrorLogger.Println("Could not evaluate conditional sections in template: " + err.Error())
		logError("Could not evaluate conditional sections in template. Input could be corrupted", true)
	}
	return string(dataJson)
}
