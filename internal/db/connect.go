package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *Queries

func Connect(file string) error {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return err
	}

	Conn = New(db)
	return nil
}
