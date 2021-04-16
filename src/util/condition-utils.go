package util

import (
	"compose-generator/model"
	"strings"

	"github.com/Knetic/govaluate"
)

// EvaluateConditionalSections takes in a string, searches for comments and uncomments it based on a condition
func EvaluateConditionalSections(
	content string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) string {
	rows := strings.Split(content, "\n")
	uncommenting := false
	for i, row := range rows {
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
			rows[i] = row[3:]
		}
	}
	return strings.Join(rows, "\n")
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
		expr, err1 := govaluate.NewEvaluableExpression(condition)
		parameters := make(map[string]interface{}, 8)
		for varName, varValue := range varMap {
			parameters[varName] = varValue
		}
		result, err2 := expr.Evaluate(parameters)
		return result.(bool) && err1 != nil && err2 != nil
	}
	return false
}
