package util

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------- EvaluateConditionalSection ---------------------------------------------------------------

func TestEvaluateConditionalSection_True1(t *testing.T) {
	content := "property1: true\n#? if services.backend contains label == \"Wordpress\" {\n#property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSection_True2(t *testing.T) {
	content := "property1: true\n#? if services.frontend contains name == \"vue\" | has templates.backend {\n#property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSection_True3(t *testing.T) {
	content := "property1: true\n#? if var.BAR == \"test1\" {\n#property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSection_False1(t *testing.T) {
	content := "property1: true\n#? if var.BAR == \"invalid\" {\n# property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSection_False2(t *testing.T) {
	content := "property1: true\n#? if has services.database {\n# property2: false\n#? }\nproperty3: true"
	expectation := "property1: true\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

// ---------------------------------------------------------------- EvaluateCondition ---------------------------------------------------------------

var templateData = map[string][]model.ServiceTemplateConfig{
	"frontend": {
		{Label: "Angular", Name: "angular"},
		{Label: "Vue", Name: "vue"},
	},
	"backend": {
		{Label: "Wordpress", Name: "wordpress"},
	},
	"database": {},
}
var templateData2 = map[string][]model.ServiceTemplateConfig{
	"db-admin": {
		{Label: "PhpMyAdmin", Name: "phpmyadmin"},
	},
	"tls-helper": {
		{Label: "Lets Encrypt Companion", Name: "letsencrypt"},
	},
}
var varMap = map[string]string{
	"FOO": "test",
	"BAR": "test1",
}

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

func TestPrepareInputData1(t *testing.T) {
	result := PrepareInputData(templateData, varMap)
	expected := "{\"services\":{\"backend\":[{\"label\":\"Wordpress\",\"name\":\"wordpress\"}],\"frontend\":[{\"label\":\"Angular\",\"name\":\"angular\"},{\"label\":\"Vue\",\"name\":\"vue\"}]},\"var\":{\"BAR\":\"test1\",\"FOO\":\"test\"}}"
	assert.Equal(t, expected, result)
}

func TestPrepareInputData2(t *testing.T) {
	result := PrepareInputData(templateData2, varMap)
	expected := "{\"services\":{\"dbadmin\":[{\"label\":\"PhpMyAdmin\",\"name\":\"phpmyadmin\"}],\"tlshelper\":[{\"label\":\"Lets Encrypt Companion\",\"name\":\"letsencrypt\"}]},\"var\":{\"BAR\":\"test1\",\"FOO\":\"test\"}}"
	assert.Equal(t, expected, result)
}
