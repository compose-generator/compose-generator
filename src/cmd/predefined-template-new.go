package cmd

import (
	"compose-generator/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/urfave/cli/v2"
)

// PredefinedTemplateNewCliFlags are the cli flags for the remove command
var PredefinedTemplateNewCliFlags = []cli.Flag{}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// NewPredefinedTemplate generates a blank predefined service template in the respective directory
func NewPredefinedTemplate(context *cli.Context) error {
	infoLogger.Println("NewPredefinedTemplate command executed")

	config, service, readme := createNewPredefinedTemplate()
	savePredefinedTemplate(config, service, readme)

	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func createNewPredefinedTemplate() (*model.PredefinedTemplateConfig, string, string) {
	config := model.PredefinedTemplateConfig{}

	// Ask for template label
	config.Label = textQuestion("Template label:")
	config.Name = strings.ReplaceAll(strings.ToLower(config.Label), " ", "-")

	// Ask for template type
	config.Type = menuQuestion("", []string{"frontend", "backend", "database", "db-admin"})
	config.Dir = GetPredefinedServicesPath() + "/" + config.Type + "/" + config.Name

	// Prepare Readme contents
	readme := fmt.Sprintf("## %s\nToDo: Insert software description here.\n\n### Setup\nToDo: Insert setup instructions here.", config.Label)

	service := ""

	return &config, service, readme
}

func savePredefinedTemplate(config *model.PredefinedTemplateConfig, service, readme string) {
	spinner := startProcess("Saving predefined service template ...")
	// Save config.json
	file, err := json.Marshal(config)
	if err != nil {
		errorLogger.Println("Error marshalling predefined template config to json: " + err.Error())
		logError("Error marshalling config to json", true)
		return
	}
	if err = ioutil.WriteFile("test.json", file, 0600); err != nil {
		errorLogger.Println("Error writing predefined template config to file: " + err.Error())
		logError("Error saving config.json", true)
		return
	}

	// Save service.yml
	if err = ioutil.WriteFile("service.yml", []byte(service), 0600); err != nil {
		errorLogger.Println("Error writing service to file: " + err.Error())
		logError("Error saving service.yml", true)
		return
	}

	// Save README.md
	if err = ioutil.WriteFile("README.md", []byte(readme), 0600); err != nil {
		errorLogger.Println("Error writing Readme to file: " + err.Error())
		logError("Error saving README.md", true)
		return
	}
	stopProcess(spinner)
}
