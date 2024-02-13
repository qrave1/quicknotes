package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/qrave1/quicknotes/internal/config"
)

type Storage struct {
	DB *sql.DB
}

func MustLoad(c *config.config) *Storage {
	db, err := sql.Open("postgres", c.DbConfig.DSN)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return &Storage{
		DB: db,
	}
}
