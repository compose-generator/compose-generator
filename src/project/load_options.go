package project

import "strings"

type LoadOptions struct {
	WorkingDir      string
	ComposeFileName string
}

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

func LoadFromComposeFile(value string) LoadOption {
	return func(o *LoadOptions) {
		o.ComposeFileName = value
	}
}

func LoadFromDir(value string) LoadOption {
	return func(o *LoadOptions) {
		o.WorkingDir = value
	}
}
