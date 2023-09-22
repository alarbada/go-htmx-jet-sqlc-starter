package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *Queries

func Connect(user, password, host, port, name, sslmode string) error {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user, password, host, port, name, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	Conn = New(db)
	return nil
}
