package cmd

import (
	"compose-generator/model"
	"errors"
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------- askForTemplate ----------------------------------------------------------------

func TestAskForTemplate1(t *testing.T) {
	// Test data
	templateList := map[string]*model.CGProjectMetadata{
		"Template 1": {
			Name:           "Template 1",
			LastModifiedAt: 1632579623478000000,
		},
		"Template 2": {
			Name:           "Template 2",
			LastModifiedAt: 1632579661510000000,
		},
	}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading template list ...", text)
		return nil
	}
	getTemplateMetadataListMockable = func() map[string]*model.CGProjectMetadata {
		return templateList
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	menuQuestionIndex = func(label string, items []string) int {
		assert.Equal(t, "Which template do you want to load?", label)
		assert.ElementsMatch(t, []string{"Template 1 (Saved at: Sep-25-21 2:20:23 PM)", "Template 2 (Saved at: Sep-25-21 2:21:01 PM)"}, items)
		if items[0] == "Template 1 (Saved at: Sep-25-21 2:20:23 PM)" {
			return 0
		}
		return 1
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := askForTemplate()
	// Assert
	assert.Equal(t, "Template 1", result)
	assert.Equal(t, 1, pelCallCount)
}

func TestAskForTemplate2(t *testing.T) {
	// Test data
	templateList := map[string]*model.CGProjectMetadata{}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading template list ...", text)
		return nil
	}
	getTemplateMetadataListMockable = func() map[string]*model.CGProjectMetadata {
		return templateList
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	menuQuestionIndex = func(label string, items []string) int {
		assert.Fail(t, "Unexpected call of menuQuestionIndex")
		return 1
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "No templates found. Use \"$ compose-generator save <template-name>\" to save one.", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	result := askForTemplate()
	// Assert
	assert.Equal(t, "", result)
	assert.Equal(t, 1, pelCallCount)
}

// ---------------------------------------------------------------- showTemplateList ---------------------------------------------------------------

func TestShowTemplateList1(t *testing.T) {
	// Test data
	templateList := map[string]*model.CGProjectMetadata{
		"Template 1": {
			Name:           "Template 1",
			LastModifiedAt: 1632579623478000000,
		},
		"Template 2": {
			Name:           "Template 2",
			LastModifiedAt: 1632579661510000000,
		},
	}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading template list ...", text)
		return nil
	}
	getTemplateMetadataListMockable = func() map[string]*model.CGProjectMetadata {
		return templateList
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	printHeadingCallCount := 0
	printHeading = func(text string) {
		printHeadingCallCount++
		assert.Equal(t, "List of all templates:", text)
	}
	// Execute test
	showTemplateList()
	// Assert
	assert.Equal(t, 2, pelCallCount)
	assert.Equal(t, 1, printHeadingCallCount)
}

func TestShowTemplateList2(t *testing.T) {
	// Test data
	templateList := map[string]*model.CGProjectMetadata{}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading template list ...", text)
		return nil
	}
	getTemplateMetadataListMockable = func() map[string]*model.CGProjectMetadata {
		return templateList
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	printHeading = func(text string) {
		assert.Fail(t, "Unexpected call of printHeading")
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "No templates found. Use \"$ compose-generator save <template-name>\" to save one.", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	showTemplateList()
	// Assert
	assert.Equal(t, 1, pelCallCount)
}

// ------------------------------------------------------------- getTemplateMetadataList -----------------------------------------------------------

func TestGetTemplateMetatdataList1(t *testing.T) {
	// Test data
	customTemplatesPath := "../../.github/test-files/templates"
	// Mock functions
	readDir = func(dirname string) ([]fs.FileInfo, error) {
		assert.Equal(t, customTemplatesPath, dirname)
		return ioutil.ReadDir(dirname)
	}
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := getTemplateMetadataList()
	// Assert
	assert.Equal(t, "template-1", result["../../.github/test-files/templates/Template 1"].ContainerName)
	assert.Equal(t, "template-2", result["../../.github/test-files/templates/Template 2"].ContainerName)
}

func TestGetTemplateMetatdataList2(t *testing.T) {
	// Test data
	customTemplatesPath := "../templates"
	// Mock functions
	readDir = func(dirname string) ([]fs.FileInfo, error) {
		assert.Equal(t, customTemplatesPath, dirname)
		return nil, errors.New("Error message")
	}
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Cannot access directory for custom templates", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	result := getTemplateMetadataList()
	// Assert
	assert.Nil(t, result)
}
