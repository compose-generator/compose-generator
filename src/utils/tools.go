package utils

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/kardianos/osext"
	"github.com/sethvargo/go-password/password"

	"compose-generator/model"
)

// PrintSafetyWarning checks if commonly used files are already existing and warns the user about it
func PrintSafetyWarning(existingCount int) {
	Pel()
	Warning(strconv.Itoa(existingCount) + " output files already exist. By continuing, those files will be overwritten!")
	result := YesNoQuestion("Do you want to continue?", true)
	if !result {
		os.Exit(0)
	}
	Pel()
}

// IsDockerized checks if Compose Generator runs within a dockerized environment
func IsDockerized() bool {
	return os.Getenv("COMPOSE_GENERATOR_DOCKERIZED") == "1"
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDirectory checks if a file is a directory
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// GetTemplatesPath returns the path to the custom templates directory
func GetTemplatesPath() string {
	if FileExists("/usr/lib/compose-generator/templates") {
		return "/usr/lib/compose-generator/templates" // Linux
	}
	filename, _ := osext.Executable()
	filename = strings.ReplaceAll(filename, "\\", "/")
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/templates") {
		return filename + "/templates" // Windows + Docker
	}
	return "../templates" // Dev
}

// GetPredefinedServicesPath returns the path to the predefined services directory
func GetPredefinedServicesPath() string {
	if FileExists("/usr/lib/compose-generator/predefined-services") {
		return "/usr/lib/compose-generator/predefined-services" // Linux
	}
	filename, _ := osext.Executable()
	filename = strings.ReplaceAll(filename, "\\", "/")
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/predefined-services") {
		return filename + "/predefined-services" // Windows + Docker
	}
	return "../predefined-services" // Dev
}

func getToolboxPath() string {
	if FileExists("/usr/lib/compose-generator/toolbox") {
		return "/usr/lib/compose-generator/toolbox" // Linux
	}
	filename, _ := osext.Executable()
	filename = strings.ReplaceAll(filename, "\\", "/")
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/toolbox") {
		return filename + "/toolbox" // Windows + Docker
	}
	return "../toolbox" // Dev
}

// ReplaceVarsInFile replaces all variables in the stated file with the contents of the map
func ReplaceVarsInFile(path string, varMap map[string]string) {
	// Read file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		Error("Could not read from "+path, err, true)
	}

	// Replace variables
	content = []byte(ReplaceVarsInString(string(content), varMap))

	// Write content back
	err = ioutil.WriteFile(path, content, 0777)
	if err != nil {
		Error("Could not write to "+path, err, true)
	}
}

// ReplaceVarsInString replaces all variables in the stated string with the contents of the map
func ReplaceVarsInString(content string, varMap map[string]string) string {
	for key, value := range varMap {
		content = strings.ReplaceAll(content, "${{"+key+"}}", value)
	}
	return content
}

// GenerateSecrets generates random strings as secrets and replaces them in the stated file
func GenerateSecrets(path string, secrets []model.Secret) map[string]string {
	// Read file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		Error("Could not read from "+path, err, true)
	}

	// Generate a password for each occurrence of _GENERATE_PW
	newContent := string(content)
	secretsMap := make(map[string]string)
	for _, s := range secrets {
		res, err := password.Generate(s.Length, 10, 0, false, false)
		if err != nil {
			Error("Password generation failed.", err, true)
		}
		newContent = strings.ReplaceAll(newContent, "${{"+s.Variable+"}}", res)
		secretsMap[s.Name] = res
	}

	// Write content back
	err = ioutil.WriteFile(path, []byte(newContent), 0777)
	if err != nil {
		Error("Could not write to "+path, err, true)
	}

	return secretsMap
}

// DownloadFile downloads a file by its url
func DownloadFile(url string, filepath string) error {
	// Ignore untrusted authorities
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Download file
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// CommandExists checks if the stated command exists on the host system
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// IsPrivileged checks if the user has root priviledges
func IsPrivileged() bool {
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	// 0 = root, 501 = non-root user
	i, err := strconv.Atoi(string(output[:len(output)-1]))

	if err != nil {
		panic(err)
	}
	return i == 0
}

// RemoveStringFromSlice searches a string in a slice and removes it
func RemoveStringFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// SliceContainsString checks if a slice contains a certain element
func SliceContainsString(slice []string, i string) bool {
	for _, ele := range slice {
		if ele == i {
			return true
		}
	}
	return false
}

// SliceContainsInt checks if a slice contains a certain element
func SliceContainsInt(slice []int, i int) bool {
	for _, ele := range slice {
		if ele == i {
			return true
		}
	}
	return false
}

// AppendStringToSliceIfMissing checks if a slice contains a string and adds it if its not existing already
func AppendStringToSliceIfMissing(slice []string, i string) []string {
	if SliceContainsString(slice, i) {
		return slice
	}
	return append(slice, i)
}

// DockerComposeUp executes 'docker compose up' in the current directory
func DockerComposeUp(detached bool) {
	Pel()
	Pl("Running docker-compose ... ")
	Pel()

	cmd := exec.Command("docker-compose", "up")
	if detached {
		cmd = exec.Command("docker-compose", "up", "-d")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// ExecuteAndWait executes a command and wait till the execution is complete
func ExecuteAndWait(c ...string) {
	cmd := exec.Command(c[0], c[1:]...)
	cmd.Start()
	cmd.Wait()
}

// ExecuteAndWaitWithOutput executes a command and return the command output as string
func ExecuteAndWaitWithOutput(c ...string) string {
	cmd := exec.Command(c[0], c[1:]...)
	output, _ := cmd.CombinedOutput()
	return strings.TrimRight(string(output), "\r\n")
}

// ExecuteOnLinux runs a command in an isolated Linux environment
func ExecuteOnLinux(c string) {
	ensureToolbox()
	// Start docker container
	absolutePath, _ := os.Getwd()
	ExecuteAndWait("docker", "run", "-i", "-v", absolutePath+":/toolbox", "compose-generator-toolbox", c)
}

func ensureToolbox() {
	// Check if toolbox image exists locally
	imageInspect := ExecuteAndWaitWithOutput("docker", "image", "inspect", "compose-generator-toolbox")
	if !strings.HasPrefix(imageInspect, "[]") { // Image exists locally -> return
		return
	}
	if FileExists(filepath.Join(getToolboxPath(), "toolbox.img")) { // Image exists as file
		// Load iamge
		ExecuteAndWait("docker", "load", "-i", filepath.Join(getToolboxPath(), "toolbox.img"))
	} else { // Image has to be built
		// Build docker image
		ExecuteAndWait("docker", "build", "-t", "compose-generator-toolbox", getToolboxPath())
		// Save docker image as file
		ExecuteAndWait("docker", "save", "-o", filepath.Join(getToolboxPath(), "toolbox.img"), "compose-generator-toolbox")
	}
}

// ClearScreen errases the console contents
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
