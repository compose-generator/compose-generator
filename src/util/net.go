/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"io"
	"net/http"
	"os"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// DownloadFile downloads a file by its url
func DownloadFile(url string, filepath string) error {
	// Ignore untrusted authorities
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	// Download file
	resp, err := client.Get(url)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			ErrorLogger.Println("Error closing downloaded file: " + err.Error())
			logError("Error closing downloaded file", true)
		}
	}()
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(filepath)
	defer func() {
		if err := out.Close(); err != nil {
			ErrorLogger.Println("Error closing downloaded file: " + err.Error())
			logError("Error closing downloaded file", true)
		}
	}()
	if err != nil {
		return err
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
