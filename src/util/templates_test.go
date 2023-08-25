/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------
// ---------------------------------------------------------- CheckForServiceTemplateUpdate --------------------------------------------------------

// -------------------------------------------------------- TemplateListToPreselectedLabelList -----------------------------------------------------

// ---------------------------------------------------------------- EvaluateProxyLabels ------------------------------------------------------------

func TestEvaluateProxyLabels1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		ProxyLabels: make(map[string]model.Labels),
		Vars:        model.Vars{},
	}
	expectedProject := &model.CGProject{
		ProxyLabels: map[string]model.Labels{
			"angular": {
				"this.is-a.test_label":               "Angular",
				"this.is-another.test_label.angular": "false",
			},
		},
		Vars: model.Vars{},
	}
	template := &model.PredefinedTemplateConfig{
		Proxied: true,
		Name:    "angular",
		Label:   "Angular",
	}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Name:  "angular",
				Label: "Angular",
			},
		},
		DatabaseServices: []model.PredefinedTemplateConfig{
			{
				Name:  "mongodb",
				Label: "MongoDB",
			},
		},
		ProxyServices: []model.PredefinedTemplateConfig{
			{
				Name:  "traefik",
				Label: "Traefik",
				ProxyLabels: []model.Label{
					{
						Name:      "this.is-a.test_label",
						Value:     "${{CURRENT_SERVICE_LABEL}}",
						Condition: "true",
					},
					{
						Name:      "this.is-another.test_label.${{CURRENT_SERVICE_NAME}}",
						Value:     "false",
						Condition: "services.database contains name == \"mongodb\"",
					},
				},
			},
		},
	}
	// Mock functions
	evaluateConditionCallCount := 0
	evaluateCondition = func(condition string, selected *model.SelectedTemplates, varMap model.Vars) bool {
		evaluateConditionCallCount++
		if evaluateConditionCallCount == 1 {
			assert.Equal(t, "true", condition)
		} else {
			assert.Equal(t, "services.database contains name == \"mongodb\"", condition)
		}
		assert.Equal(t, selectedTemplates, selected)
		assert.Equal(t, project.Vars, varMap)
		return true
	}
	// Execute test
	EvaluateProxyLabels(project, template, selectedTemplates)
	// Assert
	assert.Equal(t, expectedProject, project)
	assert.Equal(t, 2, evaluateConditionCallCount)
}

func TestEvaluateProxyLabels2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	expectedProject := &model.CGProject{}
	template := &model.PredefinedTemplateConfig{
		Proxied: false,
	}
	selectedTemplates := &model.SelectedTemplates{}
	// Execute test
	EvaluateProxyLabels(project, template, selectedTemplates)
	// Assert
	assert.Equal(t, expectedProject, project)
}

// -------------------------------------------------------------- TemplateListToLabelList ----------------------------------------------------------

func TestTemplateListToLabelList(t *testing.T) {
	// Test data
	templates := []model.PredefinedTemplateConfig{
		{Label: "Angular"},
		{Label: "Wordpress"},
		{Label: "MySQL"},
	}
	// Execute test
	result := TemplateListToLabelList(templates)
	// Assert
	assert.EqualValues(t, result, []string{"Angular", "Wordpress", "MySQL"})
}

// --------------------------------------------------------- TemplateListToPreselectedLabelList ----------------------------------------------------

func TestTemplateListToPreselectedLabelList(t *testing.T) {
	// Test data
	templates := []model.PredefinedTemplateConfig{
		{
			Label:       "Angular",
			Preselected: "services.database contains name == \"mysql\" | services.database contains name == \"mariadb\"",
		},
		{
			Label:       "Wordpress",
			Preselected: "true",
		},
		{
			Label:       "MySQL",
			Preselected: "false",
		},
	}
	selected := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Name:  "mariadb",
				Label: "MariaDB",
			},
		},
	}
	// Mock functions
	evaluateConditionCallCount := 0
	evaluateCondition = func(condition string, s *model.SelectedTemplates, varMap model.Vars) bool {
		evaluateConditionCallCount++
		assert.Equal(t, selected, s)
		assert.Nil(t, varMap)
		if evaluateConditionCallCount == 1 {
			assert.Equal(t, "services.database contains name == \"mysql\" | services.database contains name == \"mariadb\"", condition)
		} else if evaluateConditionCallCount == 2 {
			assert.Equal(t, "true", condition)
		} else if evaluateConditionCallCount == 3 {
			assert.Equal(t, "false", condition)
			return false
		}
		return true
	}
	// Execute test
	result := TemplateListToPreselectedLabelList(templates, selected)
	// Assert
	assert.EqualValues(t, result, []string{"Angular", "Wordpress"})
}
