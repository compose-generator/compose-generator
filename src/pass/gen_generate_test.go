package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/stretchr/testify/assert"
)

func TestGenerate1(t *testing.T) {
	// Test data
	predefinedTemplatesPath := "../predefined-templates"
	project := &model.CGProject{}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Name: "angular",
			},
		},
		ProxyService: []model.PredefinedTemplateConfig{
			{
				Name: "nginx",
			},
		},
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Generating configuration from 2 template(s) ...", text)
		return nil
	}
	getPredefinedServicesPath = func() string {
		return predefinedTemplatesPath
	}
	generateServiceCallCount := 0
	generateServiceMockable = func(
		proj *model.CGProject,
		selectedTemplates *model.SelectedTemplates,
		template model.PredefinedTemplateConfig,
		templateType, serviceName string,
	) {
		generateServiceCallCount++
		if generateServiceCallCount == 1 {
			assert.Equal(t, model.TemplateTypeFrontend, templateType)
			assert.Equal(t, selectedTemplates.FrontendServices[0], template)
		} else {
			assert.Equal(t, model.TemplateTypeProxy, templateType)
			assert.Equal(t, selectedTemplates.ProxyService[0], template)
		}
	}
	stopProcessCallCount := 0
	stopProcess = func(s *spinner.Spinner) {
		stopProcessCallCount++
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	Generate(project, selectedTemplates)
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, 1, stopProcessCallCount)
}

func TestGenerate2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	selectedTemplates := &model.SelectedTemplates{}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Fail(t, "Unexpected error of startProcess")
		return nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "No templates selected. Aborting ...", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	Generate(project, selectedTemplates)
	// Assert
	assert.Equal(t, 1, pelCallCount)
}
