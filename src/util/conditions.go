package util

import (
	"compose-generator/model"
	"encoding/json"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

func EvaluateConditionalSections(
	filePath string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) string {
	dataString := PrepareInputData(templateData, varMap)
	Pl(dataString)
	// Execute CCom
	// ToDo: Remove '-l yml' as soon as CCom #62 is fixed
	return ExecuteAndWaitWithOutput("ccom", "-l", "yml", "-d", dataString, "-s", filePath)
}

// EvaluateCondition evaluates the given condition to a boolean result
func EvaluateCondition(
	condition string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) bool {
	dataString := PrepareInputData(templateData, varMap)
	// Execute CCom
	result := ExecuteAndWaitWithOutput("ccom", "-m", "-s", "-d", dataString, condition)
	return result == "true"
}

func CheckIfCComIsInstalled() {
	if !CommandExists("ccom") {
		Error("CCom is not missing on your system. Please go to https://github.com/compose-generator/compose-generator/releases/latest to download the latest version.", nil, true)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func PrepareInputData(
	selectedTemplateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) string {
	// Delete empty service categories in template data
	for key, value := range selectedTemplateData {
		if len(value) == 0 {
			delete(selectedTemplateData, key)
		}
	}
	// Re-map db-admin to dbadmin and tls-helper to tlshelper
	if val, ok := selectedTemplateData["db-admin"]; ok {
		selectedTemplateData["dbadmin"] = val
		delete(selectedTemplateData, "db-admin")
	}
	if val, ok := selectedTemplateData["tls-helper"]; ok {
		selectedTemplateData["tlshelper"] = val
		delete(selectedTemplateData, "tls-helper")
	}
	// Create data object
	data := model.CComDataInput{
		Services: selectedTemplateData,
		Var:      varMap,
	}
	// Marshal to json
	dataJson, err := json.Marshal(data)
	if err != nil {
		Error("Could not evaluate conditional sections in template. Could be corrupted", err, true)
	}
	return string(dataJson)
}
