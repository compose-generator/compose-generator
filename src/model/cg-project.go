package model

import "github.com/compose-spec/compose-go/types"

type CGProject struct {
	Project          *types.Project
	WithGitignore    bool
	GitignorePaths   []string
	WithReadme       bool
	ReadmeChildPaths []string
}
