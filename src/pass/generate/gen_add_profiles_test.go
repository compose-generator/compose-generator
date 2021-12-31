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

	"github.com/briandowns/spinner"
	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAddProfiles1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: true,
		},
		Composition: &types.Project{
			Services: types.Services{
				{
					Name:     "Service 1",
					Profiles: []string{ProfileDev},
				},
				{
					Name: "Service 2",
				},
				{
					Name: "Service 3",
				},
			},
		},
	}
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: true,
		},
		Composition: &types.Project{
			Services: types.Services{
				{
					Name:     "Service 1",
					Profiles: []string{ProfileDev},
				},
				{
					Name:     "Service 2",
					Profiles: []string{ProfileDev, ProfileProduction},
				},
				{
					Name:     "Service 3",
					Profiles: []string{ProfileDev, ProfileProduction},
				},
			},
		},
	}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Adding dev and prod profiles", text)
		return nil
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	// Execute functions
	GenerateAddProfiles(project)
	// Assert
	assert.Equal(t, expectedProject, project)
}

func TestGenerateAddProfiles2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: false,
		},
	}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Fail(t, "Unexpected call of startProcess")
		return nil
	}
	// Execute functions
	GenerateAddProfiles(project)
}
