/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/model"
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

// PredefinedTemplateNewCliFlags are the cli flags for the remove command
var PredefinedTemplateNewCliFlags = []cli.Flag{}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// NewPredefinedTemplate generates a blank predefined service template in the respective directory
func NewPredefinedTemplate(context *cli.Context) error {
	infoLogger.Println("NewPredefinedTemplate command executed")

	if !isDevVersion() {
		errorLogger.Println("Executed NewPredefinedTemplate command without being dev. Aborting.")
		logError("Cannot run this command in a production environment.", true)
		return nil
	}

	// Ask user questions and save to disk subsequently
	config, service, readme := createNewPredefinedTemplate()
	savePredefinedTemplate(config, service, readme)

	return nil
}

const serviceTemplate = `image: hello-world:${{%s_VERSION}}
container_name: ${{PROJECT_NAME_CONTAINER}}-%s-%s
%snetworks:
# ToDo: Insert or remove section
ports:
# ToDo: Insert or remove section
volumes:
# ToDo: Insert or remove section
env_file:
# ToDo: Insert or delete section`

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func createNewPredefinedTemplate() (*model.PredefinedTemplateConfig, string, string) {
	config := model.PredefinedTemplateConfig{}

	// Ask for template label
	config.Label = textQuestion("Template label:")
	name := strings.ReplaceAll(strings.ToLower(config.Label), " ", "-")
	snakeUpperName := strings.ToUpper(strings.ReplaceAll(name, "-", "_"))

	// Ask for template type
	config.Type = menuQuestion("What is the closest match of specifying the type?", []string{
		model.TemplateTypeFrontend,
		model.TemplateTypeBackend,
		model.TemplateTypeDatabase,
		model.TemplateTypeDbAdmin,
	})

	// Set default configuration
	config.Preselected = "false"
	config.Proxied = config.Type == model.TemplateTypeFrontend || config.Type == model.TemplateTypeDbAdmin
	config.Files = []model.File{
		{
			Path: "service.yml",
			Type: "service",
		},
		{
			Path: "README.md",
			Type: "docs",
		},
	}
	config.Questions = []model.Question{
		{
			Text:         "Which version of " + config.Label + " do you want to use?",
			Type:         model.QuestionTypeText,
			DefaultValue: "latest",
			Variable:     snakeUpperName + "_VERSION",
		},
	}

	// Check if dir already exists
	config.Dir = getPredefinedServicesPath() + "/" + config.Type + "/" + name
	if fileExists(config.Dir) {
		errorLogger.Println("Predefined template dir '" + config.Dir + "' already exists. Aborting.")
		logError("Template dir already exists. Aborting.", true)
		return nil, "", ""
	}

	// Prepare Readme contents
	readme := fmt.Sprintf("## %s\nToDo: Insert software description here.\n\n### Setup\nToDo: Insert setup instructions here.", config.Label)

	// Prepare service contents
	restartRule := ""
	if config.Type != model.TemplateTypeDbAdmin {
		restartRule = "restart: always\n"
	}
	service := fmt.Sprintf(serviceTemplate, snakeUpperName, config.Type, name, restartRule)

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
	file, err := marshalIndent(config, "", "    ")
	if err != nil {
		errorLogger.Println("Error marshaling predefined template config to json: " + err.Error())
		logError("Error marshaling config to json", true)
		return
	}
	if err = writeFile(config.Dir+"/config.json", file, 0600); err != nil {
		errorLogger.Println("Error writing predefined template config to file: " + err.Error())
		logError("Error saving config.json", true)
		return
	}

	// Save service.yml
	if err = writeFile(config.Dir+"/service.yml", []byte(service), 0600); err != nil {
		errorLogger.Println("Error writing service to file: " + err.Error())
		logError("Error saving service.yml", true)
		return
	}

	// Save README.md
	if err = writeFile(config.Dir+"/README.md", []byte(readme), 0600); err != nil {
		errorLogger.Println("Error writing Readme to file: " + err.Error())
		logError("Error saving README.md", true)
		return
	}
	stopProcess(spinner)
}
