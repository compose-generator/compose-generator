/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------- P -------------------------------------------------

func TestP(t *testing.T) {
	// Test data

}

// ------------------------------------------------- Pl ------------------------------------------------

func TestPl(t *testing.T) {
	// Test data
	text := "This is a test."
	// Mock functions
	whiteCallCount := 0
	white = func(format string, a ...interface{}) {
		whiteCallCount++
		assert.Equal(t, text, format)
	}
	// Execute test
	Pl(text)
	// Assert
	assert.Equal(t, 1, whiteCallCount)
}

// ------------------------------------------------ Pel ------------------------------------------------

func TestPel1(t *testing.T) {
	// Mock functions
	printlnCallCount := 0
	println = func(a ...interface{}) (n int, err error) {
		printlnCallCount++
		assert.Zero(t, len(a))
		return 0, nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	Pel()
	// Assert
	assert.Equal(t, 1, printlnCallCount)
}

func TestPel2(t *testing.T) {
	// Mock functions
	printlnCallCount := 0
	println = func(a ...interface{}) (n int, err error) {
		printlnCallCount++
		assert.Zero(t, len(a))
		return 0, errors.New("Test")
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Could not print empty line", description)
		assert.NotNil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	Pel()
	// Assert
	assert.Equal(t, 1, printlnCallCount)
}

// ----------------------------------------------- Error -----------------------------------------------

func TestError1(t *testing.T) {
	// Test data
	description := "Error message"
	errorMessage := "Error: " + description + " - Error message"
	err := errors.New("Error message")
	// Mock functions
	redCallCount := 0
	red = func(format string, a ...interface{}) {
		redCallCount++
		assert.Equal(t, errorMessage, format)
	}
	getErrorMessageMockable = func(desc string, err error) string {
		assert.Equal(t, description, desc)
		assert.NotNil(t, err)
		return errorMessage
	}
	exitProgramCallCount := 0
	exitProgram = func(code int) {
		exitProgramCallCount++
		assert.Equal(t, 1, code)
	}
	// Execute test
	Error(description, err, true)
	// Assert
	assert.Equal(t, 1, redCallCount)
	assert.Equal(t, 1, exitProgramCallCount)
}

func TestError2(t *testing.T) {
	// Test data
	description := "Error message"
	errorMessage := "Error: " + description + " - Error message"
	err := errors.New("Error message")
	// Mock functions
	redCallCount := 0
	red = func(format string, a ...interface{}) {
		redCallCount++
		assert.Equal(t, errorMessage, format)
	}
	getErrorMessageMockable = func(desc string, err error) string {
		assert.Equal(t, description, desc)
		assert.NotNil(t, err)
		return errorMessage
	}
	exitProgram = func(code int) {
		assert.Fail(t, "Unexpected call of exitProgram")
	}
	// Execute test
	Error(description, err, false)
	// Assert
	assert.Equal(t, 1, redCallCount)
}

// ---------------------------------------------- Warning ----------------------------------------------

func TestWarning(t *testing.T) {
	// Test data
	description := "This is a warning."
	// Mock functions
	hiYellowCallCount := 0
	hiYellow = func(format string, a ...interface{}) {
		hiYellowCallCount++
		assert.Equal(t, "Warning: This is a warning.", format)
	}
	// Execute test
	Warning(description)
	// Assert
	assert.Equal(t, 1, hiYellowCallCount)
}

// ------------------------------------------ GetErrorMessage ------------------------------------------

func TestGetErrorMessage_ErrorNil(t *testing.T) {
	// Execute test
	result := getErrorMessage("This is a error message", nil)
	// Assert
	assert.Equal(t, "Error: This is a error message", result)
}

func TestGetErrorMessage_ErrorNotNil(t *testing.T) {
	// Test data
	err := errors.New("Test error")
	// Execute test
	result := getErrorMessage("This is a error message", err)
	// Assert
	assert.Equal(t, "Error: This is a error message: Test error", result)
}
