package util

import (
	"compose-generator/model"
	"io/ioutil"
	"strings"

	"github.com/sethvargo/go-password/password"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// ReplaceVarsInFile replaces all variables in the stated file with the contents of the map
func ReplaceVarsInFile(path string, varMap map[string]string) {
	// Read file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		Error("Could not read from "+path, err, true)
	}

	// Replace variables
	content = []byte(ReplaceVarsInString(string(content), varMap))

	// Write content back
	err = ioutil.WriteFile(path, content, 0777)
	if err != nil {
		Error("Could not write to "+path, err, true)
	}
}

// ReplaceVarsInString replaces all variables in the stated string with the contents of the map
func ReplaceVarsInString(content string, varMap map[string]string) string {
	for key, value := range varMap {
		content = strings.ReplaceAll(content, "${{"+key+"}}", value)
	}
	return content
}

// GenerateSecrets generates random strings as secrets and replaces them in the stated file
func GenerateSecrets(path string, secrets []model.Secret) map[string]string {
	// Read file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		Error("Could not read from "+path, err, true)
	}

	// Generate a password for each occurrence of _GENERATE_PW
	newContent := string(content)
	secretsMap := make(map[string]string)
	for _, s := range secrets {
		res, err := password.Generate(s.Length, 10, 0, false, false)
		if err != nil {
			Error("Password generation failed.", err, true)
		}
		newContent = strings.ReplaceAll(newContent, "${{"+s.Variable+"}}", res)
		secretsMap[s.Name] = res
	}

	// Write content back
	err = ioutil.WriteFile(path, []byte(newContent), 0777)
	if err != nil {
		Error("Could not write to "+path, err, true)
	}

	return secretsMap
}

// RemoveStringFromSlice searches a string in a slice and removes it
func RemoveStringFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// SliceContainsString checks if a slice contains a certain element
func SliceContainsString(slice []string, i string) bool {
	for _, ele := range slice {
		if ele == i {
			return true
		}
	}
	return false
}

// SliceContainsInt checks if a slice contains a certain element
func SliceContainsInt(slice []int, i int) bool {
	for _, ele := range slice {
		if ele == i {
			return true
		}
	}
	return false
}

// AppendStringToSliceIfMissing checks if a slice contains a string and adds it if its not existing already
func AppendStringToSliceIfMissing(slice []string, i string) []string {
	if SliceContainsString(slice, i) {
		return slice
	}
	return append(slice, i)
}