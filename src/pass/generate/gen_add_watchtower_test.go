/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

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
			Services: types.Services{},
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label:       "Angular",
				Name:        "angular",
				Type:        "frontend",
				AutoUpdated: false,
			},
			{
				Label:       "Vue",
				Name:        "vue",
				Type:        "frontend",
				AutoUpdated: false,
			},
		},
		DbAdminServices: []model.PredefinedTemplateConfig{
			{
				Label:       "PhpMyAdmin",
				Name:        "phpmyadmin",
				Type:        "dbadmin",
				AutoUpdated: true,
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &types.Project{
			Services: types.Services{
				{
					Name:    "companion-watchtower",
					Image:   "containrrr/watchtower:latest",
					Restart: types.RestartPolicyUnlessStopped,
					Volumes: []types.ServiceVolumeConfig{
						{
							Type:   types.VolumeTypeBind,
							Source: "/var/run/docker.sock",
							Target: "/var/run/docker.sock",
						},
					},
					DependsOn: types.DependsOnConfig{
						model.TemplateTypeFrontend:  types.ServiceDependency{},
						model.TemplateTypeBackend:   types.ServiceDependency{},
						model.TemplateTypeDatabase:  types.ServiceDependency{},
						model.TemplateTypeDbAdmin:   types.ServiceDependency{},
						model.TemplateTypeProxy:     types.ServiceDependency{},
						model.TemplateTypeTlsHelper: types.ServiceDependency{},
					},
					Command: types.ShellCommand{"--interval", "30"},
				},
			},
		},
	}
	expectedSelectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label:       "Angular",
				Name:        "angular",
				Type:        "frontend",
				AutoUpdated: true,
			},
			{
				Label:       "Vue",
				Name:        "vue",
				Type:        "frontend",
				AutoUpdated: false,
			},
		},
		DbAdminServices: []model.PredefinedTemplateConfig{
			{
				Label:       "PhpMyAdmin",
				Name:        "phpmyadmin",
				Type:        "dbadmin",
				AutoUpdated: true,
			},
		},
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
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
	GenerateAddWatchtower(project, selectedTemplates, nil)
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, expectedProject, project)
	assert.Equal(t, expectedSelectedTemplates, selectedTemplates)
}

func TestGenerateAddWatchTower2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label:       "Angular",
				Name:        "angular",
				Type:        "frontend",
				AutoUpdated: false,
			},
			{
				Label:       "Vue",
				Name:        "vue",
				Type:        "frontend",
				AutoUpdated: false,
			},
		},
		DbAdminServices: []model.PredefinedTemplateConfig{
			{
				Label:       "PhpMyAdmin",
				Name:        "phpmyadmin",
				Type:        "dbadmin",
				AutoUpdated: true,
			},
		},
	}
	expectedProject := &model.CGProject{}
	expectedSelectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label:       "Angular",
				Name:        "angular",
				Type:        "frontend",
				AutoUpdated: false,
			},
			{
				Label:       "Vue",
				Name:        "vue",
				Type:        "frontend",
				AutoUpdated: false,
			},
		},
		DbAdminServices: []model.PredefinedTemplateConfig{
			{
				Label:       "PhpMyAdmin",
				Name:        "phpmyadmin",
				Type:        "dbadmin",
				AutoUpdated: false,
			},
		},
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
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
	GenerateAddWatchtower(project, selectedTemplates, &model.GenerateConfig{FromFile: true})
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, expectedProject, project)
	assert.Equal(t, expectedSelectedTemplates, selectedTemplates)
}

func TestGenerateAddWatchTower3(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	selectedTemplates := &model.SelectedTemplates{}
	expectedProject := &model.CGProject{}
	expectedSelectedTemplates := &model.SelectedTemplates{}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
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
	GenerateAddWatchtower(project, selectedTemplates, nil)
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, expectedProject, project)
	assert.Equal(t, expectedSelectedTemplates, selectedTemplates)
}
