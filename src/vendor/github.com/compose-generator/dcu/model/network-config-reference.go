package model

// NetworkConfigurationReference represents the YAML structure of a network configuration in a docker compose file
type NetworkConfigurationReference struct {
	Name       string          `yaml:"name,omitempty"`
	External   ExtneralNetwork `yaml:"external,omitempty"`
	Ipam       IPAMNetwork     `yaml:"ipam,omitempty"`
	Driver     string          `yaml:"driver,omitempty"`
	Attachable bool            `yaml:"attachable,omitempty"`
	Labels     []string        `yaml:"labels,omitempty"`
}
