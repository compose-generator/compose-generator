package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// -------------------------------------------------------------- GenerateAddWatchtower ------------------------------------------------------------

func TestGenerateAddWatchTower1(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{},
	}
	expectedProject := &model.CGProject{}
	// Mock functions
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to add Watchtower to check for new image versions?", question)
		assert.False(t, defaultValue)
		return true
	}
	multiSelectMenuQuestion = func(label string, items []string) []string {
		assert.Equal(t, "For which services do you want to add Watchtower?", label)
		assert.Equal(t, []string{}, items)
		return []string{"Angular", "Vue", "Matomo", "Ghost"}
	}
	// Execute test
	GenerateAddWatchtower(project, selectedTemplates)
	// Assert
	assert.Equal(t, expectedProject, project)
}

func TestGenerateAddWatchTower2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	selectedTemplates := &model.SelectedTemplates{}
	expectedProject := &model.CGProject{}
	// Mock functions
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to add Watchtower to check for new image versions?", question)
		assert.False(t, defaultValue)
		return false
	}
	multiSelectMenuQuestion = func(label string, items []string) []string {
		assert.Fail(t, "Unexpected call of multiSelectMenuQuestionIndex")
		return []string{}
	}
	// Execute test
	GenerateAddWatchtower(project, selectedTemplates)
	// Assert
	assert.Equal(t, expectedProject, project)
}
