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
	varMap map[string]string,
) string {
	dataString := prepareInputData(selected, varMap)
	// Execute CCom
	// #nosec G204
	cmd := exec.Command("ccom", "-l", "yml", "-d", dataString, "-s", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		Error("Could not execute CCom", err, true)
	}
	return strings.TrimRight(string(output), "\r\n")
}

// EvaluateCondition evaluates the given condition to a boolean result
func EvaluateCondition(
	condition string,
	selected *model.SelectedTemplates,
	varMap map[string]string,
) bool {
	dataString := prepareInputData(selected, varMap)
	// Execute CCom
	// #nosec G204
	cmd := exec.Command("ccom", "-m", "-s", "-d", dataString, condition)
	output, err := cmd.CombinedOutput()
	if err != nil {
		Warning("CCom returned with an error: " + string(output))
	}
	return strings.TrimRight(string(output), "\r\n") == "true"
}

// EnsureCComIsInstalled checks if CCom is present on the current machine
func EnsureCComIsInstalled() {
	if !CommandExists("ccom") {
		Error("CCom could not be found on your system. Please go to https://github.com/compose-generator/compose-generator/releases/latest to download the latest version.", nil, true)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func prepareInputData(
	selected *model.SelectedTemplates,
	varMap map[string]string,
) string {
	// Create data object
	data := model.CComDataInput{
		Services: *selected,
		Var:      varMap,
	}
	// Marshal to json
	dataJson, err := json.Marshal(data)
	if err != nil {
		Error("Could not evaluate conditional sections in template. Input could be corrupted", err, true)
	}
	return string(dataJson)
}
