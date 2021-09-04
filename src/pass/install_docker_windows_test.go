// go:build windows
package pass

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstallDockers1(t *testing.T) {
	// Test data
	filePath := os.TempDir() + "/DockerInstaller.exe"
	// Mock functions
	P = func(text string) {
		assert.Equal(t, "Downloading Docker installer ... ", text)
	}
	downloadFile = func(url string, filepath string) error {
		assert.Equal(t, downloadUrl, url)
		assert.Equal(t, filePath, filepath)
		return nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Download of Docker installer failed", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	doneCallCount := 0
	done = func() {
		doneCallCount++
	}
	pl = func(text string) {
		assert.Equal(t, "Running installation ... ", text)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	executeWithOutput = func(c string) {
		assert.Equal(t, filePath, c)
	}
	// Execute test
	InstallDocker()
	// Assert
	assert.Equal(t, 1, doneCallCount)
	assert.Equal(t, 2, pelCallCount)
}

func TestInstallDocker2(t *testing.T) {
	// Test data
	filePath := os.TempDir() + "/DockerInstaller.exe"
	errorMessage := "Test download error"
	// Mock functions
	P = func(text string) {
		assert.Equal(t, "Downloading Docker installer ... ", text)
	}
	DownloadFile = func(url string, filepath string) error {
		assert.Equal(t, downloadUrl, url)
		assert.Equal(t, filePath, filepath)
		return errors.New(errorMessage)
	}
	errorCalled := false
	Error = func(description string, err error, exit bool) {
		assert.Equal(t, "Download of Docker installer failed", description)
		assert.NotNil(t, err)
		assert.Equal(t, errorMessage, err.Error())
		assert.True(t, exit)
		errorCalled = true
	}
	// Execute test
	InstallDocker()
	// Assert
	assert.True(t, errorCalled)
}
