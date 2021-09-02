package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAddContainerName(t *testing.T) {
	// Create test data
	service := &types.ServiceConfig{}
	project := &model.CGProject{}
	testContainerName := "Test container name"
	// Mock user input
	TextQuestionWithDefault = func(question, defaultValue string) string {
		return testContainerName
	}
	// Execute test
	AddContainerName(service, project)
	// Assert result
	assert.Equal(t, testContainerName, service.ContainerName)
}
