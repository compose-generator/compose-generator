package util

import (
	"compose-generator/model"
	"io/ioutil"
	"strings"

	"github.com/sethvargo/go-password/password"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

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
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		Error("Could not read from "+path, err, true)
	}
	content := string(contentBytes)
	secretsMap := generateSecretsAndReplaceInString(&content, secrets)

	// Write content back
	err = ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		Error("Could not write to "+path, err, true)
	}
	return secretsMap
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

func generateSecretsAndReplaceInString(content *string, secrets []model.Secret) map[string]string {
	var secretsMap = make(map[string]string)
	// Generate a secret and replace it in the content string
	for _, s := range secrets {
		res, err := password.Generate(s.Length, 10, 0, false, false)
		if err != nil {
			Error("Password generation failed.", err, true)
		}
		*content = strings.ReplaceAll(*content, "${{"+s.Variable+"}}", res)
		secretsMap[s.Name] = res
	}
	return secretsMap
}
