/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

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

// ------------------------------------------------------------ NormalizePaths --------------------------------------------------------

func TestNormalizePaths(t *testing.T) {
	// Test data
	paths := []string{"./volumes/mysql-data", "/bin/compose-generator", "./volumes/wordpress-content", "/bin/compose-generator", "./volumes/mysql-data/config"}
	expectedResult := []string{"./volumes/mysql-data", "/bin/compose-generator", "./volumes/wordpress-content"}
	// Execute test
	result := NormalizePaths(paths)
	// Assert
	assert.Equal(t, 3, len(result))
	assert.EqualValues(t, expectedResult, result)
}
