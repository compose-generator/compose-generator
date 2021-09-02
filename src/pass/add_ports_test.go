package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/AlecAivazis/survey/v2"
	"github.com/stretchr/testify/assert"

	spec "github.com/compose-spec/compose-go/types"
)

func TestAddPorts1(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{}
	expectedService := &spec.ServiceConfig{
		Ports: []spec.ServicePortConfig{
			{
				Mode:      "ingress",
				Protocol:  "tcp",
				Target:    80,
				Published: 9090,
			},
			{
				Mode:      "ingress",
				Protocol:  "tcp",
				Target:    81,
				Published: 8081,
			},
		},
	}
	// Mock functions
	pelCallCount := 0
	Pel = func() {
		pelCallCount++
	}
	yesNoCallCount := 0
	YesNoQuestion = func(question string, defaultValue bool) bool {
		yesNoCallCount++
		switch yesNoCallCount {
		case 1:
			assert.Equal(t, "Do you want to expose ports of your service?", question)
			assert.False(t, defaultValue)
		case 2:
			assert.Equal(t, "Expose another port?", question)
			assert.True(t, defaultValue)
		case 3:
			assert.Equal(t, "Expose another port?", question)
			assert.True(t, defaultValue)
			return false
		}
		return true
	}
	textQuestionCallCount := 0
	TextQuestionWithValidator = func(question string, fn survey.Validator) string {
		textQuestionCallCount++
		switch textQuestionCallCount {
		case 1:
			assert.Equal(t, "Which port do you want to expose? (inner port)", question)
			return "80"
		case 2:
			assert.Equal(t, "To which destination port on the host machine?", question)
			return "9090"
		case 3:
			assert.Equal(t, "Which port do you want to expose? (inner port)", question)
			return "81"
		case 4:
			assert.Equal(t, "To which destination port on the host machine?", question)
			return "8081"
		}
		return ""
	}
	// Exeute test
	AddPorts(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, 2, pelCallCount)
}

func TestAddPorts2(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{}
	expectedService := &spec.ServiceConfig{}
	// Mock functions
	pelCallCount := 0
	Pel = func() {
		pelCallCount++
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to expose ports of your service?", question)
		assert.False(t, defaultValue)
		return false
	}
	// Exeute test
	AddPorts(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Zero(t, pelCallCount)
}
