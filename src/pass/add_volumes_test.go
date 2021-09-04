package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

func TestAddVolumes1(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	// Mock functions
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	yesNoCallCount := 0
	YesNoQuestion = func(question string, defaultValue bool) bool {
		yesNoCallCount++
		switch yesNoCallCount {
		case 1:
			assert.Equal(t, "Do you want to add volumes to your service?", question)
			assert.False(t, defaultValue)
		case 2:
			assert.Equal(t, "Do you want to add an existing external volume (y) or link a directory / file (N)?", question)
			assert.False(t, defaultValue)
		case 3:
			assert.Equal(t, "Add another volume?", question)
			assert.True(t, defaultValue)
		case 4:
			assert.Equal(t, "Do you want to add an existing external volume (y) or link a directory / file (N)?", question)
			assert.False(t, defaultValue)
			return false
		case 5:
			assert.Equal(t, "Add another volume?", question)
			assert.True(t, defaultValue)
			return false
		}
		return true
	}
	askForExternalVolumeMockable = func(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
		if yesNoCallCount != 2 {
			assert.Fail(t, "Unexpected call of askForExternalVolume")
		}
	}
	askForFileVolumeMockable = func(service *spec.ServiceConfig, project *model.CGProject) {
		if yesNoCallCount != 4 {
			assert.Fail(t, "Unexpected call of askForFileVolume")
		}
	}
	Pel = func() {}
	// Execute test
	AddVolumes(service, project, cli)
}

func TestAddVolumes2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{}
	// Mock functions
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to add volumes to your service?", question)
		assert.False(t, defaultValue)
		return false
	}
	Pel = func() {
		assert.Fail(t, "Unexpected call of Pel")
	}
	// Execute test
	AddVolumes(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
}
