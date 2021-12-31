/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package model

import (
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------- GetAllVolumePaths -------------------------------------------------------------

func TestGetAllVolumePaths1(t *testing.T) {
	// Test data
	project := &CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Volumes: []spec.ServiceVolumeConfig{
						{
							Type:   spec.VolumeTypeBind,
							Source: "./volumes/mysql-data",
							Target: "/mysql/data",
						},
						{
							Type:   spec.VolumeTypeVolume,
							Source: "postgres-data",
							Target: "/postgres/data",
						},
					},
				},
				{
					Volumes: []spec.ServiceVolumeConfig{
						{
							Type:   spec.VolumeTypeBind,
							Source: "./volumes/wordpress-content",
							Target: "/wordpress/content",
						},
						{
							Type:   spec.VolumeTypeBind,
							Source: "/bin/compose-generator",
							Target: "/cg/compose-generator",
						},
					},
				},
			},
		},
	}
	expectedResult := []string{"./volumes/mysql-data", "./volumes/wordpress-content", "/bin/compose-generator"}
	// Execute test
	result := project.GetAllVolumePaths()
	// Assert
	assert.Equal(t, 3, len(result))
	assert.EqualValues(t, expectedResult, result)
}

func TestGetAllVolumePaths2(t *testing.T) {
	// Test data
	project := CGProject{}
	// Execute test
	result := project.GetAllVolumePaths()
	// Assert
	assert.Zero(t, len(result))
}

// --------------------------------------------------------------- GetAllBuildContextPaths ---------------------------------------------------------

func TestGetAllBuildContextPaths(t *testing.T) {
	// Test data
	project := &CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Build: &spec.BuildConfig{
						Context: "./context",
					},
				},
				{
					Build: &spec.BuildConfig{
						Context: "./build-context",
					},
				},
			},
		},
	}
	expectedResult := []string{"./context", "./build-context"}
	// Execute test
	result := project.GetAllBuildContextPaths()
	// Assert
	assert.EqualValues(t, expectedResult, result)
}

// ----------------------------------------------------------------- GetAllEnvFilePaths ------------------------------------------------------------

func TestGetAllEnvFilePaths1(t *testing.T) {
	// Test data
	project := &CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					EnvFile: []string{"./environment.env", "./volumes/wordpress/environment.env"},
				},
				{
					EnvFile: []string{"./environment.env"},
				},
			},
		},
	}
	expectedResult := []string{"./environment.env", "./volumes/wordpress/environment.env", "./environment.env"}
	// Execute test
	result := project.GetAllEnvFilePaths()
	// Assert
	assert.EqualValues(t, expectedResult, result)
}

func TestGetAllEnvFilePaths2(t *testing.T) {
	// Test data
	project := CGProject{}
	// Execute test
	result := project.GetAllEnvFilePaths()
	// Assert
	assert.Zero(t, len(result))
}

// ------------------------------------------------------------ GetAllEnvFilePathsNormalized -------------------------------------------------------

func TestGetAllEnvFilePathsNormalized1(t *testing.T) {
	// Test data
	project := &CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					EnvFile: []string{"./environment.env", "./volumes/wordpress/environment.env"},
				},
				{
					EnvFile: []string{"./environment.env"},
				},
			},
		},
	}
	expectedResult := []string{"./environment.env", "./volumes/wordpress/environment.env"}
	// Execute test
	result := project.GetAllEnvFilePathsNormalized()
	// Assert
	assert.Equal(t, 2, len(result))
	assert.EqualValues(t, expectedResult, result)
}

func TestGetAllEnvFilePathsNormalized2(t *testing.T) {
	// Test data
	project := CGProject{}
	// Execute test
	result := project.GetAllEnvFilePathsNormalized()
	// Assert
	assert.Zero(t, len(result))
}

// ------------------------------------------------------------------ GetServiceRef ----------------------------------------------------------------

func TestGetServiceRef1(t *testing.T) {
	// Test data
	project := &CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "Service 1",
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
	serviceName := "Service 2"
	// Execute test
	result := project.GetServiceRef(serviceName)
	// Assert
	assert.Equal(t, &project.Composition.Services[1], result)
}

func TestGetServiceRef2(t *testing.T) {
	// Test data
	project := &CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "Service 1",
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
	serviceName := "Service 4"
	// Execute test
	result := project.GetServiceRef(serviceName)
	// Assert
	assert.Nil(t, result)
}
