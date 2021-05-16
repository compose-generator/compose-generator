package model

// ComposeFile represents the YAML structure of docker compose file
type ComposeFile struct {
	Version  string
	Services map[string]Service
	Networks map[string]NetworkConfigurationReference `yaml:"networks,omitempty"`
	Volumes  map[string]VolumeConfigurationReference  `yaml:"volumes,omitempty"`
}
