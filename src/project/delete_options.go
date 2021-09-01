package project

import (
	"path/filepath"
	"strings"
)

type DeleteOptions struct {
	ComposeFileName string
	WorkingDir      string
}

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
	return opts
}

func DeleteComposeFileName(value string) DeleteOption {
	return func(o *DeleteOptions) {
		o.ComposeFileName = value
	}
}

func DeleteWorkingDir(value string) DeleteOption {
	return func(o *DeleteOptions) {
		o.WorkingDir = value
	}
}
