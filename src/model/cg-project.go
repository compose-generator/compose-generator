package model

import (
	"path/filepath"
	"strings"

	"github.com/compose-spec/compose-go/types"
)

type CGProject struct {
	Name              string
	ContainerName     string
	Composition       *types.Project
	WithGitignore     bool
	GitignorePatterns []string
	WithReadme        bool
	ReadmeChildPaths  []string
	AdvancedConfig    bool
	ForceConfig       bool
	WithVolumesConfig bool
}

func (p CGProject) getAllVolumePaths() []string {
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

func (p CGProject) getAllVolumePathsNormalized() []string {
	paths := p.getAllVolumePaths()
	normalizedPaths := []string{}
	for _, path := range paths {
		pathAbs, err := filepath.Abs(path)
		if err != nil {
			continue
		}
		// Check if the path is not contained in other paths
		containedInOtherPath := false
		for _, otherPath := range paths {
			otherPathAbs, err := filepath.Abs(otherPath)
			if err != nil {
				continue
			}
			if strings.HasPrefix(pathAbs, otherPathAbs) {
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
