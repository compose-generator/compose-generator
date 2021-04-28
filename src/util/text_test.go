package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------ GetErrorMessage ------------------------------------------

func TestGetErrorMessage_ErrorNil(t *testing.T) {
	result := getErrorMessage("This is a error message", nil)
	assert.Equal(t, "Error: This is a error message", result)
}

func TestGetErrorMessage_ErrorNotNil(t *testing.T) {
	err := errors.New("Test error")
	result := getErrorMessage("This is a error message", err)
	assert.Equal(t, "Error: This is a error message: Test error", result)
}
