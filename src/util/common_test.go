package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------ AppendStringToSliceIfMissing ------------------------------------------

func AppendStringToSliceIfMissing_Missing(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	slice = AppendStringToSliceIfMissing(slice, "ipsum")
	assert.Equal(t, 5, len(slice))
}

func AppendStringToSliceIfMissing_NotMissing(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	slice = AppendStringToSliceIfMissing(slice, "bar")
	assert.Equal(t, 4, len(slice))
}

// ------------------------------------------ TestSliceContainsString ------------------------------------------

func TestSliceContainsString_True(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	result := SliceContainsString(slice, "bar")
	assert.True(t, result)
}

func TestSliceContainsString_False(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	result := SliceContainsString(slice, "ipsum")
	assert.False(t, result)
}

// ------------------------------------------ TestSliceContainsInt ------------------------------------------

func TestSliceContainsInt_True(t *testing.T) {
	slice := []int{5, 1, 78, 6}
	result := SliceContainsInt(slice, 6)
	assert.True(t, result)
}

func TestSliceContainsInt_False(t *testing.T) {
	slice := []int{5, 1, 78, 6}
	result := SliceContainsInt(slice, 9)
	assert.False(t, result)
}
