/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package model

// DockerManifest represents the JSON structure of a docker manifest
type DockerManifest struct {
	Ref              string
	Descriptor       Descriptor
	SchemaV2Manifest SchemaV2Manifest
}

// Descriptor represents the descriptor section of a docker manifest
type Descriptor struct {
	MediaType string
	Digest    string
	Size      int
	Platform  Platform
}

// Platform represents the platform section of a docker manifest
type Platform struct {
	Architecture string
	Os           string
}

// SchemaV2Manifest represents the schema v2 manifest section of a docker manifest
type SchemaV2Manifest struct {
	SchemaVersion int
	MediaType     string
	Config        Config
	Layers        []Layer
}

// Config represents the config section of a docker manifest
type Config struct {
	MediaType string
	Size      int
	Digest    string
}

// Layer represents the layer section of a docker manifest
type Layer struct {
	MediaType string
	Size      int
	Digest    string
}
