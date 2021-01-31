package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/kardianos/osext"
	"github.com/sethvargo/go-password/password"

	"compose-generator/model"
)

func ExecuteSafetyFileChecks() {
	isNotSafe := true
	if IsDockerized() {
		isNotSafe = FileExists("/compose-generator/out/docker-compose.yml") || FileExists("/compose-generator/out/environment.env")
	} else {
		isNotSafe = FileExists("docker-compose.yml") || FileExists("environment.env")
	}
	if isNotSafe {
		color.Red("Warning: docker-compose.yml or environment.env already exists. By continuing, you might overwrite those files.")
		result := YesNoQuestion("Do you want to continue?", true)
		if !result {
			os.Exit(0)
		}
		fmt.Println()
	}
}

func IsDockerized() bool {
	return os.Getenv("COMPOSE_GENERATOR_DOCKERIZED") == "1"
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetTemplatesPath() string {
	filename, _ := osext.Executable()
	filename = strings.ReplaceAll(filename, "\\", "/")
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/templates") {
		return filename + "/templates"
	} else if FileExists(filename + "/../templates") {
		return filename + "/../templates"
	} else {
		return "../templates"
	}
}

func GetPredefinedTemplatesPath() string {
	filename, _ := osext.Executable()
	filename = strings.ReplaceAll(filename, "\\", "/")
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/predefined-templates") {
		return filename + "/predefined-templates"
	} else if FileExists(filename + "/../predefined-templates") {
		return filename + "/../predefined-templates"
	} else {
		return "../predefined-templates"
	}
}

func ReplaceVarsInFile(path string, envMap map[string]string) {
	// Read file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		Error("Could not read from "+path, true)
	}

	// Replace variables
	newContent := string(content)
	for key, value := range envMap {
		newContent = strings.ReplaceAll(newContent, "${{"+key+"}}", value)
	}

	// Write content back
	err = ioutil.WriteFile(path, []byte(newContent), 0777)
	if err != nil {
		Error("Could not write to "+path, true)
	}
}

func GenerateSecrets(path string, secrets []model.Secret) map[string]string {
	// Read file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		Error("Could not read from "+path, true)
	}

	// Generate a password for each occurrence of _GENERATE_PW
	newContent := string(content)
	secretsMap := make(map[string]string)
	for _, s := range secrets {
		res, err := password.Generate(s.Length, 10, 0, false, false)
		if err != nil {
			Error("Password generation failed.", true)
		}
		newContent = strings.ReplaceAll(newContent, "${{"+s.Var+"}}", res)
		secretsMap[s.Name] = res
	}

	// Write content back
	err = ioutil.WriteFile(path, []byte(newContent), 0777)
	if err != nil {
		Error("Could not write to "+path+" - "+err.Error(), true)
	}

	return secretsMap
}
