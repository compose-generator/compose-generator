package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDevVersion(t *testing.T) {
	result := IsDevVersion()
	if strings.HasSuffix(VERSION, "-dev") {
		assert.True(t, result)
	} else {
		assert.False(t, result)
	}
}
