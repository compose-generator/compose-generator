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

	// Ask user questions and save to disk subsequently
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

	// Check if dir already exists
	config.Dir = GetPredefinedServicesPath() + "/" + config.Type + "/" + config.Name
	if fileExists(config.Dir) {
		errorLogger.Println("Predefined template dir '" + config.Dir + "' already exists. Aborting.")
		logError("Template dir already exists. Aborting.", true)
		return nil, "", ""
	}

	// Prepare Readme contents
	readme := fmt.Sprintf("## %s\nToDo: Insert software description here.\n\n### Setup\nToDo: Insert setup instructions here.", config.Label)

	service := "image:"

	return &config, service, readme
}

func savePredefinedTemplate(config *model.PredefinedTemplateConfig, service, readme string) {
	pel()
	spinner := startProcess("Saving predefined service template ...")
	// Create dir
	if err := mkDir(config.Dir, 0777); err != nil {
		errorLogger.Println("Error creating the template dir '" + config.Dir + "': " + err.Error())
		logError("Error creating the template dir", true)
		return
	}

	// Save config.json
	file, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		errorLogger.Println("Error marshaling predefined template config to json: " + err.Error())
		logError("Error marshaling config to json", true)
		return
	}
	if err = ioutil.WriteFile(config.Dir+"/config.json", file, 0600); err != nil {
		errorLogger.Println("Error writing predefined template config to file: " + err.Error())
		logError("Error saving config.json", true)
		return
	}

	// Save service.yml
	if err = ioutil.WriteFile(config.Dir+"/service.yml", []byte(service), 0600); err != nil {
		errorLogger.Println("Error writing service to file: " + err.Error())
		logError("Error saving service.yml", true)
		return
	}

	// Save README.md
	if err = ioutil.WriteFile(config.Dir+"/README.md", []byte(readme), 0600); err != nil {
		errorLogger.Println("Error writing Readme to file: " + err.Error())
		logError("Error saving README.md", true)
		return
	}
	stopProcess(spinner)
}
