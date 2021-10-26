/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package project

import (
	"compose-generator/model"
	"testing"

	"github.com/compose-spec/compose-go/loader"
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
			Name:          "Example project",
			ContainerName: "example-project",
		},
		GitignorePatterns: []string{},
		ReadmeChildPaths:  []string{"README.md"},
		ForceConfig:       false,
		Vars: model.Vars{
			"PROJECT_NAME":           "Example project",
			"PROJECT_NAME_CONTAINER": "example-project",
		},
		ProxyVars: make(map[string]model.Vars),
		ProxyLabels: make(map[string]model.Labels),
		Secrets:   []model.ProjectSecret{},
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
		metadata.Name = "Example project"
		metadata.ContainerName = "example-project"
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

func TestLoadComposeFile1(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	options := LoadOptions{
		WorkingDir:      "./context/",
		ComposeFileName: "docker-compose.yml",
	}
	composeFileContent := "services:\nbackend-django:\nbuild:\ncontext: ./backend-django\ncontainer_name: example-project-backend-django\nports:\n- mode: ingress\ntarget: 8000\npublished: 8000\nprotocol: tcp\nrestart: always"
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Name:       "Project name",
			WorkingDir: "./context/",
		},
	}
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, "./context/docker-compose.yml", path)
		return true
	}
	readFile = func(filename string) ([]byte, error) {
		assert.Equal(t, "./context/docker-compose.yml", filename)
		return []byte(composeFileContent), nil
	}
	parseCompositionYAML = func(source []byte) (map[string]interface{}, error) {
		return make(map[string]interface{}), nil
	}
	loadComposition = func(configDetails spec.ConfigDetails, options ...func(*loader.Options)) (*spec.Project, error) {
		return &spec.Project{
			Name: "Project name",
		}, nil
	}
	logError = func(message string, exit bool) {
		assert.Fail(t, "Unexpected call of logError")
	}
	// Execute test
	loadComposeFile(project, options)
	// Assert
	assert.Equal(t, expectedProject, project)
}
