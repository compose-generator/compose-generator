/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAddRestart1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: true,
		},
	}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{
		Restart: "unless-stopped",
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	menuQuestion = func(label string, items []string) string {
		assert.Equal(t, "When should the service get restarted?", label)
		assert.EqualValues(t, []string{"always", "on-failure", "unless-stopped", "no"}, items)
		return "unless-stopped"
	}
	// Execute test
	AddRestart(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, 2, pelCallCount)
}

func TestAddRestart2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: false,
		},
	}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{}
	// Mock functions
	pel = func() {
		assert.Fail(t, "Unexpected call of pel")
	}
	// Execute test
	AddRestart(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
}
