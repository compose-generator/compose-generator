package util

import (
	"io/ioutil"
	"os"
	"strings"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir checks if a file is a directory
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// AddFileToGitignore takes a path and adds it to the .gitignore file in the current dir
func AddFileToGitignore(path string) {
	filename := ".gitignore"
	var f *os.File
	content := ""
	if FileExists(filename) {
		// File does exist already
		b, err1 := ioutil.ReadFile(filename)
		if err1 != nil {
			Error("Could not read "+filename+" file", err1, true)
		}
		content = string(b) + "\n"
		if strings.Contains(content, path) {
			// This path is already included
			return
		}
		f, _ = os.OpenFile(filename, os.O_WRONLY, 0755)
	} else {
		// File does not exist yet
		f, _ = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	}

	defer f.Close()
	_, err2 := f.WriteString(content + "# Docker secrets\n" + path)
	if err2 != nil {
		Error("Could not write to "+filename+" file", err2, true)
	}
}
