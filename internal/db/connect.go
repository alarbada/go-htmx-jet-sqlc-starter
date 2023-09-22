package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *Queries

func Connect(user, password, host, port, name string) error {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, name))
	if err != nil {
		return err
	}

	Conn = New(db)
	return nil
}
