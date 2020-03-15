package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"instagram-worker/internal/pkg/config"
	"instagram-worker/internal/pkg/httpclient"
	"net/http"
	"regexp"
)

var cfg = config.Load()
var header = http.Header{
	"user-agent": {cfg.HTTP.Header.UA},
}

// PageInfo the info return from the user page
type PageInfo struct {
	SharedData SharedData
	QueryHash  string
	Cookies    []*http.Cookie
}

// Request request the user page
func Request(username string) (*PageInfo, error) {
	req := httpclient.Request{
		Scheme: cfg.HTTP.Scheme,
		Host:   cfg.HTTP.Host,
		Path:   fmt.Sprintf("/%s", username),
		Header: header,
	}

	resp, err := req.Get()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to request the user's page with http status %d", resp.StatusCode)
	}

	sharedData, err := parseSharedData(resp.Body)
	if err != nil {
		return nil, err
	}

	queryHash, err := parseQueryHash(resp.Body, resp.Cookies)
	if err != nil {
		return nil, err
	}

	pageInfo := PageInfo{
		SharedData: *sharedData,
		QueryHash:  queryHash,
		Cookies:    resp.Cookies,
	}

	return &pageInfo, nil
}

func parseSharedData(content string) (*SharedData, error) {
	re := regexp.MustCompile(cfg.HTML.SharedDataRegex)
	matchs := re.FindStringSubmatch(content)

	if len(matchs) < 2 {
		return nil, errors.New("Failed to parse user's shared data by html")
	}

	var sharedData SharedData
	err := json.Unmarshal([]byte(matchs[1]), &sharedData)

	if err != nil {
		return nil, errors.New("Failed to parse user's shared data by json")
	}

	return &sharedData, nil
}

func parseQueryHash(content string, cookies []*http.Cookie) (string, error) {
	profileRe := regexp.MustCompile(`/static/bundles/es6/ProfilePageContainer.js/(.+?).js`)
	profileMatchs := profileRe.FindStringSubmatch(content)
	if len(profileMatchs) < 2 {
		return "", errors.New("Failed to parse user's query ID by html")
	}

	jsLinkHash := profileMatchs[1]

	req := httpclient.Request{
		Scheme:  cfg.HTTP.Scheme,
		Host:    cfg.HTTP.Host,
		Path:    fmt.Sprintf("/static/bundles/es6/ProfilePageContainer.js/%s.js", jsLinkHash),
		Header:  header,
		Cookies: cookies,
	}

	resp, err := req.Get()
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to request the query ID js file with http status %d", resp.StatusCode)
	}

	queryidRe := regexp.MustCompile(`queryId:"(.+?)"`)
	queryidMatchs := queryidRe.FindAllStringSubmatch(resp.Body, -1)

	if len(queryidMatchs) < 3 || len(queryidMatchs[2]) < 2 {
		return "", errors.New("Failed to get the query ID")
	}

	queryHash := queryidMatchs[2][1]

	return queryHash, nil
}
