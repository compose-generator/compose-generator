package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSecrets(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	selected := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Name: "MySQL",
				Secrets: []model.Secret{
					{
						Name:     "MySQL root password",
						Variable: "MYSQL_ROOT_PASSWORD",
						Length:   50,
					},
					{
						Name:     "MySQL user password",
						Variable: "MYSQL_USER_PASSWORD",
						Length:   30,
					},
				},
			},
		},
	}
	expectedProject := &model.CGProject{
		Secrets: []model.ProjectSecret{
			{
				Name:     "MySQL root password",
				Variable: "MYSQL_ROOT_PASSWORD",
				Value:    "tkzN4rfQMDWgpLWcQp5sWLWgHVXgWG9maFgUaG9x3u7t5sg3z2",
			},
			{
				Name:     "MySQL user password",
				Variable: "MYSQL_USER_PASSWORD",
				Value:    "HTRqX9Cb72LHSM4LahwVTtWQktFwx6",
			},
		},
	}
	// Mock functions
	p = func(text string) {
		assert.Equal(t, "Generating secrets ... ", text)
	}
	doneCallCount := 0
	done = func() {
		doneCallCount++
	}
	generatePasswordCallCount := 0
	generatePassword = func(length, numDigits, numSymbols int, noUpper, allowRepeat bool) (string, error) {
		generatePasswordCallCount++
		if generatePasswordCallCount == 1 {
			assert.Equal(t, 50, length)
			return "tkzN4rfQMDWgpLWcQp5sWLWgHVXgWG9maFgUaG9x3u7t5sg3z2", nil
		} else {
			assert.Equal(t, 30, length)
			return "HTRqX9Cb72LHSM4LahwVTtWQktFwx6", nil
		}
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	GenerateSecrets(project, selected)
	// Assert
	assert.Equal(t, 1, doneCallCount)
	assert.Equal(t, expectedProject, project)
}
