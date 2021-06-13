package util

import (
	"fmt"
	"strings"
)

// nolint: gochecknoglobals
var (
	Version = "dev"
	Commit  = ""
	Date    = ""
	BuiltBy = ""
)

func BuildVersion(version, commit, date, builtBy string) string {
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s, commit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s, built at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s, built by: %s", result, builtBy)
	}
	return result
}

// IsDevVersion checks if this version of Compose Generator is a development version
func IsDevVersion() bool {
	return Version == "dev"
}

// IsPreRelease checks if this version of Compose Generator is a pre-release version
func IsPreRelease() bool {
	return strings.Contains(Version, "alpha") || strings.Contains(Version, "beta") || strings.Contains(Version, "rc")
}
