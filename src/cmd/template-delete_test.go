/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"errors"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/stretchr/testify/assert"
)

func TestDelete1(t *testing.T) { // Happy path - Template name set
	// Test data
	customTemplatesPath := "/usr/lib/compose-generator/templates"
	expectedTemplatePath := "/usr/lib/compose-generator/templates/Test template"
	// Mock functions
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	askForTemplateMockable = func(question string) string {
		assert.Fail(t, "Unexpected call of askForTemplate")
		return ""
	}
	fileExists = func(path string) bool {
		assert.Equal(t, expectedTemplatePath, path)
		return true
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you really want to delete this template?", question)
		assert.False(t, defaultValue)
		return true
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Deleting project ...", text)
		return nil
	}
	removeAll = func(path string) error {
		assert.Equal(t, expectedTemplatePath, path)
		return nil
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	// Execute test
	result := delete("Test template", false)
	// Assert
	assert.Nil(t, result)
}

func TestDelete2(t *testing.T) { // Happy path - No template name set
	// Test data
	customTemplatesPath := "/usr/lib/compose-generator/templates"
	expectedTemplatePath := "/usr/lib/compose-generator/templates/Test template 1"
	// Mock functions
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	askForTemplateMockable = func(question string) string {
		assert.Equal(t, "Which template do you want to delete?", question)
		return "Test template 1"
	}
	fileExists = func(path string) bool {
		assert.Fail(t, "Unexpected call of fileExists")
		return true
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Fail(t, "Unexpected call of yesNoQuestion")
		return false
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Deleting project ...", text)
		return nil
	}
	removeAll = func(path string) error {
		assert.Equal(t, expectedTemplatePath, path)
		return nil
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	// Execute test
	result := delete("", true)
	// Assert
	assert.Nil(t, result)
}

func TestDelete3(t *testing.T) { // Error - template does not exist
	// Test data
	customTemplatesPath := "/usr/lib/compose-generator/templates"
	expectedTemplatePath := "/usr/lib/compose-generator/templates/Test template 2"
	// Mock functions
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	askForTemplateMockable = func(question string) string {
		assert.Fail(t, "Unexpected call of askForTemplate")
		return ""
	}
	fileExists = func(path string) bool {
		assert.Equal(t, expectedTemplatePath, path)
		return false
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Fail(t, "Unexpected call of yesNoQuestion")
		return true
	}
	logError = func(message string, exit bool) {
		assert.Equal(t, "Could not find template 'Test template 2'", message)
		assert.True(t, exit)
	}
	// Execute test
	result := delete("Test template 2", false)
	// Assert
	assert.Nil(t, result)
}

func TestDelete4(t *testing.T) { // Error - safety checks not confirmed
	// Test data
	customTemplatesPath := "../templates"
	expectedTemplatePath := "../templates/Test template 3"
	// Mock functions
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	askForTemplateMockable = func(question string) string {
		assert.Fail(t, "Unexpected call of askForTemplate")
		return ""
	}
	fileExists = func(path string) bool {
		assert.Equal(t, expectedTemplatePath, path)
		return true
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you really want to delete this template?", question)
		assert.False(t, defaultValue)
		return false
	}
	startProcess = func(text string) (s *spinner.Spinner) {
		assert.Fail(t, "Unexpected call of startProcess")
		return nil
	}
	// Execute test
	result := delete("Test template 3", false)
	// Assert
	assert.Nil(t, result)
}

func TestDelete5(t *testing.T) { // Error - error of removeAll
	// Test data
	customTemplatesPath := "/usr/lib/compose-generator/templates"
	expectedTemplatePath := "/usr/lib/compose-generator/templates/Test template 4"
	// Mock functions
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	askForTemplateMockable = func(question string) string {
		assert.Equal(t, "Which template do you want to delete?", question)
		return "Test template 4"
	}
	fileExists = func(path string) bool {
		assert.Fail(t, "Unexpected call of fileExists")
		return true
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Fail(t, "Unexpected call of yesNoQuestion")
		return false
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Deleting project ...", text)
		return nil
	}
	removeAll = func(path string) error {
		assert.Equal(t, expectedTemplatePath, path)
		return errors.New("Error message")
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Fail(t, "Unexpected call of stopProcess")
	}
	logError = func(message string, exit bool) {
		assert.Equal(t, "Could not delete template", message)
		assert.True(t, exit)
	}
	// Execute test
	result := delete("", true)
	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "Error message", result.Error())
}
