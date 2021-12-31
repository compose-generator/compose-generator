/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"path/filepath"
	"strings"
)

var replaceVarsInFileMockable = replaceVarsInFile

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateReplacePlaceholdersInConfigFiles replaces all variables in the config files, stated in the selected templates
func GenerateReplacePlaceholdersInConfigFiles(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	for _, template := range selectedTemplates.GetAll() {
		// Replace vars for all config files in this template
		infoLogger.Println("Replacing placeholders in config files for '" + template.Label + "' ...")
		spinner := startProcess("Applying custom configuration for " + template.Label + " ...")
		for _, configFile := range template.GetFilePathsByType(model.FileTypeConfig) {
			filePath := filepath.Clean(project.Composition.WorkingDir + util.ReplaceVarsInString(configFile, project.Vars))
			if fileExists(filePath) {
				// Replace all vars in file if it exists
				replaceVarsInFileMockable(filePath, project.Vars)
				// Replace conditional sections in file if it exists
				evaluateConditionalSections(filePath, selectedTemplates, project.Vars)
			}
		}
		stopProcess(spinner)
		infoLogger.Println("Replacing placeholders in config files for '" + template.Label + "' (done)")
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func replaceVarsInFile(filePath string, vars model.Vars) {
	// Read contents from file
	content, err := readFile(filePath)
	if err != nil {
		errorLogger.Println("Unable to read config file '" + filePath + "': " + err.Error())
		logError("Unable to read config file '"+filePath+"'", false)
		return
	}
	contentStr := string(content)
	// Replace all vars
	for name, value := range vars {
		contentStr = strings.ReplaceAll(contentStr, "${{"+name+"}}", value)
	}
	// Write contents back
	if err := writeFile(filePath, []byte(contentStr), 0600); err != nil {
		errorLogger.Println("Unable to write config file '" + filePath + "' back to the disk: " + err.Error())
		logError("Unable to write config file '"+filePath+"' back to the disk", false)
	}
}
