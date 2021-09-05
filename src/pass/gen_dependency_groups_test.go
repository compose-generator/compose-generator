package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateResolveDependencyGroups(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "frontend-live-poll",
					DependsOn: spec.DependsOnConfig{
						"backend": spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "backend-live-poll-api",
				},
			},
		},
	}
	selected := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Name:  "live-poll",
				Label: "Live-Poll",
			},
		},
		BackendServices: []model.PredefinedTemplateConfig{
			{
				Name:  "live-poll-api",
				Label: "Live-Poll API",
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "frontend-live-poll",
					DependsOn: spec.DependsOnConfig{
						"backend-live-poll-api": spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "backend-live-poll-api",
				},
			},
		},
	}
	// Mock functions
	pCallCount := 0
	p = func(text string) {
		pCallCount++
		assert.Equal(t, "Resolving group dependencies ... ", text)
	}
	doneCallCount := 0
	done = func() {
		doneCallCount++
	}
	// Execute test
	GenerateResolveDependencyGroups(project, selected)
	// Assert
	assert.Equal(t, 1, pCallCount)
	assert.Equal(t, 1, doneCallCount)
	assert.Equal(t, expectedProject, project)
}

func TestReplaceGroupDependency(t *testing.T) {
	// Test data
	groupName := model.TemplateTypeFrontend
	service := &spec.ServiceConfig{
		DependsOn: spec.DependsOnConfig{
			"frontend": spec.ServiceDependency{
				Condition: spec.ServiceConditionStarted,
			},
		},
	}
	templates := []model.PredefinedTemplateConfig{
		{
			Label: "Angular",
			Name:  "angular",
		},
		{
			Label: "Vue",
			Name:  "vue",
		},
	}
	expectedService := &spec.ServiceConfig{
		DependsOn: spec.DependsOnConfig{
			"frontend-angular": spec.ServiceDependency{
				Condition: spec.ServiceConditionStarted,
			},
			"frontend-vue": spec.ServiceDependency{
				Condition: spec.ServiceConditionStarted,
			},
		},
	}
	// Execute test
	replaceGroupDependency(service, templates, groupName)
	// Assert
	assert.Equal(t, expectedService, service)
}
