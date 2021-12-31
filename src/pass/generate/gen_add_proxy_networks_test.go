/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------ GenerateAddProxyNetworks -----------------------------------------------------------

func TestGenerateAddProxyNetworks1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: true,
		},
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "frontend-angular",
				},
				{
					Name: "tlshelper-letsencrypt",
				},
				{
					Name: "backend-spring-maven",
				},
				{
					Name: "proxy-nginx",
				},
			},
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Name: "angular",
				Type: model.TemplateTypeFrontend,
			},
		},
		BackendServices: []model.PredefinedTemplateConfig{
			{
				Name:    "spring-maven",
				Proxied: true,
				Type:    model.TemplateTypeBackend,
			},
			{
				Name:    "spring-gradle",
				Proxied: true,
				Type:    model.TemplateTypeBackend,
			},
		},
		ProxyServices: []model.PredefinedTemplateConfig{
			{
				Name: "nginx",
				Type: model.TemplateTypeProxy,
			},
		},
		TlsHelperServices: []model.PredefinedTemplateConfig{
			{
				Name: "letsencrypt",
				Type: model.TemplateTypeTlsHelper,
			},
		},
	}
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: true,
		},
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "frontend-angular",
				},
				{
					Name: "tlshelper-letsencrypt",
				},
				{
					Name: "backend-spring-maven",
					Networks: map[string]*spec.ServiceNetworkConfig{
						"proxy-spring-maven": nil,
					},
				},
				{
					Name: "proxy-nginx",
					Networks: map[string]*spec.ServiceNetworkConfig{
						"proxy-spring-maven": nil,
					},
				},
			},
		},
	}
	// Execute test
	GenerateAddProxyNetworks(project, selectedTemplates)
	// Assert
	assert.Equal(t, expectedProject, project)
}

func TestGenerateAddProxyNetworks2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: false,
		},
	}
	selectedTemplates := &model.SelectedTemplates{}
	// Execute test
	GenerateAddProxyNetworks(project, selectedTemplates)
}

func TestGenerateAddProxyNetworks3(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: true,
		},
	}
	selectedTemplates := &model.SelectedTemplates{}
	// Execute test
	GenerateAddProxyNetworks(project, selectedTemplates)
}

func TestGenerateAddProxyNetworks4(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: true,
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		ProxyServices: []model.PredefinedTemplateConfig{
			{
				Name: "nginx",
				Type: model.TemplateTypeProxy,
			},
		},
	}
	// Mock functions
	logErrorCallCount := 0
	logError = func(message string, exit bool) {
		logErrorCallCount++
		assert.Equal(t, "Proxy service cannot be found for network inserting", message)
		assert.True(t, exit)
	}
	// Execute test
	GenerateAddProxyNetworks(project, selectedTemplates)
	// Assert
	assert.Equal(t, 1, logErrorCallCount)
}
