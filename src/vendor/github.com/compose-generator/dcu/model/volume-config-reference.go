package model

// VolumeConfigurationReference represents the YAML structure of a volume configuration in a docker compose file
type VolumeConfigurationReference struct {
	Name     string   `yaml:"name,omitempty"`
	External bool     `yaml:"external,omitempty"`
	Driver   string   `yaml:"driver,omitempty"`
	Labels   []string `yaml:"labels,omitempty"`
}
