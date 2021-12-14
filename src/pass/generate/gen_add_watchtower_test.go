package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// -------------------------------------------------------------- GenerateAddWatchtower ------------------------------------------------------------

func TestGenerateAddWatchTower1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &types.Project{
			Services: types.Services{
				{
					Name: "frontend-angular",
				},
				{
					Name: "frontend-vue",
				},
				{
					Name: "dbadmin-phpmyadmin",
				},
			},
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label: "Angular",
				Name:  "angular",
				Type:  "frontend",
			},
			{
				Label: "Vue",
				Name:  "vue",
				Type:  "frontend",
			},
		},
		DbAdminServices: []model.PredefinedTemplateConfig{
			{
				Label: "PhpMyAdmin",
				Name:  "phpmyadmin",
				Type:  "dbadmin",
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &types.Project{
			Services: types.Services{
				{
					Name: "frontend-angular",
					Labels: types.Labels{
						"com.centurylinklabs.watchtower.enable": "true",
					},
				},
				{
					Name: "frontend-vue",
				},
				{
					Name: "dbadmin-phpmyadmin",
					Labels: types.Labels{
						"com.centurylinklabs.watchtower.enable": "true",
					},
				},
				{
					Name: "companion-watchtower",
					Volumes: []types.ServiceVolumeConfig{
						{
							Type:   types.VolumeTypeBind,
							Source: "/var/run/docker.sock",
							Target: "/var/run/docker.sock",
						},
					},
					Command: types.ShellCommand{"--interval", "30"},
				},
			},
		},
	}
	// Mock functions
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to add Watchtower to check for new image versions?", question)
		assert.False(t, defaultValue)
		return true
	}
	multiSelectMenuQuestionIndex = func(label string, items, defaultItems []string) []int {
		assert.Equal(t, "For which services do you want to add Watchtower?", label)
		assert.Equal(t, []string{"Angular", "Vue", "PhpMyAdmin"}, items)
		return []int{0, 2}
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
	multiSelectMenuQuestionIndex = func(label string, items, defaultItems []string) []int {
		assert.Fail(t, "Unexpected call of multiSelectMenuQuestionIndex")
		return []int{}
	}
	// Execute test
	GenerateAddWatchtower(project, selectedTemplates)
	// Assert
	assert.Equal(t, expectedProject, project)
}
