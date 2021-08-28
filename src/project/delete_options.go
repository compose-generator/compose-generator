package project

type DeleteOptions struct {
	ComposeFileName string
}

type DeleteOption func(*DeleteOptions)

func applyDeleteOptions(options ...DeleteOption) DeleteOptions {
	// Create options with default values
	opts := DeleteOptions{
		ComposeFileName: "docker-compose.yml",
	}
	// Apply custom options
	for _, opt := range options {
		opt(&opts)
	}
	return opts
}

func WithComposeFileName(value string) DeleteOption {
	return func(o *DeleteOptions) {
		o.ComposeFileName = value
	}
}
