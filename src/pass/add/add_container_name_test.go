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

	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAddContainerName(t *testing.T) {
	// Create test data
	testContainerName := "Test container name"
	oldContainerName := "Old container name"
	service := &types.ServiceConfig{
		ContainerName: oldContainerName,
	}
	project := &model.CGProject{}
	// Mock user input
	textQuestionWithDefault = func(question, defaultValue string) string {
		assert.Equal(t, "How do you want to call your container (best practice: lower, kebab cased):", question)
		assert.Equal(t, oldContainerName, defaultValue)
		return testContainerName
	}
	// Execute test
	AddContainerName(service, project)
	// Assert result
	assert.Equal(t, testContainerName, service.ContainerName)
}
