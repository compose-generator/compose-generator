package project

import (
	"compose-generator/model"
	"compose-generator/util"
	"io/ioutil"
	"os"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

type LoadOptions struct {
	ComposeFileName string
}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadProject loads the Docker compose project from the current directory
func LoadProject(options ...LoadOptions) *model.CGProject {
	// Load default options
	opt := LoadOptions{ComposeFileName: "docker-compose.yml"}
	// Replace with the passed options, if available
	if len(options) > 0 {
		opt = options[0]
	}

	// Create project instance
	project := &model.CGProject{
		WithGitignore: false,
		WithReadme:    false,
	}

	// Load components
	loadComposeFile(project, opt)

	return project
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func loadComposeFile(project *model.CGProject, opt LoadOptions) {
	content, err := ioutil.ReadFile(opt.ComposeFileName)
	if err != nil {
		util.Error("Unable to parse '"+opt.ComposeFileName+"'", err, true)
	}
	dict, err := loader.ParseYAML(content)
	if err != nil {
		util.Error("Unable to parse '"+opt.ComposeFileName+"' file", err, true)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		util.Error("Unable to retrieve workdir", err, true)
	}

	configs := []types.ConfigFile{
		{
			Filename: opt.ComposeFileName,
			Config:   dict,
		},
	}
	config := types.ConfigDetails{
		WorkingDir:  workingDir,
		ConfigFiles: configs,
		Environment: nil,
	}
	project.Project, err = loader.Load(config)
	if err != nil {
		util.Error("Could not load project from the current directory", err, true)
	}
}
