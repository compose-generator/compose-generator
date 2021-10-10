/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package project

import "strings"

// LoadOptions represents an option to the LoadProject function
type LoadOptions struct {
	WorkingDir      string
	ComposeFileName string
}

// LoadOption represents a callback function option for the LoadProject function
type LoadOption func(*LoadOptions)

func applyLoadOptions(options ...LoadOption) LoadOptions {
	// Create options with default values
	opts := LoadOptions{
		WorkingDir:      "./",
		ComposeFileName: "docker-compose.yml",
	}
	// Apply custom options
	for _, opt := range options {
		opt(&opts)
	}
	// Validate and corrent the passed options
	opts.WorkingDir = strings.ReplaceAll(opts.WorkingDir, "\\", "/")
	if !strings.HasSuffix(opts.WorkingDir, "/") {
		opts.WorkingDir += "/"
	}
	return opts
}

// LoadFromComposeFile is an option to set a custom compose file name
func LoadFromComposeFile(value string) LoadOption {
	return func(o *LoadOptions) {
		o.ComposeFileName = value
	}
}

// LoadFromDir is an option to set a dir where to load the project from
func LoadFromDir(value string) LoadOption {
	return func(o *LoadOptions) {
		o.WorkingDir = value
	}
}
