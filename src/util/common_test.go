/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------ ReplaceVarsInString ------------------------------------------

func TestReplaceVarsInString(t *testing.T) {
	vars := map[string]string{
		"var1": "ipsum",
		"var2": "consetetur",
		"var3": "sed",
	}
	input := "Lorem ${{var1}} dolor sit amet, ${{var2}} sadipscing elitr, ${{var3}} diam nonumy"
	output := ReplaceVarsInString(input, vars)
	assert.Equal(t, "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy", output)
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
