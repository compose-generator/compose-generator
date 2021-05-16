package model

// IPAMNetwork represents the YAML structure of an ipam network configuration in a docker compose file
type IPAMNetwork struct {
	Driver     string            `yaml:"driver,omitempty"`
	DriverOpts map[string]string `yaml:"driver_opts,omitempty"`
	Config     []string          `yaml:"config,omitempty"`
}
