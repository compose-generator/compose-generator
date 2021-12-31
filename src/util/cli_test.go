/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------- Command Exists ---------------------------------------------------------------

func TestCommandExists_Successful(t *testing.T) {
	result := CommandExists("curl")
	assert.True(t, result)
}

func TestCommandExists_Failure(t *testing.T) {
	result := CommandExists("asdgausegksk")
	assert.False(t, result)
}

// ------------------------------------------------------------ getToolboxImageVersion -----------------------------------------------------------

func TestGetToolboxImageVersion(t *testing.T) {
	result := getToolboxImageVersion()
	if Version == "dev" {
		assert.Equal(t, "dev", result)
	} else {
		assert.Equal(t, Version, result)
	}
}

// -------------------------------------------------------------- getToolboxMountPath ------------------------------------------------------------

func TestGetToolboxMountPath1(t *testing.T) {
	// Test data
	outerVolumePath := "/docker/outer/volume/path"
	// Mock functions
	isDockerizedEnvironment = func() bool {
		return true
	}
	getOuterVolumePathOnDockerizedEnvironmentMockable = func() string {
		return outerVolumePath
	}
	getwd = func() (dir string, err error) {
		assert.Fail(t, "Unexpected call of getwd")
		return "", nil
	}
	// Execute test
	result := getToolboxMountPath()
	// Assert
	assert.Equal(t, outerVolumePath, result)
}

func TestGetToolboxMountPath2(t *testing.T) {
	// Test data
	wd := "/docker/outer/volume/path"
	// Mock functions
	isDockerizedEnvironment = func() bool {
		return false
	}
	getOuterVolumePathOnDockerizedEnvironmentMockable = func() string {
		assert.Fail(t, "Unexpected call of getOuterVolumePathOnDockerizedEnvironment")
		return ""
	}
	getwd = func() (dir string, err error) {
		return wd, nil
	}
	logError = func(message string, exit bool) {
		assert.Fail(t, "Unexpected call of logError")
	}
	// Execute test
	result := getToolboxMountPath()
	// Assert
	assert.Equal(t, wd, result)
}

func TestGetToolboxMountPath3(t *testing.T) {
	// Mock functions
	isDockerizedEnvironment = func() bool {
		return false
	}
	getOuterVolumePathOnDockerizedEnvironmentMockable = func() string {
		assert.Fail(t, "Unexpected call of getOuterVolumePathOnDockerizedEnvironment")
		return ""
	}
	getwd = func() (dir string, err error) {
		return "", errors.New("Error message")
	}
	logErrorCallCount := 0
	logError = func(message string, exit bool) {
		logErrorCallCount++
		assert.Equal(t, "Could not find current working directory", message)
		assert.True(t, exit)
	}
	// Execute test
	result := getToolboxMountPath()
	// Assert
	assert.Zero(t, len(result))
	assert.Equal(t, 1, logErrorCallCount)
}
