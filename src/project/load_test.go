package project

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------ LoadProject ------------------------------------------------------------------

func TestLoadProject(t *testing.T) {
	// Test data
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: true,
			WithReadme:    false,
		},
		GitignorePatterns: []string{},
		ReadmeChildPaths:  []string{"README.md"},
		ForceConfig:       false,
	}
	expectedOptions := LoadOptions{
		ComposeFileName: "docker.yml",
		WorkingDir:      "./context/",
	}
	// Mock functions
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		if fileExistsCallCount == 1 {
			assert.Equal(t, "./context/.gitignore", path)
			return true
		}
		assert.Equal(t, "./context/README.md", path)
		return false
	}
	loadComposeFileCallCount := 0
	loadComposeFileMockable = func(project *model.CGProject, opt LoadOptions) {
		loadComposeFileCallCount++
		assert.Equal(t, expectedOptions, opt)
	}
	loadGitignoreFileCallCount := 0
	loadGitignoreFileMockable = func(project *model.CGProject, opt LoadOptions) {
		loadGitignoreFileCallCount++
		assert.Equal(t, expectedOptions, opt)
	}
	loadCGFileCallCount := 0
	loadCGFileMockable = func(metadata *model.CGProjectMetadata, opt LoadOptions) {
		loadCGFileCallCount++
		assert.Equal(t, expectedOptions, opt)
	}
	// Execute test
	result := LoadProject(LoadFromComposeFile("docker.yml"), LoadFromDir("./context"))
	// Assert
	assert.Equal(t, expectedProject, result)
	assert.Equal(t, 1, loadComposeFileCallCount)
	assert.Equal(t, 1, loadGitignoreFileCallCount)
	assert.Equal(t, 1, loadCGFileCallCount)
	assert.Equal(t, 2, fileExistsCallCount)
}

// -------------------------------------------------------------- LoadProjectMetadata --------------------------------------------------------------

func TestLoadProjectMetadata(t *testing.T) {
	// Test data
	expectedMetadata := &model.CGProjectMetadata{
		WithGitignore:   true,
		WithReadme:      false,
		Name:            "Project name",
		ProductionReady: true,
	}
	expectedOptions := LoadOptions{
		ComposeFileName: "docker.yml",
		WorkingDir:      "./context/",
	}
	// Mock functions
	loadCGFileMockable = func(metadata *model.CGProjectMetadata, opt LoadOptions) {
		assert.Equal(t, expectedOptions, opt)

		metadata.WithGitignore = true
		metadata.WithReadme = false
		metadata.ProductionReady = true
		metadata.Name = "Project name"
	}
	// Execute test
	result := LoadProjectMetadata(LoadFromComposeFile("docker.yml"), LoadFromDir("./context"))
	// Assert
	assert.Equal(t, expectedMetadata, result)
}

// -------------------------------------------------------------- LoadTemplateService --------------------------------------------------------------

func TestLoadTemplateService(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			Name: "Project name",
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label: "Angular",
			},
		},
	}
	templateType := model.TemplateTypeDatabase
	expectedServiceName := "angular"
	expectedServiceConfig := &spec.ServiceConfig{
		Name: "frontend-angular",
	}
	// Mock functions
	loadComposeFileSingleServiceMockable = func(proj *model.CGProject, templates *model.SelectedTemplates, templateTypeName, serviceName string, opt LoadOptions) *spec.ServiceConfig {
		assert.Equal(t, project, proj)
		assert.Equal(t, selectedTemplates, templates)
		assert.Equal(t, templateType, templateTypeName)
		assert.Equal(t, expectedServiceName, serviceName)
		return expectedServiceConfig
	}
	// Execute test
	result := LoadTemplateService(project, selectedTemplates, templateType, expectedServiceName)
	// Assert
	assert.Equal(t, expectedServiceConfig, result)
}

// ---------------------------------------------------------------- loadComposeFile ----------------------------------------------------------------
