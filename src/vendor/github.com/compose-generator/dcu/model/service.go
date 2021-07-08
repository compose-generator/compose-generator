package model

// Service represents the YAML structure of a service in a docker compose file
type Service struct {
	Build         string   `yaml:"build,omitempty"`
	Image         string   `yaml:"image,omitempty"`
	ContainerName string   `yaml:"container_name,omitempty"`
	Volumes       []string `yaml:"volumes,omitempty"`
	Networks      []string `yaml:"networks,omitempty"`
	Ports         []string `yaml:"ports,omitempty"`
	EnvFile       []string `yaml:"env_file,omitempty"`
	Environment   []string `yaml:"environment,omitempty"`
	Restart       string   `yaml:"restart,omitempty"`
	DependsOn     []string `yaml:"depends_on,omitempty"`
	Links         []string `yaml:"links,omitempty"`
	Labels        []string `yaml:"labels,omitempty"`
	Command       string   `yaml:"command,omitempty"`
}
