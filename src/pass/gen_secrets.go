package pass

import (
	"compose-generator/model"
	"compose-generator/util"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateSecrets generates all secrets for a stack
func GenerateSecrets(project *model.CGProject, config *model.GenerateConfig) {
	util.P("Generating secrets ... ")

	util.Done()
}
