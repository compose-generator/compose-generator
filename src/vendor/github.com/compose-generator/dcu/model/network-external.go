package model

// ExtneralNetwork represents the YAML structure of an external network configuration in a docker compose file
type ExtneralNetwork struct {
	Name string `yaml:"name,omitempty"`
}
