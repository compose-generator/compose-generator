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

func TestPel(t *testing.T) {
	// Mock functions
	printlnCallCount := 0
	println = func(a ...interface{}) (n int, err error) {
		printlnCallCount++
		assert.Zero(t, len(a))
		return 0, nil
	}
	// Execute test
	Pel()
	// Assert
	assert.Equal(t, 1, printlnCallCount)
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
