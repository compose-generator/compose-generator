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
	// Execute CCom
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
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) string {
	// Delete empty service categories in template data
	for key, value := range templateData {
		if len(value) == 0 {
			delete(templateData, key)
		}
	}
	// Re-map db-admin to dbadmin and tls-helper to tlshelper
	if val, ok := templateData["db-admin"]; ok {
		templateData["dbadmin"] = val
		delete(templateData, "db-admin")
	}
	if val, ok := templateData["tls-helper"]; ok {
		templateData["tlshelper"] = val
		delete(templateData, "tls-helper")
	}
	// Create data object
	data := model.CComDataInput{
		Services: templateData,
		Var:      varMap,
	}
	// Marshal to json
	dataJson, err := json.Marshal(data)
	if err != nil {
		Error("Could not evaluate conditional sections in template. Could be corrupted", err, true)
	}
	return string(dataJson)
}
