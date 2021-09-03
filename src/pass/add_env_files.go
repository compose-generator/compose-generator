package pass

import (
	"compose-generator/model"
	"path/filepath"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddEnvFiles asks the user if he/she wants to add env files to the configuration
func AddEnvFiles(service *spec.ServiceConfig, _ *model.CGProject) {
	if YesNoQuestion("Do you want to provide environment files to your service?", false) {
		Pel()
		for another := true; another; another = YesNoQuestion("Add another environment file?", true) {
			// Ask for env file with auto-suggested test input
			envFile := TextQuestionWithDefaultAndSuggestions("Where is your env file located?", "environment.env", func(toComplete string) (files []string) {
				files, _ = filepath.Glob(toComplete + "*.*")
				return
			})
			// Check if the selected file is valid
			if !FileExists(envFile) || IsDir(envFile) {
				Error("File is not valid. Please select another file", nil, false)
				continue
			}
			// Add env file to service
			service.EnvFile = append(service.EnvFile, envFile)
		}
		Pel()
	}
}
