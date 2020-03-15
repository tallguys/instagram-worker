package postgres

import (
	"database/sql"
	"fmt"
	"instagram-worker/internal/pkg/config"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once
var db *sql.DB

var cfg = config.Load()

func DB() *sql.DB {
	once.Do(func() {
		connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.POSTGRES.User, cfg.POSTGRES.PWD, cfg.POSTGRES.DB)
		_db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		db = _db
	})

	return db
}
