package model

// GenerateConfig represents a configuration file, which can be passed to the generate command
type GenerateConfig struct {
	ProjectName    string          `yaml:"project_name"`
	ComposeVersion string          `yaml:"compose_version,omitempty"`
	AlsoProduction bool            `yaml:"also_production,omitempty"`
	ServiceConfig  []ServiceConfig `yaml:"services,omitempty"`
}

// ServiceConfig represents a collection of services within a GenerateConfig
type ServiceConfig struct {
	Type    string            `yaml:"type"`
	Service string            `yaml:"service"`
	Params  map[string]string `yaml:"params,omitempty"`
}
