package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"errors"
	"instagram-worker/internal/pkg/config"
	"instagram-worker/internal/pkg/httpclient"
)

var cfg = config.Load()
var header = http.Header{
	"user-agent": {cfg.HTTP.Header.UA},
}

type PageInfo struct {
	SharedData SharedData
}

func Request(shortcode string) (*PageInfo, error) {
	req := httpclient.Request{
		Scheme: cfg.HTTP.Scheme,
		Host:   cfg.HTTP.Host,
		Path:   fmt.Sprintf("/p/%s", shortcode),
		Header: header,
	}

	resp, err := req.Get()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to request the post %s with http status %d", shortcode, resp.StatusCode)
	}

	sharedData, err := parseSharedData(resp.Body)
	if err != nil {
		return nil, err
	}

	pageInfo := PageInfo{
		SharedData: *sharedData,
	}

	return &pageInfo, nil
}

func parseSharedData(content string) (*SharedData, error) {
	re := regexp.MustCompile(cfg.HTML.SharedDataRegex)
	matchs := re.FindStringSubmatch(content)

	if len(matchs) < 2 {
		return nil, errors.New("Failed to parse post's' shared data by html")
	}

	var sharedData SharedData
	err := json.Unmarshal([]byte(matchs[1]), &sharedData)

	if err != nil {
		return nil, errors.New("Failed to parse post's' shared data by json")
	}

	return &sharedData, nil
}
