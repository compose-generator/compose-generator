package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/briandowns/spinner"
	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// -------------------------------------------------------- GenerateResolveDependencyGroups --------------------------------------------------------

func TestGenerateResolveDependencyGroups(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "frontend-live-poll",
					DependsOn: spec.DependsOnConfig{
						model.TemplateTypeBackend: spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "backend-live-poll-api",
					DependsOn: spec.DependsOnConfig{
						model.TemplateTypeTlsHelper: spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "proxy-nginx",
					DependsOn: spec.DependsOnConfig{
						model.TemplateTypeDbAdmin: spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "database-redis",
					DependsOn: spec.DependsOnConfig{
						model.TemplateTypeProxy: spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "db-admin-phpmyadmin",
					DependsOn: spec.DependsOnConfig{
						model.TemplateTypeFrontend: spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "tls-helper-letsencrypt",
					DependsOn: spec.DependsOnConfig{
						model.TemplateTypeDatabase: spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
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
		DatabaseServices: []model.PredefinedTemplateConfig{
			{
				Name:  "redis",
				Label: "Redis",
			},
		},
		DbAdminServices: []model.PredefinedTemplateConfig{
			{
				Name:  "phpmyadmin",
				Label: "PhpMyAdmin",
			},
		},
		ProxyService: []model.PredefinedTemplateConfig{
			{
				Name:  "nginx",
				Label: "JWilder Nginx",
			},
		},
		TlsHelperService: []model.PredefinedTemplateConfig{
			{
				Name:  "letsencrypt",
				Label: "Let's Encrypt",
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
					DependsOn: spec.DependsOnConfig{
						"tls-helper-letsencrypt": spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "proxy-nginx",
					DependsOn: spec.DependsOnConfig{
						"db-admin-phpmyadmin": spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "database-redis",
					DependsOn: spec.DependsOnConfig{
						"proxy-nginx": spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "db-admin-phpmyadmin",
					DependsOn: spec.DependsOnConfig{
						"frontend-live-poll": spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
				{
					Name: "tls-helper-letsencrypt",
					DependsOn: spec.DependsOnConfig{
						"database-redis": spec.ServiceDependency{
							Condition: spec.ServiceConditionStarted,
						},
					},
				},
			},
		},
	}
	// Mock functions
	startProcessCallCount := 0
	startProcess = func(text string) (s *spinner.Spinner) {
		startProcessCallCount++
		assert.Equal(t, "Resolving group dependencies ...", text)
		return nil
	}
	stopProcessCallCount := 0
	stopProcess = func(s *spinner.Spinner) {
		stopProcessCallCount++
		assert.Nil(t, s)
	}
	// Execute test
	GenerateResolveDependencyGroups(project, selected)
	// Assert
	assert.Equal(t, 1, startProcessCallCount)
	assert.Equal(t, 1, stopProcessCallCount)
	assert.Equal(t, expectedProject, project)
}

// ------------------------------------------------------------ ReplaceGroupDependency -------------------------------------------------------------

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
