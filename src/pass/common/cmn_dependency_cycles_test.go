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

// ---------------------------------------------------------- CommonCheckForDependencyCycles -------------------------------------------------------

func TestCommonCheckForDependencyCycles1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &spec.Project{
			Name: "test project",
		},
	}
	// Mock functions
	hasDependencyCyclesMockable = func(proj *spec.Project) bool {
		assert.Equal(t, project.Composition, proj)
		return true
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Configuration contains dependency cycles", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	CommonCheckForDependencyCycles(project)
	// Assert
	assert.Equal(t, 1, printErrorCallCount)
}

func TestCommonCheckForDependencyCycles2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &spec.Project{
			Name: "test project",
		},
	}
	// Mock functions
	hasDependencyCyclesMockable = func(proj *spec.Project) bool {
		assert.Equal(t, project.Composition, proj)
		return false
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	CommonCheckForDependencyCycles(project)
}

// ------------------------------------------------------------- VisitServiceDependencies ----------------------------------------------------------

func TestVisitServiceDependencies1(t *testing.T) {
	// Test data
	serviceName := "service 1"
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "service 1",
				DependsOn: spec.DependsOnConfig{
					"service 2": {
						Condition: spec.ServiceConditionStarted,
					},
				},
			},
			{
				Name: "service 2",
				DependsOn: spec.DependsOnConfig{
					"service 3": {
						Condition: spec.ServiceConditionStarted,
					},
				},
			},
			{
				Name: "service 3",
				DependsOn: spec.DependsOnConfig{
					"service 1": {
						Condition: spec.ServiceConditionStarted,
					},
				},
			},
		},
	}
	visitedServices := []string{}
	// Mock functions
	sliceContainsStringCallCount := 0
	sliceContainsString = func(slice []string, i string) bool {
		sliceContainsStringCallCount++
		switch sliceContainsStringCallCount {
		case 1:
			assert.EqualValues(t, []string{"service 1"}, slice)
			assert.Equal(t, "service 2", i)
		case 2:
			assert.EqualValues(t, []string{"service 1", "service 2"}, slice)
			assert.Equal(t, "service 3", i)
		case 3:
			assert.EqualValues(t, []string{"service 1", "service 2", "service 3"}, slice)
			assert.Equal(t, "service 1", i)
			return true
		}
		return false
	}
	// Execute test
	result := VisitServiceDependencies(project, serviceName, &visitedServices)
	// Assert
	assert.True(t, result)
	assert.Equal(t, 3, sliceContainsStringCallCount)
}

func TestVisitServiceDependencies2(t *testing.T) {
	// Test data
	serviceName := "service 2"
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "service 1",
				DependsOn: spec.DependsOnConfig{
					"service 2": {
						Condition: spec.ServiceConditionStarted,
					},
				},
			},
		},
	}
	visitedServices := []string{}
	// Mock functions
	sliceContainsString = func(slice []string, i string) bool {
		assert.Fail(t, "Unexpected call of sliceContainsString")
		return false
	}
	// Execute test
	result := VisitServiceDependencies(project, serviceName, &visitedServices)
	// Assert
	assert.False(t, result)
}

func TestVisitServiceDependencies3(t *testing.T) {
	// Test data
	serviceName := "service 1"
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "service 1",
				DependsOn: spec.DependsOnConfig{
					"service 2": {
						Condition: spec.ServiceConditionStarted,
					},
				},
			},
			{
				Name:      "service 2",
				DependsOn: make(spec.DependsOnConfig),
			},
		},
	}
	visitedServices := []string{}
	// Mock functions
	sliceContainsStringCallCount := 0
	sliceContainsString = func(slice []string, i string) bool {
		sliceContainsStringCallCount++
		switch sliceContainsStringCallCount {
		case 1:
			assert.EqualValues(t, []string{"service 1"}, slice)
			assert.Equal(t, "service 2", i)
		case 2:
			assert.EqualValues(t, []string{"service 1", "service 2"}, slice)
			assert.Equal(t, "service 3", i)
		}
		return false
	}
	// Execute test
	result := VisitServiceDependencies(project, serviceName, &visitedServices)
	// Assert
	assert.False(t, result)
}

// --------------------------------------------------------------- hasDependencyCycles -------------------------------------------------------------

func TestHasDependencyCycles1(t *testing.T) {
	// Test data
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "service 1",
			},
			{
				Name: "service 2",
			},
		},
	}
	// Mock functions
	visitServiceDependenciesCallCount := 0
	visitServiceDependenciesMockable = func(p *spec.Project, currentServiceName string, visitedServices *[]string) bool {
		visitServiceDependenciesCallCount++
		assert.Equal(t, project, p)
		if visitServiceDependenciesCallCount == 1 {
			assert.Equal(t, "service 1", currentServiceName)
			return false
		}
		assert.Equal(t, "service 2", currentServiceName)
		return true
	}
	// Execute test
	result := hasDependencyCycles(project)
	// Assert
	assert.True(t, result)
	assert.Equal(t, 2, visitServiceDependenciesCallCount)
}

func TestHasDependencyCycles2(t *testing.T) {
	// Test data
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "service 1",
			},
		},
	}
	// Mock functions
	visitServiceDependenciesCallCount := 0
	visitServiceDependenciesMockable = func(p *spec.Project, currentServiceName string, visitedServices *[]string) bool {
		visitServiceDependenciesCallCount++
		assert.Equal(t, project, p)
		assert.Equal(t, "service 1", currentServiceName)
		return false
	}
	// Execute test
	result := hasDependencyCycles(project)
	// Assert
	assert.False(t, result)
	assert.Equal(t, 1, visitServiceDependenciesCallCount)
}
