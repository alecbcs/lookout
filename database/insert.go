package database

import (
	"database/sql"
	"log"
	"strings"

	"github.com/alecbcs/lookout/results"
)

// Add checks if an entry is already in the database and
// if found updates the entry, else it adds the entry.
func Add(db *sql.DB, entry *results.Entry) {
	result, found := Get(db, entry.ID)
	if found == nil {
		results.Patch(result, entry)
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		Update(tx, result)
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		Insert(db, entry)
	}
}

// Insert adds a new entry into the database.
func Insert(db *sql.DB, entry *results.Entry) {
	// Ping to check that database connection still exists.
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
			"upToDate" +
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
		entry.UpToDate)

	if err != nil {
		log.Fatal(err)
	}
}

// Update patches an existing db entry with new data.
func Update(tx *sql.Tx, entry *results.Entry) {
	stmt, err := tx.Prepare(
		"UPDATE appdata SET " +
			"latestURL = ?," +
			"latestVERSION = ?," +
			"currentURL = ?," +
			"currentVERSION = ?," +
			"upToDate = ?" +
			"WHERE id = ?;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		entry.LatestURL,
		strings.Join(entry.LatestVersion, "."),
		entry.CurrentURL,
		strings.Join(entry.CurrentVersion, "."),
		entry.UpToDate,
		entry.ID)

	if err != nil {
		log.Fatal(err)
	}
}
