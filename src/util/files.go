/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"os"
	"strings"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir checks if a file is a directory
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// NormalizePaths returns the passed paths, whithout any duplicates and nested paths
func NormalizePaths(paths []string) []string {
	normalizedPaths := []string{}
	for _, path := range paths {
		// Check for duplicate
		duplicate := false
		for _, normalizedPath := range normalizedPaths {
			if path == normalizedPath {
				duplicate = true
				break
			}
		}
		if duplicate {
			continue
		}
		// Check if nested in other paths
		containedInOtherPath := false
		for _, otherPath := range paths {
			// Skip the current path
			if path == otherPath {
				continue
			}
			// Check if the current path is nested in another path
			if strings.HasPrefix(path, otherPath) {
				containedInOtherPath = true
				break
			}
		}
		// Add to normalized list if not contained anywhere
		if !containedInOtherPath {
			normalizedPaths = append(normalizedPaths, path)
		}
	}
	return normalizedPaths
}
