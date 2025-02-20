package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/sugam12/go-api-crud/cmd/api"
	"github.com/sugam12/go-api-crud/config"
	"github.com/sugam12/go-api-crud/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.EnvVars.DBUserName,
		Passwd:               config.EnvVars.DBPassword,
		Addr:                 config.EnvVars.DBAddress,
		DBName:               config.EnvVars.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	apiServer := api.NewAPIServer("localhost:8080", nil)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("sucessfully connected to db")
}
