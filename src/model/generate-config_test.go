/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------- GetServiceConfigurationsByType ------------------------------------------------------

func TestGetServiceConfigurationsByType(t *testing.T) {
	// Test data
	config := &GenerateConfig{
		ServiceConfig: []ServiceConfig{
			{
				Name:   "angular",
				Type:   TemplateTypeFrontend,
				Params: map[string]string{"ANGULAR_PORT": "80"},
			},
			{
				Name:   "redis",
				Type:   TemplateTypeDatabase,
				Params: make(map[string]string),
			},
			{
				Name:   "jira",
				Type:   TemplateTypeFrontend,
				Params: map[string]string{"JIRA_VERSION": "8.18"},
			},
		},
	}
	expectedResult := []ServiceConfig{config.ServiceConfig[0], config.ServiceConfig[2]}
	// Execute test
	result := config.GetServiceConfigurationsByType(TemplateTypeFrontend)
	// Assert
	assert.Equal(t, 2, len(result))
	assert.EqualValues(t, expectedResult, result)
}
