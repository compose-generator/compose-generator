package project

import (
	"path/filepath"
	"strings"
)

// SaveOptions represents an option to the SaveProject function
type SaveOptions struct {
	WorkingDir      string
	ComposeFileName string
}

// SaveOption represents a callback function option for the SaveProject function
type SaveOption func(*SaveOptions)

func applySaveOptions(options ...SaveOption) SaveOptions {
	// Create options with default values
	opts := SaveOptions{
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
	return opts
}

// SaveWithComposeFile is an option to set a custom compose file name
func SaveWithComposeFile(value string) SaveOption {
	return func(o *SaveOptions) {
		o.ComposeFileName = value
	}
}

// SaveIntoDir is an option to set a custom dir to save the project to
func SaveIntoDir(value string) SaveOption {
	return func(o *SaveOptions) {
		o.WorkingDir = value
	}
}
