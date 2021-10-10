/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package pass

import (
	"compose-generator/model"
	"errors"
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
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
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
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "You must specify a project name!", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	LoadGenerateConfig(project, config, "")
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, 1, printErrorCallCount)
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
	fileExists = func(path string) bool {
		assert.Equal(t, configPath, path)
		return true
	}
	openFile = func(name string) (*os.File, error) {
		assert.Equal(t, configPath, name)
		return nil, nil
	}
	readAllFromFile = func(r io.Reader) ([]byte, error) {
		return []byte{}, nil
	}
	unmarshalYaml = func(in []byte, out interface{}) error {
		return nil
	}
	// Execute test
	LoadGenerateConfig(project, config, configPath)
	// Assert
	assert.Equal(t, expectedConfig, config)
	assert.Equal(t, expectedProject, project)
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
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, configPath, path)
		return false
	}
	openFile = func(name string) (*os.File, error) {
		assert.Fail(t, "Unexpected call of openFile")
		return nil, nil
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Config file could not be found", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	LoadGenerateConfig(project, config, configPath)
	// Assert
	assert.Equal(t, 1, printErrorCallCount)
}

func TestLoadGenerateConfig5(t *testing.T) {
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
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, configPath, path)
		return true
	}
	openFile = func(name string) (*os.File, error) {
		assert.Equal(t, configPath, name)
		return nil, errors.New("Error message")
	}
	readAllFromFile = func(r io.Reader) ([]byte, error) {
		assert.Fail(t, "Unexpected call of readAllFromFile")
		return []byte{}, nil
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Could not load config file. Permissions granted?", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	LoadGenerateConfig(project, config, configPath)
	// Assert
	assert.Equal(t, 1, printErrorCallCount)
}

func TestLoadGenerateConfig6(t *testing.T) {
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
		return []byte{}, errors.New("Error message")
	}
	unmarshalYaml = func(in []byte, out interface{}) error {
		assert.Fail(t, "Unexpected call of unmarshalYaml")
		return nil
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Could not load config file. Permissions granted?", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	LoadGenerateConfig(project, config, configPath)
	// Assert
	assert.Equal(t, 1, printErrorCallCount)
}

func TestLoadGenerateConfig7(t *testing.T) {
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
		return []byte{}, nil
	}
	unmarshalYaml = func(in []byte, out interface{}) error {
		return errors.New("Error message")
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Could not unmarshal config file", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	LoadGenerateConfig(project, config, configPath)
	// Assert
	assert.Equal(t, 1, printErrorCallCount)
}
