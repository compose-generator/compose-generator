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
