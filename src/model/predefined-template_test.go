/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------- GetAll ------------------------------------------------------------------

func TestGetAll(t *testing.T) {
	// Test data
	selectedTemplates := SelectedTemplates{
		FrontendServices: []PredefinedTemplateConfig{
			{Name: "angular"},
			{Name: "vue"},
		},
		BackendServices: []PredefinedTemplateConfig{
			{Name: "django"},
			{Name: "flask"},
			{Name: "rocket"},
		},
		DatabaseServices: []PredefinedTemplateConfig{
			{Name: "faunadb"},
			{Name: "mongodb"},
		},
		DbAdminServices: []PredefinedTemplateConfig{
			{Name: "pgadmin"},
		},
		ProxyServices: []PredefinedTemplateConfig{
			{Name: "nginx"},
		},
		TlsHelperServices: []PredefinedTemplateConfig{
			{Name: "letsencrypt"},
		},
	}
	expectedResult := []PredefinedTemplateConfig{
		{Name: "angular"},
		{Name: "vue"},
		{Name: "django"},
		{Name: "flask"},
		{Name: "rocket"},
		{Name: "faunadb"},
		{Name: "mongodb"},
		{Name: "pgadmin"},
		{Name: "nginx"},
		{Name: "letsencrypt"},
	}
	// Execute test
	result := selectedTemplates.GetAll()
	// Assert
	assert.Equal(t, 10, len(result))
	assert.EqualValues(t, expectedResult, result)
}

// --------------------------------------------------------------------- GetAllRef -----------------------------------------------------------------

func TestGetAllRef(t *testing.T) {
	// Test data
	t.Fail()
	// Execute test

	// Assert
}

// -------------------------------------------------------------------- GetAllLabels ---------------------------------------------------------------

func TestGetAllLabels(t *testing.T) {
	// Test data
	selectedTemplates := SelectedTemplates{
		FrontendServices: []PredefinedTemplateConfig{
			{
				Name:  "drupal",
				Label: "Drupal",
			},
			{
				Name:  "mediawiki",
				Label: "Mediawiki",
			},
		},
		BackendServices: []PredefinedTemplateConfig{
			{
				Name:  "fiber",
				Label: "Fiber",
			},
			{
				Name:  "nexus",
				Label: "Nexus",
			},
		},
		DatabaseServices: []PredefinedTemplateConfig{
			{
				Name:  "scylladb",
				Label: "ScyllaDB",
			},
			{
				Name:  "singlestore",
				Label: "SingleStore",
			},
		},
		DbAdminServices: []PredefinedTemplateConfig{
			{
				Name:  "elasticsearch-hq",
				Label: "Elasticsearch HQ",
			},
			{
				Name:  "redis-insight",
				Label: "Redis Insight",
			},
		},
		ProxyServices: []PredefinedTemplateConfig{
			{
				Name:  "traefik",
				Label: "Traefik",
			},
		},
		TlsHelperServices: []PredefinedTemplateConfig{},
	}
	expectedResult := []string{"Drupal", "Mediawiki", "Fiber", "Nexus", "ScyllaDB", "SingleStore", "Elasticsearch HQ", "Redis Insight", "Traefik"}
	// Execute test
	result := selectedTemplates.GetAllLabels()
	// Assert
	assert.Equal(t, 9, len(result))
	assert.EqualValues(t, expectedResult, result)
}

// ---------------------------------------------------------------------- GetTotal -----------------------------------------------------------------

func TestTotal(t *testing.T) {
	// Test data
	selectedTemplates := SelectedTemplates{
		FrontendServices: []PredefinedTemplateConfig{
			{Name: "angular"},
			{Name: "vue"},
		},
		BackendServices: []PredefinedTemplateConfig{
			{Name: "django"},
			{Name: "flask"},
			{Name: "rocket"},
		},
		DatabaseServices: []PredefinedTemplateConfig{
			{Name: "faunadb"},
			{Name: "mongodb"},
		},
		DbAdminServices: []PredefinedTemplateConfig{
			{Name: "pgadmin"},
		},
	}
	// Execute test
	result := selectedTemplates.GetTotal()
	// Assert
	assert.Equal(t, 8, result)
}

// ----------------------------------------------------------------- GetFilePathsByType ------------------------------------------------------------

func TestGetFilePathsByType(t *testing.T) {
	// Test data
	templateConfig := &PredefinedTemplateConfig{
		Files: []File{
			{
				Type: FileTypeConfig,
				Path: "./test-file1.conf",
			},
			{
				Type: FileTypeDocs,
				Path: "./README.md",
			},
			{
				Type: FileTypeEnv,
				Path: "./environment.env",
			},
			{
				Type: FileTypeConfig,
				Path: "./volumes/configuration.txt",
			},
		},
	}
	expectedResult := []string{"./test-file1.conf", "./volumes/configuration.txt"}
	// Execute test
	result := templateConfig.GetFilePathsByType(FileTypeConfig)
	// Assert
	assert.Equal(t, 2, len(result))
	assert.EqualValues(t, expectedResult, result)
}
