/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package util

import (
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

// ------------------------------------------------------------ GetToolboxImageVersion -----------------------------------------------------------

func TestGetToolboxImageVersion(t *testing.T) {
	result := getToolboxImageVersion()
	if Version == "dev" {
		assert.Equal(t, "dev", result)
	} else {
		assert.Equal(t, Version, result)
	}
}
