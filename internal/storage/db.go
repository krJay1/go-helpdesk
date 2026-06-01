package storage

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error

	DB, err = sql.Open("postgres", "host=localhost dbname=pggo connect_timeout=5")
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

}
