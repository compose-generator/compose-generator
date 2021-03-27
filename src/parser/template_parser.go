package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"compose-generator/model"
	"compose-generator/utils"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// ParsePredefinedServices returns a list of all predefined templates
func ParsePredefinedServices() (configs []model.TemplateConfig) {
	templatesPath := utils.GetPredefinedServicesPath()
	files, err := ioutil.ReadDir(templatesPath)
	if err != nil {
		utils.Error("Internal error - could not load templates.", true)
	}
	filterFunc := func(s string) bool { return strings.Contains(s, "_") && !strings.Contains(s, ".") }
	fileNames := filterFilenames(files, filterFunc)

	for _, f := range fileNames {
		config := getConfigFromFile(filepath.Join(templatesPath, f))
		configs = append(configs, config)
	}
	return
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

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getConfigFromFile(dirPath string) (config model.TemplateConfig) {
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
