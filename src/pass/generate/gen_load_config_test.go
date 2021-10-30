/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadGenerateConfig1(t *testing.T) {
	// Test data
	projectName := "Compose Generator 1.0.0"
	project := &model.CGProject{}
	config := &model.GenerateConfig{}
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			Name:            projectName,
			ProductionReady: false,
			ContainerName:   "compose-generator-1.0.0",
		},
		Vars: map[string]string{
			"PROJECT_NAME":           projectName,
			"PROJECT_NAME_CONTAINER": "compose-generator-1.0.0",
		},
	}
	expectedConfig := &model.GenerateConfig{
		FromFile:        false,
		ProjectName:     projectName,
		ProductionReady: false,
	}
	// Mock functions
	heading = func(text string) {
		assert.Equal(t, "Welcome to Compose Generator! 👋", text)
	}
	pl = func(text string) {
		assert.Equal(t, "Please continue by answering a few questions:", text)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	textQuestion = func(question string) string {
		assert.Equal(t, "What is the name of your project:", question)
		return projectName
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want the output to be production-ready?", question)
		assert.False(t, defaultValue)
		return false
	}
	logError = func(message string, exit bool) {
		assert.Fail(t, "Unexpected call of logError")
	}
	// Execute test
	LoadGenerateConfig(project, config, "")
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, expectedConfig, config)
	assert.Equal(t, expectedProject, project)
}

func TestLoadGenerateConfig2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	config := &model.GenerateConfig{}
	// Mock functions
	heading = func(text string) {
		assert.Equal(t, "Welcome to Compose Generator! 👋", text)
	}
	pl = func(text string) {
		assert.Equal(t, "Please continue by answering a few questions:", text)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	textQuestion = func(question string) string {
		assert.Equal(t, "What is the name of your project:", question)
		return ""
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Fail(t, "Unexpected call of yesNoQuestion")
		return false
	}
	logErrorCallCount := 0
	logError = func(message string, exit bool) {
		logErrorCallCount++
		assert.Equal(t, "You must specify a project name!", message)
		assert.True(t, exit)
	}
	// Execute test
	LoadGenerateConfig(project, config, "")
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, 1, logErrorCallCount)
}

func TestLoadGenerateConfig3(t *testing.T) {
	// Test data
	configPath := "./config-path.yml"
	projectName := "Test Project"
	project := &model.CGProject{}
	config := &model.GenerateConfig{
		ProjectName:     projectName,
		ProductionReady: true,
		ServiceConfig: []model.ServiceConfig{
			{
				Name: "angular",
				Type: "frontend",
				Params: map[string]string{
					"PARAM1": "value 1",
					"PARAM2": "value 2",
				},
			},
		},
	}
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			Name:            projectName,
			ProductionReady: true,
			ContainerName:   "test-project",
		},
		Vars: map[string]string{
			"PROJECT_NAME":           projectName,
			"PROJECT_NAME_CONTAINER": "test-project",
		},
	}
	expectedConfig := &model.GenerateConfig{
		FromFile:        true,
		ProjectName:     projectName,
		ProductionReady: true,
		ServiceConfig: []model.ServiceConfig{
			{
				Name: "angular",
				Type: "frontend",
				Params: map[string]string{
					"PARAM1": "value 1",
					"PARAM2": "value 2",
				},
			},
		},
	}
	// Mock functions
	isUrl = func(str string) bool {
		assert.Equal(t, configPath, str)
		return true
	}
	loadConfigFromUrlCallCount := 0
	loadConfigFromUrlMockable = func(config *model.GenerateConfig, configUrl string) {
		loadConfigFromUrlCallCount++
		assert.Equal(t, configPath, configUrl)
	}
	loadConfigFromFileMockable = func(config *model.GenerateConfig, configPath string) {
		assert.Fail(t, "Unexpected call of loadConfigFromFile")
	}
	// Execute test
	LoadGenerateConfig(project, config, configPath)
	// Assert
	assert.Equal(t, expectedConfig, config)
	assert.Equal(t, expectedProject, project)
	assert.Equal(t, 1, loadConfigFromUrlCallCount)
}

func TestLoadGenerateConfig4(t *testing.T) {
	// Test data
	configPath := "./config-path.yml"
	projectName := "Test Project"
	project := &model.CGProject{}
	config := &model.GenerateConfig{
		ProjectName:     projectName,
		ProductionReady: true,
		ServiceConfig: []model.ServiceConfig{
			{
				Name: "angular",
				Type: "frontend",
				Params: map[string]string{
					"PARAM1": "value 1",
					"PARAM2": "value 2",
				},
			},
		},
	}
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			Name:            projectName,
			ProductionReady: true,
			ContainerName:   "test-project",
		},
		Vars: map[string]string{
			"PROJECT_NAME":           projectName,
			"PROJECT_NAME_CONTAINER": "test-project",
		},
	}
	expectedConfig := &model.GenerateConfig{
		FromFile:        true,
		ProjectName:     projectName,
		ProductionReady: true,
		ServiceConfig: []model.ServiceConfig{
			{
				Name: "angular",
				Type: "frontend",
				Params: map[string]string{
					"PARAM1": "value 1",
					"PARAM2": "value 2",
				},
			},
		},
	}
	// Mock functions
	isUrl = func(str string) bool {
		assert.Equal(t, configPath, str)
		return false
	}
	loadConfigFromUrlMockable = func(config *model.GenerateConfig, configUrl string) {
		assert.Fail(t, "Unexpected call of loadConfigFromUrl")
	}
	loadConfigFromFileCallCount := 0
	loadConfigFromFileMockable = func(config *model.GenerateConfig, configPath string) {
		loadConfigFromFileCallCount++
		assert.Equal(t, configPath, configPath)
	}
	// Execute test
	LoadGenerateConfig(project, config, configPath)
	// Assert
	assert.Equal(t, expectedConfig, config)
	assert.Equal(t, expectedProject, project)
	assert.Equal(t, 1, loadConfigFromFileCallCount)
}

// --------------------------------------------------------------- loadConfigFromFile --------------------------------------------------------------

func TestLoadConfigFromFile1(t *testing.T) {
	// Test data
	config := &model.GenerateConfig{}
	configPath := "./test/path/config.yml"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, configPath, path)
		return true
	}
	openFile = func(name string) (*os.File, error) {
		assert.Equal(t, configPath, name)
		return nil, nil
	}
	readAllFromFile = func(r io.Reader) ([]byte, error) {
		return []byte("This is a test"), nil
	}
	unmarshalYaml = func(in []byte, out interface{}) error {
		assert.Equal(t, []byte("This is a test"), in)
		return nil
	}
	logError = func(message string, exit bool) {
		assert.Fail(t, "Unexpected logError")
	}
	// Execute test
	loadConfigFromFile(config, configPath)
}

func TestLoadConfigFromFile2(t *testing.T) {
	// Test data
	config := &model.GenerateConfig{}
	configPath := "./test/path/config.yml"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, configPath, path)
		return false
	}
	logError = func(message string, exit bool) {
		assert.Equal(t, "Config file could not be found", message)
		assert.True(t, exit)
	}
	// Execute test
	loadConfigFromFile(config, configPath)
}

// ---------------------------------------------------------------- loadConfigFromUrl --------------------------------------------------------------
