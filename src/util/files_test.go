package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------ FileExists ------------------------------------------

func TestFileExists_Success(t *testing.T) {
	path := "../../README.md"
	assert.True(t, FileExists(path))
}

func TestFileExists_Failure(t *testing.T) {
	path := "../../non-existing.file"
	assert.False(t, FileExists(path))
}

// ------------------------------------------ IsDir ------------------------------------------

func TestIsDir_Success(t *testing.T) {
	path := "../../src"
	assert.True(t, IsDir(path))
}

func TestIsDir_Failure(t *testing.T) {
	path := "../../README.md"
	assert.False(t, IsDir(path))
}
