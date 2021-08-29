package util

import (
	"compose-generator/model"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// CheckForServiceTemplateUpdate checks if any updates are available for the predefined service templates
func CheckForServiceTemplateUpdate() {
	// Skip on dev version
	if Version == "dev" {
		return
	}
	// Create predefined templates dir if not exitsts
	predefinedTemplatesDir := GetPredefinedServicesPath()
	if !FileExists(predefinedTemplatesDir) {
		if err := os.MkdirAll(predefinedTemplatesDir, 0755); err != nil {
			Error("Could not create directory for predefined templates", err, true)
		}
	}

	fileUrl := "https://github.com/compose-generator/compose-generator/releases/download/" + Version + "/predefined-services.tar.gz"
	outputPath := GetPredefinedServicesPath() + "/predefined-services.tar.gz"
	shouldUpdate := false

	if FileExists(outputPath) { // File exists => version check
		file, err := os.Stat(outputPath)
		if err != nil {
			Error("Could not access existing template archive", err, true)
		}
		lastModifiedLocal := file.ModTime().Unix()

		// Issue HEAD request for services archive
		res, err := http.Head(fileUrl)
		if err != nil {
			Warning("Could not check for template updates")
			return
		}
		lastModified := res.Header["Last-Modified"][0]
		t, err := time.Parse(time.RFC1123, lastModified)
		if err != nil {
			Error("Cannot parse last modified of remote file", err, true)
		}
		if t.Unix() > lastModifiedLocal {
			shouldUpdate = true
		}
	} else { // File does not exist => download directly
		shouldUpdate = true
	}

	// Download update if necessary
	if shouldUpdate {
		P("Downloading predefined services update ... ")
		DownloadFile(fileUrl, outputPath)
		filepath, err := filepath.Abs(predefinedTemplatesDir)
		if err != nil {
			Error("Could not build path", err, true)
		}
		ExecuteOnLinuxWithCustomVolume("tar xfvz predefined-services.tar.gz", filepath)
		Done()
	}
}

// Asks the user all questions the predefined service contains and saves the answers to the project
func AskTemplateQuestions(project *model.CGProject, template *model.PredefinedTemplateConfig) {
	for _, question := range template.Questions {
		defaultValue := ReplaceVarsInString(question.DefaultValue, project.Vars)
		// Only ask advanced questions when the project was created in advanced mode
		if project.AdvancedConfig || !question.Advanced {
			// Question can be answered
			switch question.Type {
			case model.QuestionTypeYesNo:
				// Ask a yes/no question
				defaultValue, err := strconv.ParseBool(defaultValue)
				if err != nil {
					Error("Mistake in predefined template '"+template.Name+"'. Default value of yes/no question was no bool", err, true)
				}
				answer := YesNoQuestion(question.Text, defaultValue)
				project.Vars[question.Variable] = strconv.FormatBool(answer)
			case model.QuestionTypeText:
				// Ask a text question
				answer := ""
				if question.Validator != "" {
					// Ask a text question with validator
					validator := GetValidatorByName(question.Validator)
					answer = TextQuestionWithDefaultAndValidator(question.Text, defaultValue, validator)
				} else {
					// Ask a text question without validator
					answer = TextQuestionWithDefault(question.Text, defaultValue)
				}
				project.Vars[question.Variable] = answer
			case model.QuestionTypeMenu:
				// Ask a menu question
				answer := MenuQuestionWithDefault(question.Text, question.Options, question.DefaultValue)
				project.Vars[question.Variable] = answer
			}
		} else {
			// Advanced question falls back to default value
			project.Vars[question.Variable] = question.DefaultValue
		}
	}
}

/*
func getVarMapFromQuestions(
	varMap *map[string]string,
	usedPorts *[]int,
	questions []model.Question,
	flagAdvanced bool,
) {
	for _, q := range questions {
		defaultValue := util.ReplaceVarsInString(q.DefaultValue, *varMap)
		if !q.Advanced || (q.Advanced && flagAdvanced) {
			switch q.Type {
			case 1: // Yes/No
				defaultValue, _ := strconv.ParseBool(defaultValue)
				(*varMap)[q.Variable] = strconv.FormatBool(util.YesNoQuestion(q.Text, defaultValue))
			case 2: // Text
				if q.Validator != "" {
					var customValidator survey.Validator
					switch q.Validator {
					case "port":
						customValidator = util.PortValidator
						// Check if port was already assigned
						port, _ := strconv.Atoi(defaultValue)
						for util.SliceContainsInt(*usedPorts, port) {
							port = port + 1
						}
						defaultValue = strconv.Itoa(port)
					default:
						customValidator = func(val interface{}) error {
							validate := validator.New()
							if validate.Var(val.(string), "required,"+q.Validator) != nil {
								return errors.New("please provide a valid input")
							}
							return nil
						}
					}
					answer := util.TextQuestionWithDefaultAndValidator(q.Text, defaultValue, customValidator)
					(*varMap)[q.Variable] = answer
					if q.Validator == "port" {
						port, _ := strconv.Atoi(answer)
						*usedPorts = append(*usedPorts, port)
					}
				} else {
					(*varMap)[q.Variable] = util.TextQuestionWithDefault(q.Text, defaultValue)
				}
			case 3: // Select
				answer := util.MenuQuestionWithDefault(q.Text, q.Options, q.DefaultValue)
				(*varMap)[q.Variable] = answer
			}
		} else {
			(*varMap)[q.Variable] = defaultValue
		}
	}
}
*/

// TemplateListToLabelList converts a slice of ServiceTemplateConfig to a slice of labels
func TemplateListToLabelList(templates []model.PredefinedTemplateConfig) (labels []string) {
	for _, t := range templates {
		labels = append(labels, t.Label)
	}
	return
}

// TemplateListToPreselectedLabelList retrieves a slice of all preselected other services for each service
func TemplateListToPreselectedLabelList(templateList []model.PredefinedTemplateConfig, selected *model.SelectedTemplates) (labels []string) {
	for _, t := range templateList {
		if t.Preselected == "false" {
			continue
		}
		if t.Preselected == "true" || EvaluateCondition(t.Preselected, selected, nil) {
			labels = append(labels, t.Label)
		}
	}
	return
}
