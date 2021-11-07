/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateSecrets generates all secrets for a stack
func GenerateSecrets(project *model.CGProject, selected *model.SelectedTemplates) {
	infoLogger.Println("Generating secrets ...")
	spinner := startProcess("Generating secrets ...")
	for _, template := range selected.GetAll() {
		for _, secret := range template.Secrets {
			// Only generate the secrets that are not customizable, the customizable ones are already generated
			if !secret.Customizable {
				// Generate secret
				res, err := generatePassword(secret.Length, 10, 0, false, false)
				if err != nil {
					errorLogger.Println("Password generation failed: " + err.Error())
					logError("Password generation failed", true)
				}
				project.Secrets = append(project.Secrets, model.ProjectSecret{
					Name:     secret.Name,
					Variable: secret.Variable,
					Value:    res,
				})
			}
		}
	}
	stopProcess(spinner)
	infoLogger.Println("Generating secrets (done)")
}
