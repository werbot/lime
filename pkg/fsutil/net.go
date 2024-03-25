package fsutil

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Download downloads file from url and saves it to filepath
func Download(filepath string, url string) (err error) {
	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check HTTP response status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errMsg := fmt.Sprintf("bad status: %s", resp.Status)
		return errors.New(errMsg)
	}

	// Open file to write
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Copy file content from response body
	_, err = io.Copy(file, resp.Body)
	return err
}
