/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package project

import (
	"path/filepath"
	"strings"
)

// DeleteOptions represents an option to the DeleteProject function
type DeleteOptions struct {
	ComposeFileName string
	WorkingDir      string
}

// DeleteOption represents a callback function option for the DeleteProject function
type DeleteOption func(*DeleteOptions)

func applyDeleteOptions(options ...DeleteOption) DeleteOptions {
	// Create options with default values
	opts := DeleteOptions{
		WorkingDir:      "./",
		ComposeFileName: "docker-compose.yml",
	}
	// Apply custom options
	for _, opt := range options {
		opt(&opts)
	}
	// Validate and corrent the passed options
	opts.WorkingDir = filepath.ToSlash(opts.WorkingDir)
	if !strings.HasSuffix(opts.WorkingDir, "/") {
		opts.WorkingDir += "/"
	}
	// Check also docker-compose.yaml file
	if opts.ComposeFileName == "docker-compose.yml" && fileExists(opts.WorkingDir+"docker-compose.yaml") {
		opts.ComposeFileName = "docker-compose.yaml"
	}
	return opts
}

// DeleteComposeFileName is an option to set a custom compose file name
func DeleteComposeFileName(value string) DeleteOption {
	return func(o *DeleteOptions) {
		o.ComposeFileName = value
	}
}

// DeleteWorkingDir is an option to set a custom work dir
func DeleteWorkingDir(value string) DeleteOption {
	return func(o *DeleteOptions) {
		o.WorkingDir = value
	}
}
