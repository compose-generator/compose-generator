/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"compose-generator/model"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// CheckForServiceTemplateUpdate checks if any updates are available for the predefined service templates
func CheckForServiceTemplateUpdate() {
	// Skip on dev version or dockerized
	if IsDevVersion() || IsDockerizedEnvironment() {
		return
	}
	InfoLogger.Println("Checking for predefined service template update ...")
	// Create predefined templates dir if not exitsts
	predefinedTemplatesDir := GetPredefinedServicesPath()
	spinner := StartProcess("Checking for predefined service template updates ...")
	if !FileExists(predefinedTemplatesDir) {
		if err := os.MkdirAll(predefinedTemplatesDir, 0750); err != nil {
			ErrorLogger.Println("Could not create directory for predefined templates: " + err.Error())
			logError("Could not create directory for predefined templates", true)
		}
	}

	fileUrl := "https://github.com/compose-generator/compose-generator/releases/download/" + Version + "/predefined-services.tar.gz"
	outputPath := GetPredefinedServicesPath() + "/predefined-services.tar.gz"
	shouldUpdate := false

	if FileExists(outputPath) { // File exists => version check
		file, err := os.Stat(outputPath)
		if err != nil {
			ErrorLogger.Println("Could not access existing template archive: " + err.Error())
			logError("Could not access existing template archive", true)
		}
		lastModifiedLocal := file.ModTime().Unix()

		// Issue HEAD request for services archive
		res, err := http.Head(fileUrl)
		if err != nil {
			WarningLogger.Println("Could not check for template updates: " + err.Error())
			logWarning("Could not check for template updates")
			return
		}
		lastModified := res.Header["Last-Modified"][0]
		t, err := time.Parse(time.RFC1123, lastModified)
		if err != nil {
			ErrorLogger.Println("Cannot parse last modified of remote file: " + err.Error())
			logError("Cannot parse last modified of remote file", true)
		}
		if t.Unix() > lastModifiedLocal {
			shouldUpdate = true
		}
	} else { // File does not exist => download directly
		shouldUpdate = true
	}
	StopProcess(spinner)
	InfoLogger.Println("Checking for predefined service template update (done)")

	// Download update if necessary
	if shouldUpdate {
		if IsPrivileged() {
			// Download predefined services update
			processMessage := "Downloading predefined services update and the toolbox image (this can take a while) ..."
			InfoLogger.Println("Download predefined service template update ...")
			if IsToolboxPresent() {
				processMessage = "Downloading predefined services update ..."
			}
			spinner = StartProcess(processMessage)
			if err := DownloadFile(fileUrl, outputPath); err != nil {
				ErrorLogger.Println("Failed to download predefined services update: " + err.Error())
				logError("Failed to download predefined services update. Please check your internet connection", true)
			}
			filepath, err := filepath.Abs(predefinedTemplatesDir)
			if err != nil {
				ErrorLogger.Println("Could not build path: " + err.Error())
				logError("Could not build path", true)
			}
			ExecuteOnToolboxCustomVolume("tar xfvz predefined-services.tar.gz", filepath)
			StopProcess(spinner)
			InfoLogger.Println("Download predefined service template update (done)")
		} else {
			InfoLogger.Println("Predefined services update found")
			logError("Predefined services update found. Root privileges are required to install the update. Please run Compose Generator again with elevated privileges", true)
		}
	}
}

