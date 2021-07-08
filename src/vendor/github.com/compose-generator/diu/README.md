# Docker Inspect Utils
![GitHub release](https://img.shields.io/github/v/release/compose-generator/diu?include_prereleases)
![Go CI](https://github.com/compose-generator/diu/workflows/Go%20CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/compose-generator/diu)](https://goreportcard.com/report/github.com/compose-generator/diu)
[![Codecov](https://codecov.io/gh/compose-generator/diu/branch/main/graph/badge.svg?token=0EoAPqmDCv)](https://codecov.io/gh/compose-generator/diu)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

This Go library contains a set of useful features to parse internal Docker objects and put them into handelable objects or slices.

*Note: This library is part of the [Compose Generator](https://github.com/compose-generator/compose-generator) project, but also can be used independently.*

## Installation
```sh
go get github.com/compose-generator/diu
```

## Usage
### Get manifest of remote image
Returns a struct with following structure: [Structure](model/manifest.go)

**Example:**
```go
// You also can pass the image with a custom repository e.g.: ghcr.io/compose-generator/compose-generator
manifest, err := GetImageManifest("hello-world")
if err == nil {
    // Print layer count of hello-world image
    fmt.println("Number of layers: "+len(manifest.SchemaV2Manifest.Layers))
}
```

### Get all volumes of local Docker instance
Returns a slice of structs with following structure: [Structure](model/volume.go)

**Example:**
```go
volumes, err := GetExistingVolumes()
if err == nil && len(volumes) > 0 {
    // Print layer count of hello-world image
    fmt.println("Name of first volume: "+volumes[0].Name)
}
```

### Get all networks of local Docker instance
Returns a slice of structs with following structure: [Structure](model/network.go)

**Example:**
```go
networks, err := GetExistingNetworks()
if err == nil && len(networks) > 0 {
    // Print layer count of hello-world image
    fmt.println("Name of first networks: "+networks[0].Name)
}
```

*TODO: To be extended*

## Contribute to the project
If you want to contribute to this project, please ensure you comply with the [contribution guidelines](CONTRIBUTING.md).

© Marc Auberer 2021