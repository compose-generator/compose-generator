/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDependencies(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{
		Name: "Current service",
	}
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				spec.ServiceConfig{
					Name: "Depending service 1",
					DependsOn: spec.DependsOnConfig{
						"Current service": {Condition: spec.ServiceConditionStarted},
					},
				},
				spec.ServiceConfig{
					Name: "Current service",
				},
				spec.ServiceConfig{
					Name: "Depending service 2",
					DependsOn: spec.DependsOnConfig{
						"Current service": {Condition: spec.ServiceConditionStarted},
					},
				},
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				spec.ServiceConfig{
					Name:      "Depending service 1",
					DependsOn: spec.DependsOnConfig{},
				},
				spec.ServiceConfig{
					Name: "Current service",
				},
				spec.ServiceConfig{
					Name:      "Depending service 2",
					DependsOn: spec.DependsOnConfig{},
				},
			},
		},
	}
	// Execute test
	RemoveDependencies(service, project)
	// Assert
	assert.Equal(t, expectedProject, project)
}
