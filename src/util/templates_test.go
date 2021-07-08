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
	templates := []model.ServiceTemplateConfig{
		{Label: "Angular"},
		{Label: "Wordpress"},
		{Label: "MySQL"},
	}
	result := TemplateListToLabelList(templates)
	assert.EqualValues(t, result, []string{"Angular", "Wordpress", "MySQL"})
}
