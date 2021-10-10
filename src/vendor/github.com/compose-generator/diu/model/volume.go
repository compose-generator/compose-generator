/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package model

// DockerVolume represents the JSON structure of a docker volume
type DockerVolume struct {
	CreatedAt  string
	Driver     string
	Labels     []string
	Mountpoint string
	Name       string
	Options    []string
	Scope      string
}
