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
		ProxyService: []model.PredefinedTemplateConfig{
			{
				Name: "nginx",
				Type: model.TemplateTypeProxy,
			},
		},
		TlsHelperService: []model.PredefinedTemplateConfig{
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
	// Mock functions
	getServiceRefCallCount := 0
	getServiceRefMockable = func(proj *spec.Project, serviceName string) *spec.ServiceConfig {
		getServiceRefCallCount++
		assert.Equal(t, project.Composition, proj)
		switch getServiceRefCallCount {
		case 1:
			assert.Equal(t, "proxy-nginx", serviceName)
			return &project.Composition.Services[3]
		case 2:
			assert.Equal(t, "backend-spring-maven", serviceName)
			return &project.Composition.Services[2]
		case 3:
			assert.Equal(t, "backend-spring-gradle", serviceName)
		}
		return nil
	}
	// Execute test
	GenerateAddProxyNetworks(project, selectedTemplates)
	// Assert
	assert.Equal(t, expectedProject, project)
	assert.Equal(t, 3, getServiceRefCallCount)
}

func TestGenerateAddProxyNetworks2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: false,
		},
	}
	selectedTemplates := &model.SelectedTemplates{}
	// Mock functions
	getServiceRefMockable = func(project *spec.Project, serviceName string) *spec.ServiceConfig {
		assert.Fail(t, "Unexpected call of getServiceRef")
		return nil
	}
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
	selectedTemplates := &model.SelectedTemplates{
		ProxyService: []model.PredefinedTemplateConfig{
			{
				Name: "nginx",
				Type: model.TemplateTypeProxy,
			},
		},
	}
	// Mock functions
	getServiceRefCallCount := 0
	getServiceRefMockable = func(project *spec.Project, serviceName string) *spec.ServiceConfig {
		getServiceRefCallCount++
		return nil
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Proxy service cannot be found for network inserting", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	GenerateAddProxyNetworks(project, selectedTemplates)
	// Assert
	assert.Equal(t, 1, getServiceRefCallCount)
	assert.Equal(t, 1, printErrorCallCount)
}

// ------------------------------------------------------------------ getServiceRef ----------------------------------------------------------------

func TestGetServiceRef1(t *testing.T) {
	// Test data
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "Service 1",
			},
			{
				Name: "Service 2",
			},
			{
				Name: "Service 3",
			},
		},
	}
	serviceName := "Service 2"
	// Execute test
	result := getServiceRef(project, serviceName)
	// Assert
	assert.Equal(t, &project.Services[1], result)
}

func TestGetServiceRef2(t *testing.T) {
	// Test data
	project := &spec.Project{
		Services: spec.Services{
			{
				Name: "Service 1",
			},
			{
				Name: "Service 2",
			},
			{
				Name: "Service 3",
			},
		},
	}
	serviceName := "Service 4"
	// Execute test
	result := getServiceRef(project, serviceName)
	// Assert
	assert.Nil(t, result)
}
