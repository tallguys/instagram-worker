package account

import (
	pg "instagram-worker/internal/pkg/postgres"
	"time"
)

type Post struct {
	ID           int       `db:"post_id"`
	AccountID    int       `db:"account_id"`
	Shortcode    string    `db:"shortcode"`
	Content      string    `db:"content"`
	LikeCount    int       `db:"like_count"`
	MediaType    string    `db:"media_type"`
	TakenAt      int       `db:"taken_at"`
	MediaURL     string    `db:"media_url"`
	DownloadPath string    `db:"download_path"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func FindPostsByShortcode(accountID int, shortcode string) ([]*Post, error) {
	db := pg.DB()
	rows, err := db.Query("SELECT * FROM post WHERE account_id = $1 and shortcode = $2", accountID, shortcode)
	if err != nil {
		return nil, err
	}

	res := []*Post{}

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.AccountID, &post.Shortcode, &post.Content, &post.LikeCount, &post.MediaType, &post.TakenAt, &post.MediaURL, &post.DownloadPath, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &post)
	}

	return res, nil
}

func InsertPost(post *Post) error {
	db := pg.DB()
	_, err := db.Exec(
		"INSERT INTO post(account_id, shortcode, content, like_count, media_type, taken_at, media_url, download_path) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		post.AccountID,
		post.Shortcode,
		post.Content,
		post.LikeCount,
		post.MediaType,
		post.TakenAt,
		post.MediaURL,
		post.DownloadPath,
	)

	if err != nil {
		return err
	}

	return nil
}

func InsertPosts(posts []*Post) error {
	tx, err := pg.DB().Begin()
	if err != nil {
		return err
	}

	for _, post := range posts {
		err := InsertPost(post)
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				return rerr
			}

			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
