package util

import (
	"compose-generator/model"
	"strings"

	"github.com/PaesslerAG/gval"
)

// EvaluateConditionalSections takes in a string, searches for comments and uncomments it based on a condition
func EvaluateConditionalSections(
	content string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) string {
	rows := strings.Split(content, "\n")
	uncommenting := false
	var cleanRows []string = []string{}
	// Evaluate conditions and uncomment lines
	for _, row := range rows {
		if strings.HasPrefix(row, "#! if ") {
			// Conditional section found -> check condition
			conditions := strings.Split(row[6:strings.Index(row, " {")], "|")
			for _, c := range conditions {
				if EvaluateCondition(c, templateData, varMap) {
					uncommenting = true
					break
				}
			}
		} else if strings.HasPrefix(row, "#! }") {
			uncommenting = false
		} else if uncommenting {
			cleanRows = append(cleanRows, row[3:])
		} else if !strings.HasPrefix(row, "#! ") {
			cleanRows = append(cleanRows, row)
		}
	}
	return strings.Join(cleanRows, "\n")
}

// EvaluateCondition evaluates the given condition to a boolean result
func EvaluateCondition(
	condition string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) bool {
	if strings.HasPrefix(condition, "has service ") {
		for _, templates := range templateData {
			for _, template := range templates {
				if template.Name == condition[12:] {
					return true
				}
			}
		}
	} else if strings.HasPrefix(condition, "has ") {
		return len(templateData[condition[4:]]) > 0
	} else if strings.HasPrefix(condition, "var.") {
		condition = condition[4:]
		params := make(map[string]interface{})
		for varName, varValue := range varMap {
			params[varName] = varValue
		}
		result, err := gval.Evaluate(condition, params)
		return result.(bool) && err == nil
	}
	return false
}
