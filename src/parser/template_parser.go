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

// ParsePredefinedServices returns a list of all predefined templates
func ParsePredefinedServices() map[string][]model.ServiceTemplateConfig {
	templatesPath := util.GetPredefinedServicesPath()
	files, err := ioutil.ReadDir(templatesPath)
	if err != nil {
		util.Error("Internal error - could not load service templates.", err, true)
	}
	filterFunc := func(s string) bool {
		return s != "README.md" && s != "INSTRUCTIONS_HEADER.md" && s != "predefined-services.tar.gz"
	}

	configs := make(map[string][]model.ServiceTemplateConfig)
	for _, templateType := range filterFilenames(files, filterFunc) {
		files, err := ioutil.ReadDir(filepath.Join(templatesPath, templateType))
		if err != nil {
			util.Error("Internal error - could not load service templates.", err, true)
		}
		for _, f := range filterFilenames(files, filterFunc) {
			templatePath := filepath.Join(templatesPath, templateType, f)
			config := getConfigFromFile(templatePath)
			config.Name = f
			config.Type = templateType
			config.Dir = filepath.Join(templateType, f)
			if configs[templateType] != nil {
				configs[templateType] = append(configs[templateType], config)
				continue
			}
			configs[templateType] = []model.ServiceTemplateConfig{config}
		}
	}
	return configs
}

// ParseCustomTemplates returns a list of all custom templates
func ParseCustomTemplates() (metadatas []model.TemplateMetadata) {
	templatesPath := util.GetCustomTemplatesPath()
	// Create templates dir, if it not exists
	if !util.FileExists(templatesPath) {
		os.MkdirAll(templatesPath, 0777)
	}
	// Read template directory names
	files, err := ioutil.ReadDir(templatesPath)
	if err != nil {
		util.Error("Internal error - could not load templates", err, true)
	}
	var fileNames []string
	for _, n := range files {
		fileNames = append(fileNames, n.Name())
	}
	// Read template metadata
	for _, f := range fileNames {
		metadata := getMetadataFromFile(filepath.Join(templatesPath, f))
		metadatas = append(metadatas, metadata)
	}
	return
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getConfigFromFile(dirPath string) (config model.ServiceTemplateConfig) {
	// Read JSON file
	jsonFile, err := os.Open(dirPath + "/config.json")
	if err != nil {
		util.Error("Internal error - unable to load config file of template "+dirPath, err, true)
	}
	defer jsonFile.Close()

	// Parse json to TemplateConfig struct
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &config)
	return
}

func getMetadataFromFile(dirPath string) (metadata model.TemplateMetadata) {
	// Read JSON file
	jsonFile, err := os.Open(dirPath + "/metadata.json")
	if err != nil {
		util.Error("Internal error - unable to load metadata file of template "+dirPath, err, true)
	}
	defer jsonFile.Close()

	// Parse json to TemplateMetadata struct
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &metadata)
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
