package project

import (
	"compose-generator/model"
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
		err := remove(opt.WorkingDir + "README.md")
		if err != nil {
			printWarning("File 'README.md' could not be deleted")
		}
	}
}

func deleteEnvFiles(project *model.CGProject, opt DeleteOptions) {
	for _, envFilePath := range project.GetAllEnvFilePathsNormalized() {
		// Try to delete the env file
		err := remove(opt.WorkingDir + envFilePath)
		if err != nil {
			printWarning("File '" + envFilePath + "' could not be deleted")
		}
	}
}

func deleteGitignore(project *model.CGProject, opt DeleteOptions) {
	if project.WithGitignore {
		err := remove(opt.WorkingDir + ".gitignore")
		if err != nil {
			printWarning("File '.gitignore' could not be deleted")
		}
	}
}

func deleteVolumes(project *model.CGProject, opt DeleteOptions) {
	for _, volumePath := range normalizePaths(project.GetAllVolumePaths()) {
		err := removeAll(volumePath)
		if err != nil {
			printWarning("Volume '" + volumePath + "' could not be deleted")
		}
	}
}

func deleteComposeFile(project *model.CGProject, opt DeleteOptions) {
	err := remove(opt.WorkingDir + "docker-compose.yml")
	if err != nil {
		printWarning("File 'docker-compose.yml' could not be deleted")
	}
}
