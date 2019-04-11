package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS appdata(id TEXT, url TEXT, version TEXT);")
	if err != nil {
		log.Fatal(err)
	}
}

// Add inserts a new entry into the database.
func Add(id string, url string, version string) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO appdata(id, url, version) VALUES(?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(id, url, version)
	if err != nil {
		log.Fatal(err)
	}
}