// AskTemplateQuestions asks the user all questions the predefined service contains and saves the answers to the project
func AskTemplateQuestions(project *model.CGProject, template *model.PredefinedTemplateConfig) {
	for _, question := range template.Questions {
		text := ReplaceVarsInString(question.Text, project.Vars)
		defaultValue := ReplaceVarsInString(question.DefaultValue, project.Vars)

		if question.Validator == "port" {
			// If the port is already in use, find unused one
			port, err := strconv.Atoi(defaultValue)
			if err != nil {
				ErrorLogger.Println("Could not convert port to integer: " + err.Error())
				logError("Could not convert port to integer. Please check template", true)
			}
			for SliceContainsInt(project.Ports, port) {
				port++
			}
			defaultValue = strconv.Itoa(port)
		}
		// Only ask advanced questions when the project was created in advanced mode
		if project.AdvancedConfig || !question.Advanced {
			// Question can be answered
			switch question.Type {
			case model.QuestionTypeYesNo:
				// Ask a yes/no question
				defaultValue, err := strconv.ParseBool(defaultValue)
				if err != nil {
					ErrorLogger.Println("Default value of yes/no was no bool in '" + template.Name + "': " + err.Error())
					logError("Mistake in predefined template '"+template.Name+"'. Default value of yes/no question was no bool", true)
				}
				answer := strconv.FormatBool(YesNoQuestion(text, defaultValue))
				InfoLogger.Println("User chose: " + question.Variable + "=" + answer)
				project.Vars[question.Variable] = answer
			case model.QuestionTypeText:
				// Ask a text question
				answer := ""
				if question.Validator != "" {
					// Ask a text question with validator
					validator := GetValidatorByName(question.Validator)
					answer = TextQuestionWithDefaultAndValidator(text, defaultValue, validator)
					if question.Validator == "port" {
						port, err := strconv.Atoi(answer)
						if err != nil {
							ErrorLogger.Println("Could not convert port to integer: " + err.Error())
							logError("Could not convert port to integer. Please check template", true)
						}
						project.Ports = append(project.Ports, port)
					}
				} else {
					// Ask a text question without validator
					answer = TextQuestionWithDefault(text, defaultValue)
				}
				InfoLogger.Println("User chose: " + question.Variable + "=" + answer)
				project.Vars[question.Variable] = answer
			case model.QuestionTypeMenu:
				// Ask a menu question
				answer := MenuQuestionWithDefault(text, question.Options, question.DefaultValue)
				InfoLogger.Println("User chose: " + question.Variable + "=" + answer)
				project.Vars[question.Variable] = answer
			}
		} else {
			// Advanced question falls back to default value
			InfoLogger.Println("Falling back to default: " + question.Variable + "=" + question.DefaultValue)
			project.Vars[question.Variable] = question.DefaultValue
		}
	}
}

// AskTemplateProxyQuestions aks the user all collected proxy questions for every selected service
func AskTemplateProxyQuestions(project *model.CGProject, template *model.PredefinedTemplateConfig, selectedTemplates *model.SelectedTemplates) {
	// Ask proxy questions only if the service wants to get proxied
	if template.Proxied {
		proxyVars := make(model.Vars)
		for _, question := range selectedTemplates.GetAllProxyQuestions() {
			// Replace vars
			text := ReplaceVarsInString(question.Text, project.Vars)
			defaultValue := ReplaceVarsInString(question.DefaultValue, project.Vars)
			// Replace current service variables
			text = strings.ReplaceAll(text, "${{CURRENT_SERVICE_LABEL}}", template.Label)
			text = strings.ReplaceAll(text, "${{CURRENT_SERVICE_NAME}}", template.Name)
			defaultValue = strings.ReplaceAll(defaultValue, "${{CURRENT_SERVICE_LABEL}}", template.Label)
			defaultValue = strings.ReplaceAll(defaultValue, "${{CURRENT_SERVICE_NAME}}", template.Name)

			// Only ask advanced questions when the project was created in advanced mode
			if project.AdvancedConfig || !question.Advanced {
				// Question can be answered
				switch question.Type {
				case model.QuestionTypeYesNo:
					// Ask a yes/no question
					defaultValue, err := strconv.ParseBool(defaultValue)
					if err != nil {
						ErrorLogger.Println("Default value of yes/no was no bool in '" + template.Name + "': " + err.Error())
						logError("Mistake in proxy question configuration. Default value of yes/no question was no bool", true)
					}
					answer := strconv.FormatBool(YesNoQuestion(text, defaultValue))
					proxyVars[question.Variable] = answer
					InfoLogger.Println("User chose: " + question.Variable + "=" + answer)
				case model.QuestionTypeText:
					// Ask a text question
					answer := ""
					if question.Validator != "" {
						// Ask a text question with validator
						validator := GetValidatorByName(question.Validator)
						answer = TextQuestionWithDefaultAndValidator(text, defaultValue, validator)
						if question.Validator == "port" {
							port, err := strconv.Atoi(answer)
							if err != nil {
								ErrorLogger.Println("Could not convert port to integer: " + err.Error())
								logError("Could not convert port to integer. Please check template", true)
							}
							project.Ports = append(project.Ports, port)
						}
					} else {
						// Ask a text question without validator
						answer = TextQuestionWithDefault(text, defaultValue)
					}
					proxyVars[question.Variable] = answer
					InfoLogger.Println("User chose: " + question.Variable + "=" + answer)
				case model.QuestionTypeMenu:
					// Ask a menu question
					answer := MenuQuestionWithDefault(text, question.Options, question.DefaultValue)
					proxyVars[question.Variable] = answer
					InfoLogger.Println("User chose: " + question.Variable + "=" + answer)
				}
			} else {
				// Advanced question falls back to default value
				proxyVars[question.Variable] = question.DefaultValue
				InfoLogger.Println("Falling back to default: " + question.Variable + "=" + question.DefaultValue)
			}
		}
		// Add collected proxy vars to project
		project.ProxyVars[template.Name] = proxyVars
	}
}

