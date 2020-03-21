package httpclient

import (
	"net/http"
	"net/url"
	"os"

	"io"
)

func Download(downloadURL *url.URL, filepath string) error {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, downloadURL.String(), nil)
	if err != nil {
		return err
	}
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := createFile(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func createFile(filepath string) (*os.File, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
