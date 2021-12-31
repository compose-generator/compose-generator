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

	"github.com/stretchr/testify/assert"

	spec "github.com/compose-spec/compose-go/types"
)

func TestAddDepends1(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				spec.ServiceConfig{
					Name:      "Service 1",
					DependsOn: make(spec.DependsOnConfig),
				},
				spec.ServiceConfig{
					Name:      "Service 2",
					DependsOn: make(spec.DependsOnConfig),
				},
				spec.ServiceConfig{
					Name:      "Service 3",
					DependsOn: make(spec.DependsOnConfig),
				},
			},
		},
	}
	expectedService := &spec.ServiceConfig{
		DependsOn: spec.DependsOnConfig{
			"Service 2": {
				Condition: spec.ServiceConditionStarted,
			},
			"Service 1": {
				Condition: spec.ServiceConditionStarted,
			},
		},
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Do you want your service to depend on other services?", question)
		assert.False(t, defaultValue)
		return true
	}
	multiSelectMenuQuestion = func(label string, items []string) (result []string) {
		assert.Equal(t, "Which ones?", label)
		assert.EqualValues(t, []string{"Service 1", "Service 2", "Service 3"}, items)
		return []string{"Service 2", "Service 1"}
	}
	// Execute test
	AddDepends(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, 2, pelCallCount)
}

func TestAddDepends2(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{},
			},
		},
	}
	expectedService := &spec.ServiceConfig{}
	// Mock functions
	pel = func() {
		assert.Fail(t, "Unexpected call of pel")
	}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Do you want your service to depend on other services?", question)
		assert.False(t, defaultValue)
		return false
	}
	// Execute test
	AddDepends(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
}

func TestAddDepends3(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{
		Composition: &spec.Project{},
	}
	expectedService := &spec.ServiceConfig{}
	// Mock functions
	pel = func() {
		assert.Fail(t, "Unexpected call of pel")
	}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Do you want your service to depend on other services?", question)
		assert.False(t, defaultValue)
		return false
	}
	// Execute test
	AddDepends(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
}
