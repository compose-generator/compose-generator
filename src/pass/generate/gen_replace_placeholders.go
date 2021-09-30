package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"strings"
)

var replaceVarsInFileMockable = replaceVarsInFile
var replaceConditionalSectionsInFileMockable = replaceConditionalSectionsInFile

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateReplacePlaceholdersInConfigFiles replaces all variables in the config files, stated in the selected templates
func GenerateReplacePlaceholdersInConfigFiles(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	for _, template := range selectedTemplates.GetAll() {
		// Replace vars for all config files in this template
		spinner := startProcess("Applying custom config for " + template.Label + " ...")
		for _, configFile := range template.GetFilePathsByType(model.FileTypeConfig) {
			filePath := project.Composition.WorkingDir + "/" + util.ReplaceVarsInString(configFile, project.Vars)
			if fileExists(filePath) {
				// Replace all vars in file if it exists
				replaceVarsInFileMockable(filePath, project.Vars)
				// Replace conditional sections in file if it exists
				replaceConditionalSectionsInFileMockable(filePath, selectedTemplates, project.Vars)
			}
		}
		stopProcess(spinner)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func replaceVarsInFile(filePath string, vars map[string]string) {
	// Read contents from file
	content, err := readFile(filePath)
	if err != nil {
		printError("Unable to read config file '"+filePath+"'", err, false)
		return
	}
	contentStr := string(content)
	// Replace all vars
	for name, value := range vars {
		contentStr = strings.ReplaceAll(contentStr, "${{"+name+"}}", value)
	}
	// Write contents back
	if err := writeFile(filePath, []byte(contentStr), 0600); err != nil {
		printError("Unable to write config file '"+filePath+"' back to the disk", err, false)
	}
}

func replaceConditionalSectionsInFile(filePath string, selectedTemplates *model.SelectedTemplates, vars map[string]string) {
	// Evaluate conditional sections
	evaluateConditionalSections(filePath, selectedTemplates, vars)
}
