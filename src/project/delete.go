package project

import "compose-generator/model"

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// DeleteProject deletes a project from the disk
func DeleteProject(project *model.CGProject, options ...DeleteOption) {
	opts := applyDeleteOptions(options...)

	deleteReadme(project, opts)
	deleteEnvFile(project, opts)
	deleteGitignore(project, opts)
	deleteComposeFile(project, opts)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func deleteReadme(project *model.CGProject, options DeleteOptions) {

}

func deleteEnvFile(project *model.CGProject, options DeleteOptions) {

}

func deleteGitignore(project *model.CGProject, options DeleteOptions) {

}

func deleteComposeFile(project *model.CGProject, options DeleteOptions) {

}
