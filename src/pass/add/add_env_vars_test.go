/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/AlecAivazis/survey/v2"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAddEnvVars(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{}
	value1 := "value 1"
	value2 := "value 2"
	expectedService := &spec.ServiceConfig{
		Environment: spec.MappingWithEquals{
			"VARIABLE1_NAME": &value1,
			"VARIABLE2_NAME": &value2,
		},
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	yesNoCallCount := 0
	yesNoQuestion = func(question string, defaultValue bool) bool {
		yesNoCallCount++
		switch yesNoCallCount {
		case 1:
			assert.Equal(t, "Do you want to provide environment variables to your service?", question)
			assert.False(t, defaultValue)
		case 2:
			assert.Equal(t, "Expose another environment variable?", question)
			assert.True(t, defaultValue)
		case 3:
			assert.Equal(t, "Expose another environment variable?", question)
			assert.True(t, defaultValue)
			return false
		}
		return true
	}
	textQuestionWithValidator = func(question string, fn survey.Validator) string {
		assert.Equal(t, "Variable name (BEST_PRACTICE_IS_CAPS):", question)
		if yesNoCallCount == 1 {
			return "VARIABLE1_NAME"
		}
		return "VARIABLE2_NAME"
	}
	textQuestion = func(question string) string {
		assert.Equal(t, "Variable value:", question)
		if yesNoCallCount == 1 {
			return "value 1"
		}
		return "value 2"
	}
	// Execute test
	AddEnvVars(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, 2, pelCallCount)
}
