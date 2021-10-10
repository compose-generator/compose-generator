/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// -------------------------------------------------------------------- Install --------------------------------------------------------------------

func TestInstall1(t *testing.T) {
	// Test data
	expectedDockerVersion := "Docker version 20.10.8, build 3967b7d"
	// Mock functions
	isDockerizedEnvironment = func() bool {
		return false
	}
	installDockerPassCallCount := 0
	installDockerPass = func() {
		installDockerPassCallCount++
	}
	commandExists = func(cmd string) bool {
		assert.Equal(t, "docker", cmd)
		return true
	}
	getDockerVersion = func() string {
		return expectedDockerVersion
	}
	printSuccessMessageCallCount := 0
	printSuccessMessage = func(text string) {
		printSuccessMessageCallCount++
		assert.Equal(t, "Congrats! You have installed "+expectedDockerVersion+". You now can start by executing 'compose-generator generate' to generate your compose file.", text)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	// Execute test
	result := Install(nil)
	// Assert
	assert.Nil(t, result)
	assert.Equal(t, 1, installDockerPassCallCount)
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, 1, printSuccessMessageCallCount)
}

func TestInstall2(t *testing.T) {
	// Mock functions
	isDockerizedEnvironment = func() bool {
		return true
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "You are currently using the dockerized version of Compose Generator. To use this command, please install Compose Generator on your system. Visit https://www.compose-generator.com/install/linux or https://www.compose-generator.com/install/windows for more details.", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	result := Install(nil)
	// Assert
	assert.Nil(t, result)
	assert.Equal(t, 1, printErrorCallCount)
}

func TestInstall3(t *testing.T) {
	// Mock functions
	isDockerizedEnvironment = func() bool {
		return false
	}
	installDockerPassCallCount := 0
	installDockerPass = func() {
		installDockerPassCallCount++
	}
	commandExists = func(cmd string) bool {
		assert.Equal(t, "docker", cmd)
		return false
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "An error occurred while installing Docker", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	result := Install(nil)
	// Assert
	assert.Nil(t, result)
	assert.Equal(t, 1, installDockerPassCallCount)
	assert.Zero(t, pelCallCount)
}
