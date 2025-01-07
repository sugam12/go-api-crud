package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func NewDBConnect() (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=GoLang sslmode=disable password=admin host=localhost")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
