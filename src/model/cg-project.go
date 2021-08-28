package model

import (
	"strings"

	"github.com/compose-spec/compose-go/types"
)

type CGProject struct {
	CGProjectMetadata
	Composition       *types.Project
	GitignorePatterns []string
	ReadmeChildPaths  []string
	ForceConfig       bool
	WithVolumesConfig bool
}

type CGProjectMetadata struct {
	Name           string
	ContainerName  string
	WithGitignore  bool
	WithReadme     bool
	AdvancedConfig bool
	CreatedBy      string
	CreatedAt      int64
	LastModifiedBy string
	LastModifiedAt int64
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
			if volume.Type == types.VolumeTypeBind {
				paths = append(paths, volume.Source)
			}
		}
	}
	return paths
}

// GetAllVolumePathsNormalized returns the paths to all volumes, whithout any duplicates and nested paths
func (p CGProject) GetAllVolumePathsNormalized() []string {
	paths := p.GetAllVolumePaths()
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
		// Check if nested in other paths
		containedInOtherPath := false
		for _, otherPath := range paths {
			// Skip the current path
			if path == otherPath {
				continue
			}
			// Check if the current path is nested in another path
			if strings.HasPrefix(path, otherPath) {
				containedInOtherPath = true
				break
			}
		}
		// Add to normalized list if not contained anywhere
		if !containedInOtherPath {
			normalizedPaths = append(normalizedPaths, path)
		}
	}
	return normalizedPaths
}
