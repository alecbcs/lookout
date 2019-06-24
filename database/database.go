package database

import (
	"database/sql"
	"log"
	"strings"

	"github.com/alecbcs/lookout/results"

	_ "github.com/mattn/go-sqlite3" // Import sqlite3 driver for database interaction.
)

// Open opens a database and creates one if not found.
func Open(databaseName string) (db *sql.DB) {
	var err error
	db, err = sql.Open("sqlite3", databaseName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS appdata(" +
			"id TEXT PRIMARY KEY," +
			"latestURL TEXT," +
			"latestVERSION TEXT," +
			"currentURL TEXT," +
			"currentVERSION TEXT," +
			"outOFdate BOOL" +
			");")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Add inserts a new entry into the database.
func Add(db *sql.DB, entry *results.Entry) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(
		"INSERT INTO appdata(" +
			"id," +
			"latestURL," +
			"latestVERSION," +
			"currentURL," +
			"currentVERSION," +
			"outOfdate" +
			") VALUES(?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		entry.ID,
		entry.LatestURL,
		strings.Join(entry.LatestVersion, "."),
		entry.CurrentURL,
		strings.Join(entry.CurrentVersion, "."),
		entry.OutOFdate)

	if err != nil {
		log.Fatal(err)
	}
}

func Get(db *sql.DB, id string) (*results.Entry, string) {
	var (
		latestURL      string
		latestVersion  string
		currentURL     string
		currentVersion string
		outOfdate      bool
		result         *results.Entry
	)

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.Query("SELECT * FROM appdata WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if !row.Next() {
		return result, "Could not find: " + id
	}
	err = row.Scan(&id, &latestURL, &latestVersion, &currentURL, &currentVersion, &outOfdate)
	if err != nil {
		log.Fatal(err)
	}
	result = results.NewEntry(id, latestURL, latestVersion, currentURL, currentVersion, outOfdate)
	return result, ""
}
