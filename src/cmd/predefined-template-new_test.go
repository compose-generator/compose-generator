package cmd

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewPredefinedTempate1(t *testing.T) { // Happy path
	// Test data
	expectedConfig := &model.PredefinedTemplateConfig{
		Label:       "ChilliDB",
		Name:        "chillidb",
		Type:        model.TemplateTypeDatabase,
		Dir:         "/usr/lib/compose-generator/predefined-services/database/chillidb",
		Preselected: "false",
		Proxied:     false,
		Files: []model.File{
			{
				Path: "service.yml",
				Type: "service",
			},
			{
				Path: "README.md",
				Type: "docs",
			},
		},
		Questions: []model.Question{
			{
				Text:         "Which version of ChilliDB do you want to use?",
				Type:         model.QuestionTypeText,
				DefaultValue: "latest",
				Variable:     "CHILLIDB_VERSION",
			},
		},
	}
	expectedService := "image: hello-world:${{CHILLIDB_VERSION}}\ncontainer_name: ${{PROJECT_NAME_CONTAINER}}-database-chillidb\nrestart: always\nnetworks:\n# ToDo: Insert or remove section\nports:\n# ToDo: Insert or remove section\nvolumes:\n# ToDo: Insert or remove section\nenv_file:\n# ToDo: Insert or delete section"
	expectedReadme := "## ChilliDB\nToDo: Insert software description here.\n\n### Setup\nToDo: Insert setup instructions here."
	predefinedServicesPath := "/usr/lib/compose-generator/predefined-services"
	// Mock functions
	textQuestion = func(question string) string {
		assert.Equal(t, "Template label:", question)
		return "ChilliDB"
	}
	menuQuestion = func(label string, items []string) string {
		assert.Equal(t, "What is the closest match of specifying the type?", label)
		assert.Equal(t, []string{model.TemplateTypeFrontend, model.TemplateTypeBackend, model.TemplateTypeDatabase, model.TemplateTypeDbAdmin}, items)
		return model.TemplateTypeDatabase
	}
	getPredefinedServicesPath = func() string {
		return predefinedServicesPath
	}
	fileExists = func(path string) bool {
		assert.Equal(t, predefinedServicesPath+"/database/chillidb", path)
		return false
	}
	logError = func(message string, exit bool) {
		assert.Fail(t, "Unexpected call of logError")
	}
	// Execute test
	config, service, readme := createNewPredefinedTemplate()
	// Assert
	assert.Equal(t, expectedConfig, config)
	assert.Equal(t, expectedService, service)
	assert.Equal(t, expectedReadme, readme)
}
