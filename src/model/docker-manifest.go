package model

// Json structure of docker compose file
type DockerManifest struct {
	Ref              string
	Descriptor       Service
	SchemaV2Manifest SchemaV2Manifest
}

type Descriptor struct {
	MediaType string
	Digest    string
	Size      int
	Platform  Platform
}

type Platform struct {
	Architecture string
	Os           string
}

type SchemaV2Manifest struct {
	SchemaVersion int
	MediaType     string
	Config        Config
	Layers        []Layer
}

type Config struct {
	MediaType string
	Size      int
	Digest    string
}

type Layer struct {
	MediaType string
	Size      int
	Digest    string
}
