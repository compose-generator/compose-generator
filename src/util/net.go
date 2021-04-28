package util

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

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
