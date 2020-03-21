package account

import (
	"errors"
	pg "instagram-worker/internal/pkg/postgres"
	"time"

	"github.com/lib/pq"
)

type Account struct {
	ID           int       `db:"account_id"`
	Username     string    `db:"username"`
	LatestCursor string    `db:"latest_cursor"`
	Shortcodes   []string  `db:"shortcodes"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type AccountNotFoundError struct{}

func (e *AccountNotFoundError) Error() string {
	return "Account not found"
}

func FindAccountByUsername(username string) (*Account, error) {
	db := pg.DB()
	rows, err := db.Query("SELECT * FROM account WHERE username = $1", username)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var account Account
		err := rows.Scan(&account.ID, &account.Username, &account.LatestCursor, pq.Array(&account.Shortcodes), &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, err
		}

		return &account, nil
	}

	return nil, &AccountNotFoundError{}
}

func UpdateAccount(ID int, latestCursor string, shortcodes []string) error {
	db := pg.DB()
	res, err := db.Exec("UPDATE account SET latest_cursor = $1, shortcodes = $2, updated_at = CURRENT_TIMESTAMP WHERE account_id = $3", latestCursor, pq.Array(shortcodes), ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Update account unsuccessfully")
	}

	return nil
}

func InsertAccount(username string) error {
	db := pg.DB()
	res, err := db.Exec("INSERT INTO account(username, latest_cursor, shortcodes) values ($1, '', array[]::varchar[])", username)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Insert account unsuccessfully")
	}

	return nil
}
