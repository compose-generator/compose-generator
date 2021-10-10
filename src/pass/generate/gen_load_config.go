/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package pass

import (
	"compose-generator/model"
	"strings"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadGenerateConfig loads a generate configuration from a file
func LoadGenerateConfig(project *model.CGProject, config *model.GenerateConfig, configPath string) {
	if configPath == "" {
		// Welcome Message
		heading("Welcome to Compose Generator! 👋")
		pl("Please continue by answering a few questions:")
		pel()

		// Ask the user for the config information
		config.ProjectName = textQuestion("What is the name of your project:")
		if config.ProjectName == "" {
			printError("You must specify a project name!", nil, true)
			return
		}
		config.ProductionReady = yesNoQuestion("Do you want the output to be production-ready?", false)
		config.FromFile = false
	} else {
		// Take the given config file and load config from there
		if fileExists(configPath) {
			yamlFile, err := openFile(configPath)
			if err != nil {
				printError("Could not load config file. Permissions granted?", err, true)
				return
			}
			content, err := readAllFromFile(yamlFile)
			if err != nil {
				printError("Could not load config file. Permissions granted?", err, true)
				return
			}
			// Parse yaml
			if err := unmarshalYaml(content, &config); err != nil {
				printError("Could not unmarshal config file", err, true)
				return
			}
			config.FromFile = true
		} else {
			printError("Config file could not be found", nil, true)
			return
		}
	}

	project.Name = config.ProjectName
	project.ProductionReady = config.ProductionReady
	project.ContainerName = strings.ReplaceAll(strings.ToLower(project.Name), " ", "-")
	if project.Vars == nil {
		project.Vars = make(map[string]string)
	}
	project.Vars["PROJECT_NAME"] = project.Name
	project.Vars["PROJECT_NAME_CONTAINER"] = project.ContainerName
}
