package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------- RemoveVolumes -------------------------------------------

func TestRemoveVolumes1(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{
		Name: "current-service",
		Volumes: []spec.ServiceVolumeConfig{
			{
				Source: "./volumes/test-volume",
			},
			{
				Source: "test",
			},
		},
	}
	project := &model.CGProject{
		Composition: &spec.Project{
			Volumes: spec.Volumes{
				"test": {},
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Volumes: spec.Volumes{},
		},
	}
	// Mock functions
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you really want to delete all attached volumes of 'current-service' on disk?", question)
		assert.False(t, defaultValue)
		return true
	}
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		return true
	}
	isVolumeUsedByOtherServicesMockable = func(volume *spec.ServiceVolumeConfig, serv *spec.ServiceConfig, proj *model.CGProject) bool {
		if fileExistsCallCount == 1 {
			assert.Equal(t, service.Volumes[0], *volume)
		} else {
			assert.Equal(t, service.Volumes[1], *volume)
		}
		assert.Equal(t, service, serv)
		assert.Equal(t, project, proj)
		return false
	}
	removeAllCallCount := 0
	removeAll = func(path string) error {
		removeAllCallCount++
		if removeAllCallCount == 1 {
			assert.Equal(t, "./volumes/test-volume", path)
		} else {
			assert.Equal(t, "test", path)
		}
		return nil
	}
	// Execute test
	RemoveVolumes(service, project)
	// Assert
	assert.Equal(t, 2, fileExistsCallCount)
	assert.Equal(t, 2, removeAllCallCount)
	assert.Equal(t, expectedProject, project)
}

func TestRemoveVolumes2(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{
		Name: "current-service",
		Volumes: []spec.ServiceVolumeConfig{
			{
				Source: "./volumes/test-volume",
			},
			{
				Source: "test",
			},
		},
	}
	project := &model.CGProject{
		Composition: &spec.Project{
			Volumes: spec.Volumes{
				"test": {},
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Volumes: spec.Volumes{
				"test": {},
			},
		},
	}
	// Mock functions
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you really want to delete all attached volumes of 'current-service' on disk?", question)
		assert.False(t, defaultValue)
		return false
	}
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		return false
	}
	isVolumeUsedByOtherServicesMockable = func(volume *spec.ServiceVolumeConfig, serv *spec.ServiceConfig, proj *model.CGProject) bool {
		if fileExistsCallCount == 1 {
			assert.Equal(t, service.Volumes[0], *volume)
		} else {
			assert.Equal(t, service.Volumes[1], *volume)
		}
		assert.Equal(t, service, serv)
		assert.Equal(t, project, proj)
		return false
	}
	removeAll = func(path string) error {
		assert.Fail(t, "Unexpected call of removeAll")
		return nil
	}
	// Execute test
	RemoveVolumes(service, project)
	// Assert
	assert.Equal(t, 2, fileExistsCallCount)
	assert.Equal(t, expectedProject, project)
}

// ------------------------------------ isVolumeUsedByOtherServices ------------------------------------

func TestIsVolumeUsedByOtherServices1(t *testing.T) {
	// Test data
	volume := &spec.ServiceVolumeConfig{
		Source: "./volumes/frontend-react",
	}
	service := &spec.ServiceConfig{}
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "other-service",
					Volumes: []spec.ServiceVolumeConfig{
						{
							Source: "./volumes/../volumes/frontend-react",
						},
						{
							Source: "../random-other.file",
						},
					},
				},
				{
					Name: "current-service",
				},
			},
		},
	}
	// Execute test
	result := isVolumeUsedByOtherServices(volume, service, project)
	// Assert
	assert.True(t, result)
}

func TestIsVolumeUsedByOtherServices2(t *testing.T) {
	// Test data
	volume := &spec.ServiceVolumeConfig{
		Source: "./volumes/database-orientdb",
	}
	service := &spec.ServiceConfig{
		Name: "current-service",
	}
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "other-service",
					Volumes: []spec.ServiceVolumeConfig{
						{
							Source: "../random-other.file",
						},
						{
							Source: "./volumes/database-orientdb/config",
						},
					},
				},
				{
					Name: "current-service",
				},
			},
		},
	}
	// Execute test
	result := isVolumeUsedByOtherServices(volume, service, project)
	// Assert
	assert.True(t, result)
}

func TestIsVolumeUsedByOtherServices3(t *testing.T) {
	// Test data
	volume := &spec.ServiceVolumeConfig{
		Source: "./volumes/backend-gin",
	}
	service := &spec.ServiceConfig{}
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "other-service",
				},
				{
					Name: "current-service",
				},
			},
		},
	}
	// Execute test
	result := isVolumeUsedByOtherServices(volume, service, project)
	// Assert
	assert.False(t, result)
}
