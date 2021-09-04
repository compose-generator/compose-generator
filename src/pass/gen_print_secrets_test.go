package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrintSecrets(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Secrets: []model.ProjectSecret{
			{
				Name:     "Secret No. 1",
				Variable: "SECRET_1",
				Value:    "HTRqX9Cb72LHSM4LahwVTtWQktFwx6",
			},
			{
				Name:     "Secret No. 2",
				Variable: "SECRET_2",
				Value:    "tkzN4rfQMDWgpLWcQp5sWLWgHVXgWG9maFgUaG9x3u7t5sg3z2",
			},
		},
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	pCallCount := 0
	p = func(text string) {
		pCallCount++
		if pCallCount == 1 {
			assert.Equal(t, "🔑   Secret No. 1: ", text)
		} else {
			assert.Equal(t, "🔑   Secret No. 2: ", text)
		}
	}
	pl = func(text string) {
		assert.Equal(t, "Following secrets were automatically generated:", text)
	}
	printSecretValue = func(format string, a ...interface{}) {
		if pCallCount == 1 {
			assert.Equal(t, "HTRqX9Cb72LHSM4LahwVTtWQktFwx6", format)
		} else {
			assert.Equal(t, "tkzN4rfQMDWgpLWcQp5sWLWgHVXgWG9maFgUaG9x3u7t5sg3z2", format)
		}
	}
	// Execute test
	GeneratePrintSecrets(project)
	// Assert
	assert.Equal(t, 1, pelCallCount)
}
