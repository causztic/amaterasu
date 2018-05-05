package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgres driver
)

var db *sql.DB
var err error

// InitDB initializes the database connection
func InitDB() {
	connStr := "user=yaojie dbname=godb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
