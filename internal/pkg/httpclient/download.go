package httpclient

import (
	"net/http"
	"net/url"
	"os"
	"strings"

	"fmt"
	"io"
)

func Download(rawurl string, downloadDir string) (string, error) {
	err := os.MkdirAll(downloadDir, 777)
	if err != nil {
		return "", err
	}

	downloadURL, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, downloadURL.String(), nil)
	if err != nil {
		return "", err
	}
	res, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	filepath, err := buildFilepath(downloadURL, downloadDir)
	if err != nil {
		return "", err
	}

	file, err := createFile(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func buildFilepath(downloadURL *url.URL, downloadDir string) (string, error) {
	urlPath := downloadURL.Path
	segments := strings.Split(urlPath, "/")

	filename := segments[len(segments)-1]

	filepath := fmt.Sprintf("%s/%s", downloadDir, filename)

	return filepath, nil
}

func createFile(filepath string) (*os.File, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
