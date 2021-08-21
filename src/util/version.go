package util

import (
	"fmt"
	"runtime/debug"
	"strings"
)

// nolint: gochecknoglobals
var (
	Version = "dev"
	Commit  = ""
	Date    = ""
	BuiltBy = ""
)

// BuildVersion returns a version string for the current verison, including commit id, date and user
func BuildVersion(version, commit, date, builtBy string) string {
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nmodule version: %s, checksum: %s", result, info.Main.Version, info.Main.Sum)
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
