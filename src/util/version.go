package util

import "strings"

const VERSION = "0.7.0"

// IsDevVersion checks if this version of Compose Generator is a development version
func IsDevVersion() bool {
	return strings.HasSuffix(VERSION, "dev")
}

// IsPreRelease checks if this version of Compose Generator is a pre-release version
func IsPreRelease() bool {
	return strings.Contains(VERSION, "alpha") || strings.Contains(VERSION, "beta") || strings.Contains(VERSION, "rc")
}
