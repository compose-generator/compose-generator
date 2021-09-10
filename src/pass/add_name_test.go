package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAddName1(t *testing.T) {
	// Create test data
	service := &types.ServiceConfig{}
	project := &model.CGProject{
		Composition: &types.Project{},
	}
	testName := "Test name"
	testContainerName := "test-name"
	// Mock user input
	textQuestionWithDefault = func(question, defaultValue string) string {
		return testName
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Error function was called")
	}
	// Execute test
	AddName(service, project)
	// Assert result
	assert.Equal(t, testName, service.Name)
	assert.Equal(t, testContainerName, service.ContainerName)
}

func TestAddName2(t *testing.T) {
	// Create test data
	testName1 := "Test name"
	testName2 := "Other test name"
	testContainerName := "other-test-name"
	service := &types.ServiceConfig{}
	project := &model.CGProject{
		Composition: &types.Project{
			Services: types.Services{
				types.ServiceConfig{
					Name: testName1,
				},
			},
		},
	}
	// Mock user input
	callCounter := 0
	textQuestionWithDefault = func(question, defaultValue string) string {
		callCounter++
		if callCounter == 2 {
			return testName2
		}
		return testName1
	}
	printError = func(description string, err error, exit bool) {
		if callCounter == 2 {
			assert.Fail(t, "Error function was called")
		} else {
			assert.Equal(t, "This service name already exists. Please choose a different one", description)
		}
	}
	// Execute test
	AddName(service, project)
	// Assert result
	assert.Equal(t, testName2, service.Name)
	assert.Equal(t, testContainerName, service.ContainerName)
}
