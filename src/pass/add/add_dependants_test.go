package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

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
