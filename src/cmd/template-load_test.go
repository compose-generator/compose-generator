package cmd

import (
	"compose-generator/model"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/stretchr/testify/assert"
)

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
		assert.EqualValues(t, []string{"Template 1 (Saved at: Sep-25-21 2:20:23 PM)", "Template 2 (Saved at: Sep-25-21 2:21:01 PM)"}, items)
		return 0
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
		assert.EqualValues(t, []string{"Template 1 (Saved at: Sep-25-21 2:20:23 PM)", "Template 2 (Saved at: Sep-25-21 2:21:01 PM)"}, items)
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
	assert.Equal(t, "Template 2", result)
	assert.Equal(t, 1, pelCallCount)
}
