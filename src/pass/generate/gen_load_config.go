/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var loadConfigFromUrlMockable = loadConfigFromUrl
var loadConfigFromFileMockable = loadConfigFromFile

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadGenerateConfig loads a generate configuration from a file
func LoadGenerateConfig(project *model.CGProject, config *model.GenerateConfig, configInput string) {
	if configInput == "" {
		// Welcome Message
		heading("Welcome to Compose Generator! 👋")
		pl("Please continue by answering a few questions:")
		pel()

		// Ask the user for the config information
		config.ProjectName = textQuestion("What is the name of your project:")
		if config.ProjectName == "" {
			errorLogger.Println("No project name specified")
			logError("You must specify a project name!", true)
			return
		}
		infoLogger.Println("Project name: '" + config.ProjectName + "'")
		config.ProductionReady = yesNoQuestion("Do you want the output to be production-ready?", false)
		config.FromFile = false
		infoLogger.Println("Production-ready: '" + strconv.FormatBool(config.ProductionReady) + "'")
	} else {
		// Check if the input is an url
		if isUrl(configInput) {
			loadConfigFromUrlMockable(config, configInput)
		} else {
			loadConfigFromFileMockable(config, configInput)
		}
		config.FromFile = true
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

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func loadConfigFromFile(config *model.GenerateConfig, configPath string) {
	// Take the given config file and load config from there
	if fileExists(configPath) {
		infoLogger.Println("Config file was attached")
		yamlFile, err := openFile(configPath)
		if err != nil {
			errorLogger.Println("Could not load config file: " + err.Error())
			logError("Could not load config file. Permissions granted?", true)
			return
		}
		content, err := readAllFromFile(yamlFile)
		if err != nil {
			errorLogger.Println("Could not load config file: " + err.Error())
			logError("Could not load config file. Permissions granted?", true)
			return
		}
		// Parse yaml
		if err := unmarshalYaml(content, &config); err != nil {
			errorLogger.Println("Could not unmarshal config file: " + err.Error())
			logError("Could not unmarshal config file", true)
			return
		}
	} else {
		errorLogger.Println("Config file could not be found")
		logError("Config file could not be found", true)
	}
}

func loadConfigFromUrl(config *model.GenerateConfig, configUrl string) {
	// Make web request
	response, err := http.Get(configUrl)
	if err != nil {
		errorLogger.Println("Config url could not be read")
		logError("Config url could not be read", true)
	}
	defer response.Body.Close()
	// Read response
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorLogger.Println("Config not parse yaml")
		logError("Config not parse yaml", true)
	}
	// Parse yaml
	if err := unmarshalYaml(bytes, &config); err != nil {
		errorLogger.Println("Could not unmarshal config file: " + err.Error())
		logError("Could not unmarshal config file", true)
	}
}
