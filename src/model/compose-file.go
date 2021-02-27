package model

// ComposeFile represents the YAML structure of docker compose file
type ComposeFile struct {
	Version  string
	Services map[string]Service
}

// Service represents the YAML structure of a service in a docker compose file
type Service struct {
	Build         string   `yaml:"build,omitempty"`
	Image         string   `yaml:"image,omitempty"`
	ContainerName string   `yaml:"container_name,omitempty"`
	Restart       string   `yaml:"restart,omitempty"`
	DependsOn     []string `yaml:"depends_on,omitempty"`
	Links         []string `yaml:"links,omitempty"`
	Ports         []string `yaml:"ports,omitempty"`
	Volumes       []string `yaml:"volumes,omitempty"`
	Environment   []string `yaml:"environment,omitempty"`
	EnvFile       []string `yaml:"env_file,omitempty"`
}
