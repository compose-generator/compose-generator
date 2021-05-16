package model

// DockerNetwork represents the JSON structure of a docker network
type DockerNetwork struct {
	Name       string
	ID         string
	Created    string
	Scope      string
	Driver     string
	EnableIPv6 bool
	IPAM       IPAM
	Internal   bool
	Attachable bool
	Ingress    bool
	ConfigFrom ConfigFrom
	ConfigOnly bool
	Options    map[string]string
	Labels     map[string]string
}

// IPAM represents the JSON structure of a docker network IPAM configuration
type IPAM struct {
	Driver  string
	Options map[string]string
	Config  []IPAMConfig
}

// IPAMConfig represents the JSON structure of a docker network IPAM configuration
type IPAMConfig struct {
	Subnet  string
	Gateway string
}

// ConfigFrom represents the JSON structure of a docker network ConfigFrom configuration
type ConfigFrom struct {
	Network string
}
