package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"testing"

	"github.com/stretchr/testify/assert"

	spec "github.com/compose-spec/compose-go/types"
)

func TestAddEnvFiles1(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{}
	expectedService := &spec.ServiceConfig{
		EnvFile: spec.StringList{
			"./test/env.env",
		},
	}
	// Mock functions
	pelCallCount := 0
	Pel = func() {
		pelCallCount++
	}
	yesNoCallCount := 0
	YesNoQuestion = func(question string, defaultValue bool) (result bool) {
		yesNoCallCount++
		switch yesNoCallCount {
		case 1:
			assert.Equal(t, "Do you want to provide environment files to your service?", question)
			assert.False(t, defaultValue)
		case 2:
			assert.Equal(t, "Add another environment file?", question)
			assert.True(t, defaultValue)
		case 3:
			assert.Equal(t, "Add another environment file?", question)
			assert.True(t, defaultValue)
			return false
		}
		return true
	}
	TextQuestionWithDefaultAndSuggestions = func(question, defaultValue string, fn util.Suggest) string {
		assert.Equal(t, "Where is your env file located?", question)
		assert.Equal(t, "environment.env", defaultValue)
		if yesNoCallCount == 1 {
			return "./test/env.env"
		}
		return "environment.env"
	}
	FileExists = func(path string) bool {
		if yesNoCallCount == 1 {
			assert.Equal(t, "./test/env.env", path)
		} else {
			assert.Equal(t, "environment.env", path)
			return false
		}
		return true
	}
	IsDir = func(path string) bool {
		if yesNoCallCount == 1 {
			assert.Equal(t, "./test/env.env", path)
		} else {
			assert.Equal(t, "environment.env", path)
		}
		return false
	}
	Error = func(description string, err error, exit bool) {
		assert.Equal(t, "File is not valid. Please select another file", description)
		assert.Nil(t, err)
		assert.False(t, exit)
	}
	// Execute test
	AddEnvFiles(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, 2, pelCallCount)
}

func TestAddEnvFiles2(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{}
	expectedService := &spec.ServiceConfig{}
	// Mock functions
	pelCallCount := 0
	Pel = func() {
		pelCallCount++
	}
	YesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Do you want to provide environment files to your service?", question)
		assert.False(t, defaultValue)
		return false
	}
	// Execute test
	AddEnvFiles(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Zero(t, pelCallCount)
}
