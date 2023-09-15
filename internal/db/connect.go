package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	*Queries
}

func Connect(file string) (*Connection, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	return &Connection{New(db)}, nil
}
