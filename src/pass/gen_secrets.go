package pass

import (
	"compose-generator/model"

	"github.com/sethvargo/go-password/password"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateSecrets generates all secrets for a stack
func GenerateSecrets(project *model.CGProject, selected *model.SelectedTemplates) {
	p("Generating secrets ... ")
	if project.Vars == nil {
		project.Vars = make(map[string]string)
	}
	for _, template := range selected.GetAll() {
		for _, secret := range template.Secrets {
			// Generate secret
			res, err := password.Generate(secret.Length, 10, 0, false, false)
			if err != nil {
				printError("Password generation failed.", err, true)
			}
			project.Secrets = append(project.Secrets, model.ProjectSecret{
				Name:     secret.Name,
				Variable: secret.Variable,
				Value:    res,
			})
		}
	}
	done()
}
