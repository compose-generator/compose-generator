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
	/*// Test data
	serviceName := ""
	project := &spec.Project{}
	visitedServices := &[]string{}
	// Mock functions
	sliceContainsString = func(slice []string, i string) bool {

	}
	// Execute test
	VisitServiceDependencies(project, serviceName, visitedServices)
	// Assert*/

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
