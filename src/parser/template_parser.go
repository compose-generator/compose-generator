package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"compose-generator/model"
	"compose-generator/util"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GetAvailablePredefinedTemplates returns a list of all predefined templates
func GetAvailablePredefinedTemplates() *model.AvailableTemplates {
	// Initialize available templates with empty lists
	availableTemplates := &model.AvailableTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{},
		BackendServices:  []model.PredefinedTemplateConfig{},
		DatabaseServices: []model.PredefinedTemplateConfig{},
		DbAdminServices:  []model.PredefinedTemplateConfig{},
		ProxyService:    []model.PredefinedTemplateConfig{},
		TlsHelperService: []model.PredefinedTemplateConfig{},
	}

	// Find available templates
	templatesPath := util.GetPredefinedServicesPath()
	files, err := ioutil.ReadDir(templatesPath)
	if err != nil {
		util.Error("Could not load predefined service templates", err, true)
	}
	filterFunc := func(s string) bool {
		return s != "README.md" && s != "INSTRUCTIONS_HEADER.md" && s != "predefined-services.tar.gz"
	}

	// Search through service template types
	for _, templateType := range filterFilenames(files, filterFunc) {
		files, err := ioutil.ReadDir(filepath.Join(templatesPath, templateType))
		if err != nil {
			util.Error("Could not load predefined service templates", err, true)
		}
		// Search through service templates
		for _, f := range filterFilenames(files, filterFunc) {
			templatePath := filepath.Join(templatesPath, templateType, f)
			config := getConfigFromFile(templatePath)
			config.Name = f
			config.Type = templateType
			config.Dir = templatePath
			// Save configuration in the selected templates object
			switch templateType {
			case model.TemplateTypeFrontend:
				availableTemplates.FrontendServices = append(availableTemplates.FrontendServices, config)
			case model.TemplateTypeBackend:
				availableTemplates.BackendServices = append(availableTemplates.BackendServices, config)
			case model.TemplateTypeDatabase:
				availableTemplates.DatabaseServices = append(availableTemplates.DatabaseServices, config)
			case model.TemplateTypeDbAdmin:
				availableTemplates.DbAdminServices = append(availableTemplates.DbAdminServices, config)
			case model.TemplateTypeProxy:
				availableTemplates.ProxyService = append(availableTemplates.ProxyService, config)
			case model.TemplateTypeTlsHelper:
				availableTemplates.TlsHelperService = append(availableTemplates.TlsHelperService, config)
			}
		}
	}
	return availableTemplates
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getConfigFromFile(dirPath string) (config model.PredefinedTemplateConfig) {
	// Read JSON file
	jsonFile, err := os.Open(dirPath + "/config.json")
	if err != nil {
		util.Error("Unable to load config file of template "+dirPath, err, true)
	}
	defer jsonFile.Close()

	// Parse json to TemplateConfig struct
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		util.Error("Unable to load config file of template "+dirPath, err, true)
	}
	json.Unmarshal(bytes, &config)
	return
}

func filterFilenames(ss []os.FileInfo, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s.Name()) {
			ret = append(ret, s.Name())
		}
	}
	return
}
