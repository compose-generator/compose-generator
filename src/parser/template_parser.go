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

func ParsePredefinedTemplates() []model.TemplateConfig {
	templates_path := utils.GetPredefinedTemplatesPath()
	files, err := ioutil.ReadDir(templates_path)
	if err != nil {
		utils.Error("Internal error - could not load templates.", true)
	}
	filter_func := func(s string) bool { return strings.Contains(s, "_") && !strings.Contains(s, ".") }
	file_names := filterFilenames(files, filter_func)

	var configs []model.TemplateConfig
	for _, f := range file_names {
		config := getConfigFromFile(filepath.Join(templates_path, f))
		configs = append(configs, config)
	}
	return configs
}

func ParseTemplates() []model.TemplateMetadata {
	templates_path := utils.GetTemplatesPath()
	files, err := ioutil.ReadDir(templates_path)
	if err != nil {
		utils.Error("Internal error - could not load templates.", true)
	}
	var file_names []string
	for _, n := range files {
		file_names = append(file_names, n.Name())
	}

	var metadatas []model.TemplateMetadata
	for _, f := range file_names {
		metadata := getMetadataFromFile(filepath.Join(templates_path, f))
		metadatas = append(metadatas, metadata)
	}
	return metadatas
}

func getConfigFromFile(dir_path string) model.TemplateConfig {
	// Read JSON file
	jsonFile, err := os.Open(dir_path + "/config.json")
	if err != nil {
		utils.Error("Internal error - unable to load config file of template "+dir_path, true)
	}

	// Parse json to TemplateConfig struct
	bytes, _ := ioutil.ReadAll(jsonFile)
	var config model.TemplateConfig
	json.Unmarshal(bytes, &config)

	// Close file
	jsonFile.Close()

	return config
}

func getMetadataFromFile(dir_path string) model.TemplateMetadata {
	// Read JSON file
	jsonFile, err := os.Open(dir_path + "/metadata.json")
	if err != nil {
		utils.Error("Internal error - unable to load metadata file of template "+dir_path, true)
	}

	// Parse json to TemplateMetadata struct
	bytes, _ := ioutil.ReadAll(jsonFile)
	var metadata model.TemplateMetadata
	json.Unmarshal(bytes, &metadata)

	// Close file
	jsonFile.Close()

	return metadata
}

func filterFilenames(ss []os.FileInfo, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s.Name()) {
			ret = append(ret, s.Name())
		}
	}
	return
}
