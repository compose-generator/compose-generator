package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------- Command Exists ---------------------------------------------------------------

func TestCommandExists_Succesful(t *testing.T) {
	result := CommandExists("ls")
	assert.True(t, result)
}

func TestCommandExists_Failure(t *testing.T) {
	result := CommandExists("asdgausegksk")
	assert.False(t, result)
}
