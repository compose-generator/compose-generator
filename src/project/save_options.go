package project

import (
	"path/filepath"
	"strings"
)

type SaveOptions struct {
	WorkingDir      string
	ComposeFileName string
}

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

func SaveWithComposeFile(value string) SaveOption {
	return func(o *SaveOptions) {
		o.ComposeFileName = value
	}
}

func SaveIntoDir(value string) SaveOption {
	return func(o *SaveOptions) {
		o.WorkingDir = value
	}
}
