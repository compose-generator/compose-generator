// go:build linux
package pass

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstallDocker1(t *testing.T) {
	// Test data
	filePath := os.TempDir() + "/install-docker.sh"
	// Mock functions
	IsPrivileged = func() bool {
		return true
	}
	P = func(text string) {
		assert.Equal(t, "Installing Docker ... ", text)
	}
	executeWaitCallCount := 0
	ExecuteAndWait = func(c ...string) {
		executeWaitCallCount++
		if executeWaitCallCount == 1 {
			assert.EqualValues(t, []string{"chmod", "+x", filePath}, c)
		} else {
			assert.EqualValues(t, []string{"sh", filePath}, c)
		}
	}
	DownloadFile = func(url string, filepath string) error {
		assert.Equal(t, downloadUrl, url)
		assert.Equal(t, filePath, filepath)
		return nil
	}
	// Execute test
	InstallDocker()
	// Assert
	assert.Equal(t, 2, executeWaitCallCount)
}

func TestInstallDocker2(t *testing.T) {
	// Test data
	filePath := os.TempDir() + "/install-docker.sh"
	errorMessage := "Test download error"
	// Mock functions
	IsPrivileged = func() bool {
		return true
	}
	pCallCount := 0
	P = func(text string) {
		pCallCount++
	}
	Error = func(description string, err error, exit bool) {
		assert.Equal(t, "Download of Docker install script failed", description)
		assert.NotNil(t, err)
		assert.Equal(t, errorMessage, err.Error())
		assert.True(t, exit)
	}
	DownloadFile = func(url string, filepath string) error {
		assert.Equal(t, downloadUrl, url)
		assert.Equal(t, filePath, filepath)
		return errors.New(errorMessage)
	}
	ExecuteAndWait = func(c ...string) {}
	// Execute test
	InstallDocker()
	// Assert
	assert.Equal(t, 1, pCallCount)
}

func TestInstallDocker3(t *testing.T) {
	// Mock functions
	IsPrivileged = func() bool {
		return false
	}
	P = func(text string) {
		assert.Fail(t, "Unexpected call of P")
	}
	// Execute test
	InstallDocker()
}
