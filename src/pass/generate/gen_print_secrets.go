package pass

import (
	"compose-generator/model"
	"compose-generator/util"
)

// GeneratePrintSecrets prints all generated secrets to the console
func GeneratePrintSecrets(project *model.CGProject) {
	pel()
	pl("Following secrets were automatically generated:")
	for _, secret := range project.Secrets {
		secretName := util.ReplaceVarsInString(secret.Name, project.Vars)
		p("🔑   " + secretName + ": ")
		printSecretValue(secret.Value)
	}
}
