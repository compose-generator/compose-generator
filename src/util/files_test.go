package util

import (
	"io/ioutil"
	"os"
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

// ------------------------------------------ AddFileToGitignore ------------------------------------------

func TestAddFileToGitignore_NotExisting(t *testing.T) {
	ignoredPath := "./environment-test-file.env"
	// Delete a potentially existing gitignore file
	os.Remove(".gitignore")
	// Execute method
	AddFileToGitignore(ignoredPath)
	content, err := ioutil.ReadFile(".gitignore")
	// Remove output file
	os.Remove(".gitignore")
	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, "# Docker secrets\n"+ignoredPath, string(content))
}

func TestAddFileToGitignore_Existing(t *testing.T) {
	ignoredPath := "./environment-test-file.env"
	// Prepare existing gitignore file
	initialContent := "./demo-dir/*"
	ioutil.WriteFile(".gitignore", []byte(initialContent), 0755)
	// Execute method
	AddFileToGitignore(ignoredPath)
	content, err := ioutil.ReadFile(".gitignore")
	// Remove output file
	os.Remove(".gitignore")
	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, initialContent+"\n# Docker secrets\n"+ignoredPath, string(content))
}

func TestAddFileToGitignore_PathAlreadyIncluded(t *testing.T) {
	ignoredPath := "./environment-test-file.env"
	// Prepare existing gitignore file
	initialContent := "./demo-dir/*\n" + ignoredPath
	ioutil.WriteFile(".gitignore", []byte(initialContent), 0755)
	// Execute method
	AddFileToGitignore(ignoredPath)
	content, err := ioutil.ReadFile(".gitignore")
	// Remove output file
	os.Remove(".gitignore")
	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, initialContent, string(content))
}
