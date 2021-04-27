package util

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------- EvaluateConditionalSection ---------------------------------------------------------------

func TestEvaluateConditionalSection_True1(t *testing.T) {
	content := "property1: true\n#! if has service wordpress {\n#! property2: false\n#! }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSection_True2(t *testing.T) {
	content := "property1: true\n#! if has service postgres|has backend {\n#! property2: false\n#! }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSection_True3(t *testing.T) {
	content := "property1: true\n#! if var.BAR == \"test1\" {\n#! property2: false\n#! }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.Equal(t, expectation, result)
}

func TestEvaluateConditionalSection_False(t *testing.T) {
	content := "property1: true\n#! if var.BAR == \"invalid\" {\n#! property2: false\n#! }\nproperty3: true"
	expectation := "property1: true\nproperty2: false\nproperty3: true"
	result := EvaluateConditionalSections(content, templateData, varMap)
	assert.NotEqual(t, expectation, result)
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
var varMap = map[string]string{
	"FOO": "test",
	"BAR": "test1",
}

func TestEvaluateCondition_True1(t *testing.T) {
	condition := "has frontend"
	result := EvaluateCondition(condition, templateData, varMap)
	assert.True(t, result)
}

func TestEvaluateCondition_True2(t *testing.T) {
	condition := "has service angular"
	result := EvaluateCondition(condition, templateData, varMap)
	assert.True(t, result)
}

func TestEvaluateCondition_True3(t *testing.T) {
	condition := "var.FOO == \"test\""
	result := EvaluateCondition(condition, templateData, varMap)
	assert.True(t, result)
}

func TestEvaluateCondition_False1(t *testing.T) {
	condition := "has service postgres"
	result := EvaluateCondition(condition, templateData, varMap)
	assert.False(t, result)
}

func TestEvaluateCondition_False2(t *testing.T) {
	condition := "invalid condition"
	result := EvaluateCondition(condition, templateData, varMap)
	assert.False(t, result)
}
