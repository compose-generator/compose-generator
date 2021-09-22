package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"strings"
)

var replaceVarsInFileMockable = replaceVarsInFile

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateReplaceVarsInConfigFiles replaces all variables in the config files, stated in the selected templates
func GenerateReplaceVarsInConfigFiles(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	for _, template := range selectedTemplates.GetAll() {
		// Replace vars for all config files in this template
		spinner := startProcess("Applying custom config for " + template.Label + " ...")
		for _, configFile := range template.GetFilePathsByType(model.FileTypeConfig) {
			filePath := project.Composition.WorkingDir + "/" + util.ReplaceVarsInString(configFile, project.Vars)
			// Replace all vars in file if it exists
			replaceVarsInFileMockable(filePath, project.Vars)
		}
		stopProcess(spinner)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func replaceVarsInFile(filePath string, vars map[string]string) {
	if fileExists(filePath) {
		// Read contents from file
		content, err := readFile(filePath)
		if err != nil {
			util.Error("Unable to parse .gitignore file", err, true)
		}
		contentStr := string(content)
		// Replace all vars
		for name, value := range vars {
			contentStr = strings.ReplaceAll(contentStr, "${{"+name+"}}", value)
		}
		// Write contents back
		if err := writeFile(filePath, []byte(contentStr), 0600); err != nil {
			util.Error("Unable to write config file '"+filePath+"' back to the disk", err, true)
		}
	}
}
