package model

import "github.com/compose-spec/compose-go/types"

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
