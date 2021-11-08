package cmd

import (
	"compose-generator/model"
	"io/fs"
	"os"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------- CreateNewPredefinedTempate ----------------------------------------------------------

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

func TestCreateNewPredefinedTempate2(t *testing.T) { // Template already exists
	// Test data
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
		return true
	}
	logError = func(message string, exit bool) {
		assert.Equal(t, "Template dir already exists. Aborting.", message)
		assert.True(t, exit)
	}
	// Execute test
	config, service, readme := createNewPredefinedTemplate()
	// Assert
	assert.Nil(t, config)
	assert.Empty(t, service)
	assert.Empty(t, readme)
}

// ------------------------------------------------------------- SavePredefinedTemplate ------------------------------------------------------------

func TestSavePredefinedTemplate1(t *testing.T) { // Happy path
	// Test data
	config := &model.PredefinedTemplateConfig{
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
	service := "image: hello-world:${{CHILLIDB_VERSION}}\ncontainer_name: ${{PROJECT_NAME_CONTAINER}}-database-chillidb\nrestart: always\nnetworks:\n# ToDo: Insert or remove section\nports:\n# ToDo: Insert or remove section\nvolumes:\n# ToDo: Insert or remove section\nenv_file:\n# ToDo: Insert or delete section"
	readme := "## ChilliDB\nToDo: Insert software description here.\n\n### Setup\nToDo: Insert setup instructions here."
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	startProcess = func(text string) (s *spinner.Spinner) {
		assert.Equal(t, "Saving predefined service template ...", text)
		return nil
	}
	mkDir = func(name string, perm os.FileMode) error {
		assert.Equal(t, "/usr/lib/compose-generator/predefined-services/database/chillidb", name)
		return nil
	}
	marshalIdent = func(v interface{}, prefix, indent string) ([]byte, error) {
		assert.Empty(t, prefix)
		assert.Equal(t, 4, len(indent))
		return []byte{}, nil
	}
	writeFileCallCount := 0
	writeFile = func(filename string, data []byte, perm fs.FileMode) error {
		writeFileCallCount++
		switch writeFileCallCount {
		case 1:
			assert.Equal(t, "/usr/lib/compose-generator/predefined-services/database/chillidb/config.json", filename)
		case 2:
			assert.Equal(t, "/usr/lib/compose-generator/predefined-services/database/chillidb/service.yml", filename)
		case 3:
			assert.Equal(t, "/usr/lib/compose-generator/predefined-services/database/chillidb/README.md", filename)
		}
		return nil
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	// Execute test
	savePredefinedTemplate(config, service, readme)
	// Assert
	assert.Equal(t, 1, pelCallCount)
}
