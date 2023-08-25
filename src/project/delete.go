/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package project

import (
	"compose-generator/model"
)

var deleteReadmeMockable = deleteReadme
var deleteEnvFilesMockable = deleteEnvFiles
var deleteGitignoreMockable = deleteGitignore
var deleteVolumesMockable = deleteVolumes
var deleteComposeFileMockable = deleteComposeFile

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// DeleteProject deletes a project from the disk
func DeleteProject(project *model.CGProject, options ...DeleteOption) {
	// Apply options
	opts := applyDeleteOptions(options...)

	// Delete all components
	deleteReadmeMockable(project, opts)
	deleteEnvFilesMockable(project, opts)
	deleteGitignoreMockable(project, opts)
	deleteVolumesMockable(project, opts)
	deleteComposeFileMockable(project, opts)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func deleteReadme(project *model.CGProject, opt DeleteOptions) {
	if project.WithReadme {
		infoLogger.Println("Deleting Readme ...")
		err := remove(opt.WorkingDir + "README.md")
		if err != nil {
			warningLogger.Println("File 'README.md' could not be deleted: " + err.Error())
			logWarning("File 'README.md' could not be deleted")
		}
		infoLogger.Println("Deleting Readme (done)")
	}
}

func deleteEnvFiles(project *model.CGProject, opt DeleteOptions) {
	for _, envFilePath := range project.GetAllEnvFilePathsNormalized() {
		infoLogger.Println("Deleting env file '" + envFilePath + "' ...")
		// Try to delete the env file
		err := remove(opt.WorkingDir + envFilePath)
		if err != nil {
			warningLogger.Println("File '" + opt.WorkingDir + envFilePath + "' could not be deleted: " + err.Error())
			logWarning("File '" + opt.WorkingDir + envFilePath + "' could not be deleted")
		}
		infoLogger.Println("Deleting env file '" + envFilePath + "' (done)")
	}
}

func deleteGitignore(project *model.CGProject, opt DeleteOptions) {
	if project.WithGitignore {
		infoLogger.Println("Deleting Gitignore ...")
		err := remove(opt.WorkingDir + ".gitignore")
		if err != nil {
			warningLogger.Println("File '.gitignore' could not be deleted: " + err.Error())
			logWarning("File '.gitignore' could not be deleted")
		}
		infoLogger.Println("Deleting Gitignore (done)")
	}
}

func deleteVolumes(project *model.CGProject, opt DeleteOptions) {
	for _, volumePath := range normalizePaths(project.GetAllVolumePaths()) {
		infoLogger.Println("Deleting volume '" + volumePath + "' ...")
		err := removeAll(volumePath)
		if err != nil {
			warningLogger.Println("Volume '" + volumePath + "' could not be deleted: " + err.Error())
			logWarning("Volume '" + volumePath + "' could not be deleted")
		}
		infoLogger.Println("Deleting volume '" + volumePath + "' (done)")
	}
}

func deleteComposeFile(project *model.CGProject, opt DeleteOptions) {
	infoLogger.Println("Deleting Compose file ...")
	err := remove(opt.WorkingDir + opt.ComposeFileName)
	if err != nil {
		warningLogger.Println("File '" + opt.ComposeFileName + "' could not be deleted: " + err.Error())
		logWarning("File '" + opt.ComposeFileName + "' could not be deleted")
	}
	infoLogger.Println("Deleting Compose file (done)")
}
