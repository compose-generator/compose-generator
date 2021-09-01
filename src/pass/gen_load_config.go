package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadGenerateConfig loads a generate configuration from a file
func LoadGenerateConfig(project *model.CGProject, config *model.GenerateConfig, configPath string) {
	if configPath == "" {
		// Welcome Message
		util.Heading("Welcome to Compose Generator! 👋")
		util.Pl("Please continue by answering a few questions:")
		util.Pel()

		// Ask the user for the config information
		config.ProjectName = util.TextQuestion("What is the name of your project:")
		if config.ProjectName == "" {
			util.Error("You must specify a project name!", nil, true)
		}
		config.ProductionReady = util.YesNoQuestion("Do you want the output to be production-ready?", false)
		config.FromFile = false
	} else {
		// Take the given config file and load config from there
		if util.FileExists(configPath) {
			yamlFile, err1 := os.Open(configPath)
			content, err2 := ioutil.ReadAll(yamlFile)
			if err1 != nil {
				util.Error("Could not load config file. Permissions granted?", err1, true)
			}
			if err2 != nil {
				util.Error("Could not load config file. Permissions granted?", err2, true)
			}
			// Parse yaml
			yaml.Unmarshal(content, &config)
			config.FromFile = true
		} else {
			util.Error("Config file could not be found", nil, true)
		}
	}

	project.Name = config.ProjectName
	project.ProductionReady = config.ProductionReady
	project.ContainerName = strings.ReplaceAll(strings.ToLower(project.Name), " ", "-")
	project.Vars["PROJECT_NAME"] = project.Name
	project.Vars["PROJECT_NAME_CONTAINER"] = project.ContainerName
}
