/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateChooseBackends1(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	config := &model.GenerateConfig{
		FromFile: true,
	}
	available := &model.AvailableTemplates{
		BackendServices: []model.PredefinedTemplateConfig{
			{
				Name:  "test-backend",
				Label: "Test Backend",
				Questions: []model.Question{
					{
						Text:         "Question 1",
						Variable:     "QUESTION_1",
						DefaultValue: "10",
					},
					{
						Text:         "Question 2",
						Variable:     "QUESTION_2",
						DefaultValue: "Test",
					},
				},
			},
		},
	}
	selected := &model.SelectedTemplates{}
	expectedProject := &model.CGProject{
		Vars: map[string]string{
			"QUESTION_1": "10",
			"QUESTION_2": "Extended test",
		},
	}
	expectedSelected := &model.SelectedTemplates{
		BackendServices: []model.PredefinedTemplateConfig{
			{
				Name:  "test-backend",
				Label: "Test Backend",
				Questions: []model.Question{
					{
						Text:         "Question 1",
						Variable:     "QUESTION_1",
						DefaultValue: "10",
					},
					{
						Text:         "Question 2",
						Variable:     "QUESTION_2",
						DefaultValue: "Test",
					},
				},
			},
		},
	}
	// Mock functions
	getServiceConfigurationsByType = func(config *model.GenerateConfig, templateType string) []model.ServiceConfig {
		assert.Equal(t, model.TemplateTypeBackend, templateType)
		return []model.ServiceConfig{
			{
				Name: "test-backend",
				Type: model.TemplateTypeBackend,
				Params: map[string]string{
					"QUESTION_2": "Extended test",
				},
			},
		}
	}
	// Execute test
	GenerateChooseBackends(project, available, selected, config)
	// Assert
	assert.Equal(t, expectedSelected, selected)
	assert.Equal(t, expectedProject, project)
}

func TestGenerateChooseBackends2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	config := &model.GenerateConfig{
		FromFile: false,
	}
	available := &model.AvailableTemplates{
		BackendServices: []model.PredefinedTemplateConfig{
			{
				Name:  "test-backend",
				Label: "Test Backend",
				Questions: []model.Question{
					{
						Text:         "Question 1",
						Variable:     "QUESTION_1",
						DefaultValue: "10",
					},
					{
						Text:         "Question 2",
						Variable:     "QUESTION_2",
						DefaultValue: "Test",
					},
				},
			},
		},
	}
	selected := &model.SelectedTemplates{}
	expectedProject := &model.CGProject{}
	expectedSelected := &model.SelectedTemplates{
		BackendServices: []model.PredefinedTemplateConfig{
			{
				Name:  "test-backend",
				Label: "Test Backend",
				Questions: []model.Question{
					{
						Text:         "Question 1",
						Variable:     "QUESTION_1",
						DefaultValue: "10",
					},
					{
						Text:         "Question 2",
						Variable:     "QUESTION_2",
						DefaultValue: "Test",
					},
				},
			},
		},
	}
	// Mock functions
	templateListToLabelList = func(templates []model.PredefinedTemplateConfig) []string {
		assert.Equal(t, expectedSelected.BackendServices, templates)
		return []string{"Test Backend"}
	}
	templateListToPreselectedLabelList = func(templates []model.PredefinedTemplateConfig, selected *model.SelectedTemplates) []string {
		assert.Equal(t, expectedSelected.BackendServices, templates)
		return []string{}
	}
	multiSelectMenuQuestionIndex = func(label string, items, defaultItems []string) []int {
		assert.Equal(t, "Which backend services do you need?", label)
		assert.Equal(t, []string{"Test Backend", "Custom backend service"}, items)
		return []int{0}
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	askTemplateQuestions = func(project *model.CGProject, template *model.PredefinedTemplateConfig) {
		assert.Equal(t, available.BackendServices[0], *template)
	}
	askTemplateProxyQuestions = func(project *model.CGProject, template *model.PredefinedTemplateConfig, selectedTemplates *model.SelectedTemplates) {
		assert.Equal(t, available.BackendServices[0], *template)
	}
	askSecretQuestions = func(project *model.CGProject, template *model.PredefinedTemplateConfig) {
		assert.Equal(t, available.BackendServices[0], *template)
	}
	askForCustomVolumePaths = func(project *model.CGProject, template *model.PredefinedTemplateConfig) {
		assert.Equal(t, available.BackendServices[0], *template)
	}
	// Execute test
	GenerateChooseBackends(project, available, selected, config)
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, expectedSelected, selected)
	assert.Equal(t, expectedProject, project)
}
