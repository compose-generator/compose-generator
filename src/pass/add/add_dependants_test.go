/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------ AddDependants ----------------------------------------------------------------

func TestAddDependants1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				spec.ServiceConfig{
					Name:      "Service 1",
					DependsOn: make(spec.DependsOnConfig),
				},
				spec.ServiceConfig{
					Name: "Service 2",
					DependsOn: spec.DependsOnConfig{
						"Service 1": {
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				spec.ServiceConfig{
					Name:      "Service 3",
					DependsOn: make(spec.DependsOnConfig),
				},
			},
		},
	}
	service := &spec.ServiceConfig{
		Name: "Service 0",
		DependsOn: spec.DependsOnConfig{
			"Service 2": {
				Condition: spec.ServiceConditionStarted,
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				spec.ServiceConfig{
					Name:      "Service 1",
					DependsOn: make(spec.DependsOnConfig),
				},
				spec.ServiceConfig{
					Name: "Service 2",
					DependsOn: spec.DependsOnConfig{
						"Service 1": {
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				spec.ServiceConfig{
					Name: "Service 3",
					DependsOn: spec.DependsOnConfig{
						"Service 0": {
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
			},
		},
	}
	// Mock functions
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want other services depend on the new one?", question)
		assert.False(t, defaultValue)
		return true
	}
	multiSelectMenuQuestion = func(label string, items []string) (result []string) {
		assert.Equal(t, "Which ones?", label)
		assert.EqualValues(t, []string{"Service 1", "Service 2", "Service 3"}, items)
		return []string{"Service 1", "Service 3"}
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	printWarningCallCount := 0
	printWarning = func(description string) {
		printWarningCallCount++
		assert.Equal(t, "Could not add dependency from 'Service 1' to 'Service 0' because it would cause a cycle", description)
	}
	checkForDependencyCycleCallCount := 0
	checkForDependencyCycleMockable = func(currentService *spec.ServiceConfig, otherServiceName string, project *spec.Project) bool {
		checkForDependencyCycleCallCount++
		if checkForDependencyCycleCallCount == 1 {
			assert.Equal(t, "Service 1", otherServiceName)
			return true
		}
		assert.Equal(t, "Service 3", otherServiceName)
		return false
	}
	// Execute test
	AddDependants(service, project)
	// Assert
	assert.Equal(t, expectedProject, project)
	assert.Equal(t, 2, pelCallCount)
	assert.Equal(t, 1, printWarningCallCount)
	assert.Equal(t, 2, checkForDependencyCycleCallCount)
}

func TestAddDependants2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	expectedProject := &model.CGProject{}
	// Mock functions
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Do you want other services depend on the new one?", question)
		assert.False(t, defaultValue)
		return false
	}
	multiSelectMenuQuestion = func(label string, items []string) (result []string) {
		return []string{}
	}
	pel = func() {
		assert.Fail(t, "Unexpected call of pel")
	}
	// Execute test
	AddDependants(service, project)
	// Assert
	assert.Equal(t, expectedProject, project)
}

// ------------------------------------------------------------- checkForDependencyCycle -----------------------------------------------------------

func TestCheckForDependencyCycle1(t *testing.T) {
	// Test data
	otherServiceName := "second service"
	currentService := &spec.ServiceConfig{
		Name: "current-service",
		DependsOn: spec.DependsOnConfig{
			"third service": {
				Condition: spec.ServiceConditionStarted,
			},
			"second service": {
				Condition: spec.ServiceConditionStarted,
			},
		},
	}
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "first service",
			},
			{
				Name: "second service",
			},
			{
				Name: "third service",
				DependsOn: spec.DependsOnConfig{
					"first service": {
						Condition: spec.ServiceConditionStarted,
					},
				},
			},
		},
	}
	// Mock functions
	visitServiceDependenciesCallCount := 0
	visitServiceDependencies = func(p *spec.Project, currentServiceName string, visitedServices *[]string) bool {
		visitServiceDependenciesCallCount++
		assert.EqualValues(t, []string{currentService.Name, otherServiceName}, *visitedServices)
		assert.Contains(t, []string{"second service", "third service"}, currentServiceName)
		return visitServiceDependenciesCallCount != 1
	}
	// Execute test
	result := checkForDependencyCycle(currentService, otherServiceName, project)
	// Assert
	assert.True(t, result)
	assert.Equal(t, 2, visitServiceDependenciesCallCount)
}

func TestCheckForDependencyCycle2(t *testing.T) {
	otherServiceName := "second service"
	currentService := &spec.ServiceConfig{
		Name: "current-service",
		DependsOn: spec.DependsOnConfig{
			"third service": {
				Condition: spec.ServiceConditionStarted,
			},
			"first service": {
				Condition: spec.ServiceConditionStarted,
			},
		},
	}
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "first service",
			},
			{
				Name: "second service",
			},
			{
				Name: "third service",
			},
		},
	}
	// Mock functions
	visitServiceDependenciesCallCount := 0
	visitServiceDependencies = func(p *spec.Project, currentServiceName string, visitedServices *[]string) bool {
		visitServiceDependenciesCallCount++
		assert.EqualValues(t, []string{currentService.Name, otherServiceName}, *visitedServices)
		assert.Contains(t, []string{"first service", "third service"}, currentServiceName)
		return false
	}
	// Execute test
	result := checkForDependencyCycle(currentService, otherServiceName, project)
	// Assert
	assert.False(t, result)
	assert.Equal(t, 2, visitServiceDependenciesCallCount)
}
