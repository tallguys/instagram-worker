package crawler

import (
	"fmt"
	postBusiness "instagram-worker/internal/app/model/business/post"
	postPersist "instagram-worker/internal/app/model/persistence/post"
	"instagram-worker/internal/pkg/httpclient"
	"net/url"
	"strings"

	logger "github.com/sirupsen/logrus"
)

const IMAGE_TYPE = "image"
const VIDEO_TYPE = "video"

func buildFilepath(downloadURL *url.URL, username string, shortcode string, likeCount int) string {
	urlPath := downloadURL.Path
	segments := strings.Split(urlPath, "/")

	filename := segments[len(segments)-1]

	filepath := fmt.Sprintf("%s/%s/%d_%s_%s", cfg.Download.Folder, username, likeCount, shortcode, filename)

	return filepath
}

func download(rawurl string, username string, shortcode string, likeCount int) (string, error) {
	downloadURL, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	filepath := buildFilepath(downloadURL, username, shortcode, likeCount)

	err = httpclient.Download(downloadURL, filepath)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func process(username string, shortcode string, accountID int) error {
	_posts, err := postPersist.FindPostsByShortcode(accountID, shortcode)
	if err != nil {
		return err
	}

	if len(_posts) > 0 {
		logger.Info(fmt.Sprintf("%s had processed before, skip.", shortcode))
		return nil
	}

	pageInfo, err := postBusiness.Request(shortcode)
	if err != nil {
		return err
	}

	postPage := pageInfo.SharedData.EntryData.PostPage[0]

	var content string
	if len(postPage.Graphql.ShortcodeMedia.EdgeMediaToCaption.Edges) > 0 {
		content = postPage.Graphql.ShortcodeMedia.EdgeMediaToCaption.Edges[0].Node.Text
	} else {
		content = ""
	}
	takenAt := postPage.Graphql.ShortcodeMedia.TakenAtTimestamp
	likeCount := postPage.Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count
	typename := postPage.Graphql.ShortcodeMedia.Typename

	posts := []*postPersist.Post{}

	switch typename {
	case "GraphSidecar":
		for _, edge := range postPage.Graphql.ShortcodeMedia.EdgeSidecarToChildren.Edges {
			var url string
			var mediaType string
			if edge.Node.IsVideo {
				url = edge.Node.VideoURL
				mediaType = VIDEO_TYPE
			} else {
				url = edge.Node.DisplayURL
				mediaType = IMAGE_TYPE
			}

			downloadPath, err := download(url, username, shortcode, likeCount)
			if err != nil {
				return err
			}

			post := postPersist.Post{
				AccountID:    accountID,
				Shortcode:    shortcode,
				Content:      content,
				LikeCount:    likeCount,
				MediaType:    mediaType,
				TakenAt:      takenAt,
				MediaURL:     url,
				DownloadPath: downloadPath,
			}

			posts = append(posts, &post)
		}
		break
	case "GraphImage":
		url := postPage.Graphql.ShortcodeMedia.DisplayURL
		downloadPath, err := download(url, username, shortcode, likeCount)
		if err != nil {
			return err
		}

		post := postPersist.Post{
			AccountID:    accountID,
			Shortcode:    shortcode,
			Content:      content,
			LikeCount:    likeCount,
			MediaType:    IMAGE_TYPE,
			TakenAt:      takenAt,
			MediaURL:     url,
			DownloadPath: downloadPath,
		}

		posts = append(posts, &post)
		break
	case "GraphVideo":
		url := postPage.Graphql.ShortcodeMedia.VideoURL
		downloadPath, err := download(url, username, shortcode, likeCount)
		if err != nil {
			return err
		}

		post := postPersist.Post{
			AccountID:    accountID,
			Shortcode:    shortcode,
			Content:      content,
			LikeCount:    likeCount,
			MediaType:    VIDEO_TYPE,
			TakenAt:      takenAt,
			MediaURL:     url,
			DownloadPath: downloadPath,
		}

		posts = append(posts, &post)
		break
	default:
		return fmt.Errorf("Unknown typename: %s", typename)
	}

	err = postPersist.InsertPosts(posts)
	if err != nil {
		return err
	}

	return nil
}
