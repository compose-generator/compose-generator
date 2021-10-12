/*
Copyright © 2021 Compose Generator Contributors
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
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
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
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Could not find current working directory", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	result := getToolboxMountPath()
	// Assert
	assert.Zero(t, len(result))
	assert.Equal(t, 1, printErrorCallCount)
}
