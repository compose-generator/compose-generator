/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package util

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test data
var templateData = &model.SelectedTemplates{
	FrontendServices: []model.PredefinedTemplateConfig{
		{Label: "Angular", Name: "angular"},
		{Label: "Vue", Name: "vue"},
	},
	BackendServices: []model.PredefinedTemplateConfig{
		{Label: "Wordpress", Name: "wordpress"},
	},
}

var templateData2 = &model.SelectedTemplates{
	DbAdminServices: []model.PredefinedTemplateConfig{
		{Label: "PhpMyAdmin", Name: "phpmyadmin"},
	},
	TlsHelperService: []model.PredefinedTemplateConfig{
		{Label: "Lets Encrypt Companion", Name: "letsencrypt"},
	},
}

var varMap = map[string]string{
	"FOO": "test",
	"BAR": "test1",
}

// ---------------------------------------------------------------- EvaluateConditionalSection ---------------------------------------------------------------

func TestEvaluateConditionalSectionToString_True1(t *testing.T) {
	content := "property1: true\n#? if services.backend contains label == \"Wordpress\" {\n#property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSectionsToString(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSectionToString_True2(t *testing.T) {
	content := "property1: true\n#? if services.frontend contains name == \"vue\" | has templates.backend {\n#property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSectionsToString(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSectionToString_True3(t *testing.T) {
	content := "property1: true\n#? if var.BAR == \"test1\" {\n#property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSectionsToString(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSectionToString_False1(t *testing.T) {
	content := "property1: true\n#? if var.BAR == \"invalid\" {\n# property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty3: true"
	result := EvaluateConditionalSectionsToString(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSectionToString_False2(t *testing.T) {
	content := "property1: true\n#? if has services.database {\n# property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty3: true"
	result := EvaluateConditionalSectionsToString(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

// ---------------------------------------------------------------- EvaluateCondition ---------------------------------------------------------------

func TestEvaluateCondition_True1(t *testing.T) {
	condition := "has services.frontend"
	result := EvaluateCondition(condition, templateData, varMap)
	assert.True(t, result)
}

func TestEvaluateCondition_True2(t *testing.T) {
	condition := "services.frontend[0].name == \"angular\""
	result := EvaluateCondition(condition, templateData, varMap)
	assert.True(t, result)
}

func TestEvaluateCondition_True3(t *testing.T) {
	condition := "var.FOO == \"test\""
	result := EvaluateCondition(condition, templateData, varMap)
	assert.True(t, result)
}

func TestEvaluateCondition_False1(t *testing.T) {
	condition := "services.database[1].name == \"postgres\""
	result := EvaluateCondition(condition, templateData, varMap)
	assert.False(t, result)
}

func TestEvaluateCondition_False2(t *testing.T) {
	condition := "invalid condition"
	result := EvaluateCondition(condition, templateData, varMap)
	assert.False(t, result)
}

// ----------------------------------------------------------------- PrepareDataInput ---------------------------------------------------------------

func TestPrepareInputData1(t *testing.T) {
	result := prepareInputData(templateData, varMap)
	expected := "{\"services\":{\"frontend\":[{\"label\":\"Angular\",\"name\":\"angular\"},{\"label\":\"Vue\",\"name\":\"vue\"}],\"backend\":[{\"label\":\"Wordpress\",\"name\":\"wordpress\"}]},\"var\":{\"BAR\":\"test1\",\"FOO\":\"test\"}}"
	assert.Equal(t, expected, result)
}

func TestPrepareInputData2(t *testing.T) {
	result := prepareInputData(templateData2, varMap)
	expected := "{\"services\":{\"dbadmin\":[{\"label\":\"PhpMyAdmin\",\"name\":\"phpmyadmin\"}],\"tlshelper\":[{\"label\":\"Lets Encrypt Companion\",\"name\":\"letsencrypt\"}]},\"var\":{\"BAR\":\"test1\",\"FOO\":\"test\"}}"
	assert.Equal(t, expected, result)
}

// -------------------------------------------------------------- CheckIfCComIsInstalled ------------------------------------------------------------

func TestCheckIfCComIsInstalled1(t *testing.T) {
	// Mock functions
	commandExists = func(cmd string) bool {
		assert.Equal(t, "ccom", cmd)
		return true
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	EnsureCComIsInstalled()
}

func TestCheckIfCComIsInstalled2(t *testing.T) {
	// Mock functions
	commandExists = func(cmd string) bool {
		assert.Equal(t, "ccom", cmd)
		return false
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "CCom could not be found on your system. Please go to https://github.com/compose-generator/compose-generator/releases/latest to download the latest version.", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	EnsureCComIsInstalled()
	// Assert
	assert.Equal(t, 1, printErrorCallCount)
}

// -------------------------------------------------------------- CheckIfDockerIsRunning ------------------------------------------------------------

func TestCheckIfDockerIsRunning1(t *testing.T) {
	// Mock functions
	isDockerRunning = func() bool {
		return true
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	EnsureDockerIsRunning()
}

func TestCheckIfDockerIsRunning2(t *testing.T) {
	// Mock functions
	isDockerRunning = func() bool {
		return false
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Docker engine is not running. Please start it and execute Compose Generator again.", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	EnsureDockerIsRunning()
	// Assert
	assert.Equal(t, 1, printErrorCallCount)
}
