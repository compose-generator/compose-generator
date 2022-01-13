/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"strings"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// ReplaceVarsInString replaces all variables in the stated string with the contents of the map
func ReplaceVarsInString(content string, varMap map[string]string) string {
	for key, value := range varMap {
		content = strings.ReplaceAll(content, "${{"+key+"}}", value)
	}
	return content
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
