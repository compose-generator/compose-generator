package project

import (
	"compose-generator/model"
	"compose-generator/util"
	"io/ioutil"
	"os/user"
	"path"
	"strings"
	"time"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
	"github.com/spf13/viper"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadProject loads the Docker compose project from the current directory
func LoadProject(options ...LoadOption) *model.CGProject {
	opts := applyLoadOptions(options...)

	// Create project instance
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: util.FileExists(opts.WorkingDir + ".gitignore"),
			WithReadme:    util.FileExists(opts.WorkingDir + "README.md"),
		},
		GitignorePatterns: []string{},
		ReadmeChildPaths:  []string{"README.md"},
		ForceConfig:       false,
	}

	// Load components
	loadComposeFile(project, opts)
	loadGitignoreFile(project, opts)
	loadCGFile(&project.CGProjectMetadata, opts)

	return project
}

func LoadProjectMetadata(options ...LoadOption) *model.CGProjectMetadata {
	opts := applyLoadOptions(options...)

	// Create project metadata instance
	metadata := &model.CGProjectMetadata{
		WithGitignore: util.FileExists(opts.WorkingDir + ".gitignore"),
		WithReadme:    util.FileExists(opts.WorkingDir + "README.md"),
	}

	// Load metadata
	loadCGFile(metadata, opts)

	return metadata
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func loadComposeFile(project *model.CGProject, opt LoadOptions) {
	if !util.FileExists(opt.WorkingDir + opt.ComposeFileName) {
		util.Error("Compose file not found", nil, true)
	}
	content, err := ioutil.ReadFile(opt.WorkingDir + opt.ComposeFileName)
	if err != nil {
		util.Error("Unable to parse '"+opt.ComposeFileName+"'", err, true)
	}
	dict, err := loader.ParseYAML(content)
	if err != nil {
		util.Error("Unable to parse '"+opt.ComposeFileName+"' file", err, true)
	}

	configs := []types.ConfigFile{
		{
			Filename: opt.ComposeFileName,
			Config:   dict,
		},
	}
	config := types.ConfigDetails{
		WorkingDir:  opt.WorkingDir,
		ConfigFiles: configs,
	}
	project.Composition, err = loader.Load(config)
	if err != nil {
		util.Error("Could not load project from the current directory", err, true)
	}
}

func loadGitignoreFile(project *model.CGProject, opt LoadOptions) {
	if project.WithGitignore {
		// Load patterns from .gitignore file
		content, err := ioutil.ReadFile(opt.WorkingDir + ".gitignore")
		if err != nil {
			util.Error("Unable to parse .gitignore file", err, true)
		}
		contentStr := strings.ReplaceAll(string(content), "\r\n", "\n")
		// Save them into the project
		for _, line := range strings.Split(contentStr, "\n") {
			trimmedLine := strings.TrimSpace(line)
			if len(trimmedLine) > 0 && !strings.HasPrefix(trimmedLine, "#") {
				project.GitignorePatterns = append(project.GitignorePatterns, trimmedLine)
			}
		}
	}
}

func loadCGFile(metadata *model.CGProjectMetadata, opt LoadOptions) {
	// Get default config values
	defaultProjectName := path.Base(opt.WorkingDir)
	defaultContainerName := strings.ReplaceAll(strings.ToLower(defaultProjectName), " ", "-")
	defaultAdvancedMode := false
	defaultProductionReady := false
	defaultCreatedBy := "unknown"
	defaultModifiedBy := "unknown"
	if user, err := user.Current(); err == nil {
		defaultCreatedBy = user.Name
		defaultModifiedBy = user.Name
	}
	defaultCreatedAt := time.Now().UnixNano()
	defaultModifiedAt := time.Now().UnixNano()

	configFileName := ".cg.yml"
	if util.FileExists(opt.WorkingDir + configFileName) {
		config := viper.New()
		// Set default values
		config.SetDefault("project-name", defaultProjectName)
		config.SetDefault("project-container-name", defaultContainerName)
		config.SetDefault("advanced-config", defaultAdvancedMode)
		config.SetDefault("production-ready", defaultProductionReady)
		config.SetDefault("created-by", defaultCreatedBy)
		config.SetDefault("created-at", defaultCreatedAt)
		config.SetDefault("modified-by", defaultModifiedBy)
		config.SetDefault("modified-at", defaultModifiedAt)
		// Load config file
		config.SetConfigName(".cg")
		config.SetConfigType("yml")
		config.AddConfigPath(opt.WorkingDir)
		err := config.ReadInConfig()
		if err != nil {
			util.Error("Could not read '"+configFileName+"' file", err, true)
		}
		// Assign values
		metadata.Name = config.GetString("project-name")
		metadata.ContainerName = config.GetString("project-container-name")
		metadata.AdvancedConfig = config.GetBool("advanced-config")
		metadata.ProductionReady = config.GetBool("production-ready")
		metadata.CreatedBy = config.GetString("created-by")
		metadata.CreatedAt = config.GetInt64("created-at")
		metadata.LastModifiedBy = config.GetString("modified-by")
		metadata.LastModifiedAt = config.GetInt64("modified-at")
	} else {
		metadata.Name = defaultProjectName
		metadata.ContainerName = defaultContainerName
		metadata.AdvancedConfig = defaultAdvancedMode
		metadata.CreatedBy = defaultCreatedBy
		metadata.CreatedAt = defaultCreatedAt
		metadata.LastModifiedBy = defaultModifiedBy
		metadata.LastModifiedAt = defaultModifiedAt
	}
}