// EvaluateProxyLabels adds proxy labels with values to the project based on the answers of the proxy questions
func EvaluateProxyLabels(project *model.CGProject, template *model.PredefinedTemplateConfig, selectedTemplates *model.SelectedTemplates) {
	if template.Proxied {
		proxyLabels := make(model.Labels)
		for _, label := range selectedTemplates.GetAllProxyLabels() {
			// Replace vars
			name := ReplaceVarsInString(label.Name, project.Vars)
			value := ReplaceVarsInString(label.Value, project.Vars)
			// Replace current service variables
			name = strings.ReplaceAll(name, "${{CURRENT_SERVICE_LABEL}}", template.Label)
			name = strings.ReplaceAll(name, "${{CURRENT_SERVICE_NAME}}", template.Name)
			value = strings.ReplaceAll(value, "${{CURRENT_SERVICE_LABEL}}", template.Label)
			value = strings.ReplaceAll(value, "${{CURRENT_SERVICE_NAME}}", template.Name)

			proxyLabels[name] = value
			InfoLogger.Println("Specified label: " + label.Name + "=" + label.Value)
		}
		// Add collected proxy labels to project
		project.ProxyLabels[template.Name] = proxyLabels
	}
}

// AskForCustomVolumePaths asks the user for custom volume paths for a template
func AskForCustomVolumePaths(project *model.CGProject, template *model.PredefinedTemplateConfig) {
	for _, volume := range template.Volumes {
		defaultValue := ReplaceVarsInString(volume.DefaultValue, project.Vars)
		// Only ask advanced questions when the project was created in advanced mode
		if project.AdvancedConfig || !volume.Advanced {
			answer := ""
			// Ask a text question with validator
			answer = TextQuestionWithDefault(volume.Text, defaultValue)
			project.Vars[volume.Variable] = answer
			InfoLogger.Println("User chose: " + volume.Variable + "=" + answer)
		} else {
			// Advanced question falls back to default value
			project.Vars[volume.Variable] = volume.DefaultValue
			InfoLogger.Println("Falling back to default: " + volume.Variable + "=" + volume.DefaultValue)
		}
	}
}

// TemplateListToLabelList converts a slice of ServiceTemplateConfig to a slice of labels
func TemplateListToLabelList(templates []model.PredefinedTemplateConfig) (labels []string) {
	for _, t := range templates {
		labels = append(labels, t.Label)
	}
	return
}

// TemplateListToPreselectedLabelList retrieves a slice of all preselected other services for each service
func TemplateListToPreselectedLabelList(templates []model.PredefinedTemplateConfig, selected *model.SelectedTemplates) (labels []string) {
	for _, t := range templates {
		if t.Preselected == "false" {
			continue
		}
		if t.Preselected == "true" || EvaluateCondition(t.Preselected, selected, nil) {
			labels = append(labels, t.Label)
		}
	}
	return
}
