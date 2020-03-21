package crawler

import (
	"fmt"
	postBusiness "instagram-worker/internal/app/model/business/post"
	postPersist "instagram-worker/internal/app/model/persistence/post"
	"instagram-worker/internal/pkg/httpclient"

	logger "github.com/sirupsen/logrus"
)

const IMAGE_TYPE = "image"
const VIDEO_TYPE = "video"

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
	dir := fmt.Sprintf("%s/%s/%s", cfg.Download.Folder, username, shortcode)

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

			downloadPath, err := httpclient.Download(url, dir)
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
		downloadPath, err := httpclient.Download(url, dir)
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
		downloadPath, err := httpclient.Download(url, dir)
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
