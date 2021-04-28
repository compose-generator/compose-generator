package util

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------ ReplaceVarsInFile ------------------------------------------

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

// ------------------------------------------ GenerateSecrets ------------------------------------------

func TestGenerateSecretsAndReplaceInString(t *testing.T) {
	var secrets = []model.Secret{
		{
			Name:     "MySQL root password",
			Variable: "_PW_MYSQL_ROOT",
			Length:   50,
		},
		{
			Name:     "MySQL application password",
			Variable: "_PW_MYSQL_APPLICATION",
			Length:   30,
		},
	}
	content := "Here goes the root pw: ${{_PW_MYSQL_ROOT}}\nand here goes the app pw: ${{_PW_MYSQL_APPLICATION}}"
	secretsMap := generateSecretsAndReplaceInString(&content, secrets)
	assert.Equal(t, 50, len(secretsMap["MySQL root password"]))
	assert.Equal(t, 30, len(secretsMap["MySQL application password"]))
}

// ------------------------------------------ AppendStringToSliceIfMissing ------------------------------------------

func TestAppendStringToSliceIfMissing_Missing(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	slice = AppendStringToSliceIfMissing(slice, "ipsum")
	assert.Equal(t, 5, len(slice))
}

func TestAppendStringToSliceIfMissing_NotMissing(t *testing.T) {
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

// ------------------------------------------ RemoveStringFromSlice ------------------------------------------

func TestRemoveStringFromSlice_Found(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	result := RemoveStringFromSlice(slice, "foo")
	assert.Equal(t, 3, len(result))
}

func TestRemoveStringFromSlice_NotFound(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	result := RemoveStringFromSlice(slice, "abc")
	assert.Equal(t, 4, len(result))
}
