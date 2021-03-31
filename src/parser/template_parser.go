package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"compose-generator/model"
	"compose-generator/utils"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// ParsePredefinedServices returns a list of all predefined templates
func ParsePredefinedServices() map[string][]model.ServiceTemplateConfig {
	templatesPath := utils.GetPredefinedServicesPath()
	files, err := ioutil.ReadDir(templatesPath)
	if err != nil {
		utils.Error("Internal error - could not load service templates.", true)
	}
	filterFunc := func(s string) bool { return s != "README.md" }

	configs := make(map[string][]model.ServiceTemplateConfig)
	for _, templateType := range filterFilenames(files, filterFunc) {
		files, err := ioutil.ReadDir(filepath.Join(templatesPath, templateType))
		if err != nil {
			utils.Error("Internal error - could not load service templates.", true)
		}
		for _, f := range filterFilenames(files, filterFunc) {
			templatePath := filepath.Join(templatesPath, templateType, f)
			config := getConfigFromFile(templatePath)
			config.Name = f
			config.Type = templateType
			config.Dir = filepath.Join(templateType, f)
			if configs[templateType] != nil {
				configs[templateType] = append(configs[templateType], config)
			} else {
				configs[templateType] = []model.ServiceTemplateConfig{config}
			}
		}
	}
	return configs
}

// ParseTemplates returns a list of all custom templates
func ParseTemplates() (metadatas []model.TemplateMetadata) {
	templatesPath := utils.GetTemplatesPath()
	files, err := ioutil.ReadDir(templatesPath)
	if err != nil {
		utils.Error("Internal error - could not load templates.", true)
	}
	var fileNames []string
	for _, n := range files {
		fileNames = append(fileNames, n.Name())
	}

	for _, f := range fileNames {
		metadata := getMetadataFromFile(filepath.Join(templatesPath, f))
		metadatas = append(metadatas, metadata)
	}
	return
}

// TemplateListToTemplateLabelList converts a list of service templates to a list of labels
func TemplateListToTemplateLabelList(templates []model.ServiceTemplateConfig) (labels []string) {
	for _, t := range templates {
		labels = append(labels, t.Label)
	}
	return
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getConfigFromFile(dirPath string) (config model.ServiceTemplateConfig) {
	// Read JSON file
	jsonFile, err := os.Open(dirPath + "/config.json")
	if err != nil {
		utils.Error("Internal error - unable to load config file of template "+dirPath, true)
	}

	// Parse json to TemplateConfig struct
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &config)

	// Close file
	jsonFile.Close()
	return
}

func getMetadataFromFile(dirPath string) (metadata model.TemplateMetadata) {
	// Read JSON file
	jsonFile, err := os.Open(dirPath + "/metadata.json")
	if err != nil {
		utils.Error("Internal error - unable to load metadata file of template "+dirPath, true)
	}

	// Parse json to TemplateMetadata struct
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &metadata)

	// Close file
	jsonFile.Close()
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
