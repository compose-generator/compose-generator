/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package util

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------ CheckForServiceTemplateUpdate ------------------------------------------

// ------------------------------------------ TemplateListToPreselectedLabelList ------------------------------------------

// ------------------------------------------ TemplateListToLabelList ------------------------------------------

func TestTemplateListToLabelList(t *testing.T) {
	templates := []model.PredefinedTemplateConfig{
		{Label: "Angular"},
		{Label: "Wordpress"},
		{Label: "MySQL"},
	}
	result := TemplateListToLabelList(templates)
	assert.EqualValues(t, result, []string{"Angular", "Wordpress", "MySQL"})
}
