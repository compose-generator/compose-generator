package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"path/filepath"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddEnvFiles asks the user if he/she wants to add env files to the configuration
func AddEnvFiles(service *spec.ServiceConfig, _ *model.CGProject) {
	if util.YesNoQuestion("Do you want to provide an environment file to your service?", false) {
		util.Pel()
		for another := true; another; another = util.YesNoQuestion("Add another environment file?", true) {
			// Ask for env file with auto-suggested test input
			envFile := util.TextQuestionWithDefaultAndSuggestions("Where is your env file located?", "environment.env", func(toComplete string) (files []string) {
				files, _ = filepath.Glob(toComplete + "*.*")
				return
			})
			// Check if the selected file is valid
			if !util.FileExists(envFile) || util.IsDir(envFile) {
				util.Error("File is not valid. Please select another file", nil, false)
				continue
			}
			// Add env file to service
			service.EnvFile = append(service.EnvFile, envFile)
		}
		util.Pel()
	}
}
