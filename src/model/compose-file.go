package model

// ComposeFile represents the YAML structure of docker compose file
type ComposeFile struct {
	Version  string
	Services map[string]Service
	Networks map[string]Network `yaml:"networks,omitempty"`
}

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
}

// Network represents the YAML structure of a network configuration in a docker compose file
type Network struct {
	External ExtneralNetwork `yaml:"external,omitempty"`
	Ipam     IPAMNetwork     `yaml:"ipam,omitempty"`
}

// IPAMNetwork represents the YAML structure of an ipam network configuration in a docker compose file
type IPAMNetwork struct {
	Driver     string            `yaml:"driver,omitempty"`
	DriverOpts map[string]string `yaml:"driver_opts,omitempty"`
	Config     []string          `yaml:"config,omitempty"`
}

// ExtneralNetwork represents the YAML structure of an external network configuration in a docker compose file
type ExtneralNetwork struct {
	Name string `yaml:"name,omitempty"`
}
