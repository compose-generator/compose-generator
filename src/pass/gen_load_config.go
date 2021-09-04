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
		Heading("Welcome to Compose Generator! 👋")
		Pl("Please continue by answering a few questions:")
		Pel()

		// Ask the user for the config information
		config.ProjectName = TextQuestion("What is the name of your project:")
		if config.ProjectName == "" {
			Error("You must specify a project name!", nil, true)
		}
		config.ProductionReady = YesNoQuestion("Do you want the output to be production-ready?", false)
		config.FromFile = false
	} else {
		// Take the given config file and load config from there
		if FileExists(configPath) {
			yamlFile, err := OpenFile(configPath)
			if err != nil {
				Error("Could not load config file. Permissions granted?", err, true)
			}
			content, err := ReadAllFromFile(yamlFile)
			if err != nil {
				Error("Could not load config file. Permissions granted?", err, true)
			}
			// Parse yaml
			UnmarshalYaml(content, &config)
			config.FromFile = true
		} else {
			Error("Config file could not be found", nil, true)
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
