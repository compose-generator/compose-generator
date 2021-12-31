/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package model

// GenerateConfig represents a configuration file, which can be passed to the generate command
type GenerateConfig struct {
	ProjectName     string          `yaml:"project_name"`
	ProductionReady bool            `yaml:"production_ready,omitempty"`
	ServiceConfig   []ServiceConfig `yaml:"services,omitempty"`
	FromFile        bool
}

// ServiceConfig represents a collection of services within a GenerateConfig
type ServiceConfig struct {
	Type   string            `yaml:"type"`
	Name   string            `yaml:"name"`
	Params map[string]string `yaml:"params,omitempty"`
}

// GetServiceConfigurationsByType returns all specified service configurations by the name of their type
func (c GenerateConfig) GetServiceConfigurationsByType(templateType string) []ServiceConfig {
	services := []ServiceConfig{}
	for _, service := range c.ServiceConfig {
		if service.Type == templateType {
			services = append(services, service)
		}
	}
	return services
}
