/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

// go:build linux
package pass

import (
	"errors"
	"os"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/stretchr/testify/assert"
)

func TestInstallDocker1(t *testing.T) {
	// Test data
	filePath := os.TempDir() + "/install-docker.sh"
	// Mock functions
	isPrivileged = func() bool {
		return true
	}
	startProcess = func(text string) (s *spinner.Spinner) {
		assert.Equal(t, "Installing Docker ...", text)
		return nil
	}
	stopProcessCallCount := 0
	stopProcess = func(s *spinner.Spinner) {
		stopProcessCallCount++
	}
	executeWaitCallCount := 0
	executeAndWait = func(c ...string) {
		executeWaitCallCount++
		if executeWaitCallCount == 1 {
			assert.EqualValues(t, []string{"chmod", "+x", filePath}, c)
		} else {
			assert.EqualValues(t, []string{"sh", filePath}, c)
		}
	}
	downloadFile = func(url string, filepath string) error {
		assert.Equal(t, downloadUrl, url)
		assert.Equal(t, filePath, filepath)
		return nil
	}
	// Execute test
	InstallDocker()
	// Assert
	assert.Equal(t, 2, executeWaitCallCount)
	assert.Equal(t, 1, stopProcessCallCount)
}

func TestInstallDocker2(t *testing.T) {
	// Test data
	filePath := os.TempDir() + "/install-docker.sh"
	// Mock functions
	isPrivileged = func() bool {
		return true
	}
	startProcessCallCount := 0
	startProcess = func(text string) (s *spinner.Spinner) {
		startProcessCallCount++
		return nil
	}
	stopProcessCallCount := 0
	stopProcess = func(s *spinner.Spinner) {
		stopProcessCallCount++
	}
	logError = func(message string, exit bool) {
		assert.Equal(t, "Download of Docker install script failed", message)
		assert.True(t, exit)
	}
	downloadFile = func(url string, filepath string) error {
		assert.Equal(t, downloadUrl, url)
		assert.Equal(t, filePath, filepath)
		return errors.New("Error message")
	}
	executeAndWait = func(c ...string) {}
	// Execute test
	InstallDocker()
	// Assert
	assert.Equal(t, 1, startProcessCallCount)
	assert.Equal(t, 1, stopProcessCallCount)
}

func TestInstallDocker3(t *testing.T) {
	// Mock functions
	isPrivileged = func() bool {
		return false
	}
	startProcess = func(text string) (s *spinner.Spinner) {
		assert.Fail(t, "Unexpected call of startProcess")
		return nil
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Fail(t, "Unexpected call of stopProcess")
	}
	// Execute test
	InstallDocker()
}
