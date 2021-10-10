/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package model

import (
	"path/filepath"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
)

// CGProject represents a Compose Generator project structure
type CGProject struct {
	CGProjectMetadata
	Composition       *spec.Project
	GitignorePatterns []string
	ReadmeChildPaths  []string
	ForceConfig       bool
	WithVolumesConfig bool
	Secrets           []ProjectSecret
	Vars              Vars
	ProxyVars         map[string]Vars
	Ports             []int
}

type Vars map[string]string

// CGProjectMetadata represents the metadata that is attached to a CGProject
type CGProjectMetadata struct {
	Name            string
	ContainerName   string
	WithGitignore   bool
	WithReadme      bool
	AdvancedConfig  bool
	ProductionReady bool
	CreatedBy       string
	CreatedAt       int64
	LastModifiedBy  string
	LastModifiedAt  int64
}

// GetAllVolumePaths returns the paths to all volumes, known by the project
func (p CGProject) GetAllVolumePaths() []string {
	paths := []string{}
	// Return empty list when no composition is attached
	if p.Composition == nil {
		return paths
	}
	// Search for volume paths in all services
	for _, service := range p.Composition.Services {
		for _, volume := range service.Volumes {
			if volume.Type == spec.VolumeTypeBind {
				paths = append(paths, volume.Source)
			}
		}
	}
	return paths
}

// GetAllBuildContextPaths returns the paths to all build contexts, known by the project
func (p CGProject) GetAllBuildContextPaths() []string {
	paths := []string{}
	// Return empty list when no composition is attached
	if p.Composition == nil {
		return paths
	}
	// Search for volume paths in all services
	for _, service := range p.Composition.Services {
		if service.Build != nil {
			path := service.Build.Context
			if strings.HasSuffix(path, "Dockerfile") {
				paths = append(paths, filepath.Dir(path))
			} else {
				paths = append(paths, path)
			}
		}
	}
	return paths
}

// GetAllEnvFilePaths returns all env file paths for the project
func (p CGProject) GetAllEnvFilePaths() []string {
	paths := []string{}
	// Return empty list when no composition is attached
	if p.Composition == nil {
		return paths
	}
	// Search for volume paths in all services
	for _, service := range p.Composition.Services {
		for _, envFile := range service.EnvFile {
			paths = append(paths, envFile)
		}
	}
	return paths
}

// GetAllEnvFilePathsNormalized returns all env file paths for the project without nested and duplicate paths
func (p CGProject) GetAllEnvFilePathsNormalized() []string {
	paths := p.GetAllEnvFilePaths()
	normalizedPaths := []string{}
	for _, path := range paths {
		// Check for duplicate
		duplicate := false
		for _, normalizedPath := range normalizedPaths {
			if path == normalizedPath {
				duplicate = true
				break
			}
		}
		if duplicate {
			continue
		}
		// No duplicate => add to list
		normalizedPaths = append(normalizedPaths, path)
	}
	return normalizedPaths
}

// ProjectSecret represents a secret in a CGProject
type ProjectSecret struct {
	Name     string
	Variable string
	Value    string
}
