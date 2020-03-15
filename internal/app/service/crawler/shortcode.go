package crawler

import (
	"encoding/json"
	"fmt"
	accountBusiness "instagram-worker/internal/app/model/business/account"
	accountPersist "instagram-worker/internal/app/model/persistence/account"
	"instagram-worker/internal/pkg/config"
	"instagram-worker/internal/pkg/httpclient"
	"net/http"
	"net/url"
)

var cfg = config.Load()
var header = http.Header{
	"user-agent": {cfg.HTTP.Header.UA},
}

func uniqShortcodes(shortcodes []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, shortcode := range shortcodes {
		set[shortcode] = struct{}{}
	}
	return set
}

func scrollDown(pageInfo *accountBusiness.PageInfo, endCursor string) ([]string, bool, string, error) {
	profile := pageInfo.SharedData.EntryData.ProfilePage[0].Graphql.User
	variables := struct {
		ID    string `json:"id"`
		First int8   `json:"first"`
		After string `json:"after"`
	}{
		ID:    profile.ID,
		First: 50,
		After: endCursor,
	}

	variablesJson, err := json.Marshal(variables)
	if err != nil {
		return nil, false, "", err
	}

	req := httpclient.Request{
		Scheme: cfg.HTTP.Scheme,
		Host:   cfg.HTTP.Host,
		Path:   "/graphql/query/",
		Query: url.Values{
			"query_hash": {pageInfo.QueryHash},
			"variables":  {string(variablesJson)},
		},
		Header:  header,
		Cookies: pageInfo.Cookies,
	}

	resp, err := req.Get()
	if err != nil {
		return nil, false, "", err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, false, "", fmt.Errorf("Failed to request the user's page with http status %d", resp.StatusCode)
	}

	var respData struct {
		Data struct {
			User struct {
				EdgeOwnerToTimelineMedia accountBusiness.EdgeOwnerToTimelineMedia `json:"edge_owner_to_timeline_media"`
			} `json:"user"`
		} `json:"data"`
	}

	err = json.Unmarshal([]byte(resp.Body), &respData)
	if err != nil {
		return nil, false, "", err
	}

	res := []string{}

	for _, edge := range respData.Data.User.EdgeOwnerToTimelineMedia.Edges {
		res = append(res, edge.Node.Shortcode)
	}

	_hasNextPage := respData.Data.User.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
	_endCursor := respData.Data.User.EdgeOwnerToTimelineMedia.PageInfo.EndCursor

	return res, _hasNextPage, _endCursor, nil
}

func fetchShortCodes(pageInfo *accountBusiness.PageInfo, account accountPersist.Account) error {
	profile := pageInfo.SharedData.EntryData.ProfilePage[0].Graphql.User

	edges := profile.EdgeOwnerToTimelineMedia.Edges

	var endCursor string
	var hasNextPage bool

	shortcodeSet := uniqShortcodes(account.Shortcodes)
	res := account.Shortcodes

	if account.LatestCursor == "" { // first time or finish scrolling before, restart scrolling.
		scrolled := true

		for _, edge := range edges {
			shortcode := edge.Node.Shortcode

			_, ok := shortcodeSet[shortcode]
			if !ok {
				scrolled = false
				res = append(res, shortcode)
			}
		}

		if scrolled { // think the page had been scrolled if all of the shortcodes return from this turn existed
			return nil
		}

		endCursor = profile.EdgeOwnerToTimelineMedia.PageInfo.EndCursor
		hasNextPage = profile.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage

		err := accountPersist.UpdateAccount(account.ID, endCursor, res)
		if err != nil {
			return err
		}
	} else { // scroll to some way but not finish
		endCursor = account.LatestCursor
		hasNextPage = true
	}

	for hasNextPage { // loop to scroll down to fetch shortcodes
		shortcodes, _hasNextPage, _endCursor, err := scrollDown(pageInfo, endCursor)

		endCursor = _endCursor
		hasNextPage = _hasNextPage

		if err != nil {
			return err
		}

		scrolled := true
		for _, shortcode := range shortcodes {
			_, ok := shortcodeSet[shortcode]
			if !ok {
				scrolled = false
				res = append(res, shortcode)
			}
		}

		if scrolled {
			return nil
		}

		err = accountPersist.UpdateAccount(account.ID, endCursor, res)
		if err != nil {
			return err
		}

		fmt.Printf("fetched %d shortcodes\n", len(res))
	}

	err := accountPersist.UpdateAccount(account.ID, "", res)
	if err != nil {
		return err
	}

	return nil
}
