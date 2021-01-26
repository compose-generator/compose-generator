package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"compose-generator/model"
	"compose-generator/utils"
)

func ParseTemplates() []model.Config {
	templates_path := utils.GetTemplatesPath()
	files, err := ioutil.ReadDir(templates_path)
	if err != nil {
		utils.Error("Internal error - could not load templates.", true)
	}
	filter_func := func(s string) bool { return strings.Contains(s, "_") && !strings.Contains(s, ".") }
	file_names := filterFilenames(files, filter_func)

	var configs []model.Config
	for _, f := range file_names {
		fmt.Println(f)
		config := getConfigFromFile(filepath.Join(templates_path, f))
		configs = append(configs, config)
	}
	return configs
}

func getConfigFromFile(dir_path string) model.Config {
	// Read JSON file
	jsonFile, err := os.Open(dir_path + "/config.json")
	if err != nil {
		utils.Error("Internal error - unable to load config file of template "+dir_path+".", true)
	}

	// Parse json to Config struct
	bytes, _ := ioutil.ReadAll(jsonFile)
	var config model.Config
	json.Unmarshal(bytes, &config)

	// Close file
	jsonFile.Close()

	return config
}

func filterFilenames(ss []os.FileInfo, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s.Name()) {
			ret = append(ret, s.Name())
		}
	}
	return
}
