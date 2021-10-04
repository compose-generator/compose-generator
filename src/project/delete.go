package project

import (
	"compose-generator/model"
	"compose-generator/util"
	"os"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// DeleteProject deletes a project from the disk
func DeleteProject(project *model.CGProject, options ...DeleteOption) {
	opts := applyDeleteOptions(options...)

	deleteReadme(project, opts)
	deleteEnvFiles(project, opts)
	deleteGitignore(project, opts)
	deleteVolumes(project, opts)
	deleteComposeFile(project, opts)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func deleteReadme(project *model.CGProject, opt DeleteOptions) {
	if project.WithReadme {
		err := os.Remove(opt.WorkingDir + "README.md")
		if err != nil {
			util.Warning("File 'README.md' could not be deleted")
		}
	}
}

func deleteEnvFiles(project *model.CGProject, opt DeleteOptions) {
	for _, envFilePath := range project.GetAllEnvFilePathsNormalized() {
		// Try to delete the env file
		err := os.Remove(opt.WorkingDir + envFilePath)
		if err != nil {
			util.Warning("File '" + envFilePath + "' could not be deleted")
		}
	}
}

func deleteGitignore(project *model.CGProject, opt DeleteOptions) {
	if project.WithGitignore {
		err := os.Remove(opt.WorkingDir + ".gitignore")
		if err != nil {
			util.Warning("File '.gitignore' could not be deleted")
		}
	}
}

func deleteVolumes(project *model.CGProject, opt DeleteOptions) {
	for _, volumePath := range util.NormalizePaths(project.GetAllVolumePaths()) {
		err := os.RemoveAll(volumePath)
		if err != nil {
			util.Warning("Volume '" + volumePath + "' could not be deleted")
		}
	}
}

func deleteComposeFile(project *model.CGProject, opt DeleteOptions) {
	err := os.Remove(opt.WorkingDir + "docker-compose.yml")
	if err != nil {
		util.Warning("File 'docker-compose.yml' could not be deleted")
	}
}
