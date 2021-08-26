package model

import "github.com/compose-spec/compose-go/types"

type CGProject struct {
	Name             string
	ContainerName    string
	Project          *types.Project
	WithGitignore    bool
	GitignorePaths   []string
	WithReadme       bool
	ReadmeChildPaths []string
}
