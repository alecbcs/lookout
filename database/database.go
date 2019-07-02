package database

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/DataDrake/cuppa/version"
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
	// Create the appdata table if is doesn't already exist.
	// This will also create the database if it doesn't exist.
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS appdata(" +
			"id TEXT PRIMARY KEY," +
			"latestURL TEXT," +
			"latestVERSION TEXT," +
			"currentURL TEXT," +
			"currentVERSION TEXT," +
			"upToDate BOOL" +
			");")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Add checks if an entry is already in the database and
// if found updates the entry, else it adds the entry.
func Add(db *sql.DB, entry *results.Entry) {
	result, found := Get(db, entry.ID)
	if found == nil {
		results.Patch(result, entry)
		Update(db, result)
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
func Update(db *sql.DB, entry *results.Entry) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(
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

// Get finds and returns an entry from the database.
func Get(db *sql.DB, id string) (*results.Entry, error) {
	var (
		latestURL      string
		latestVersion  string
		currentURL     string
		currentVersion string
		upToDate       bool
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
		return result, errors.New("Could not find: " + id)
	}
	err = row.Scan(
		&id,
		&latestURL,
		&latestVersion,
		&currentURL,
		&currentVersion,
		&upToDate)
	if err != nil {
		log.Fatal(err)
	}
	result = results.New(
		id,
		latestURL,
		version.NewVersion(latestVersion),
		currentURL,
		version.NewVersion(currentVersion),
		upToDate)
	return result, nil
}

// GetAll opens a channel and reads each entry into the channel.
func GetAll(db *sql.DB, output chan *results.Entry) {
	var (
		id             string
		latestURL      string
		latestVersion  string
		currentURL     string
		currentVersion string
		upToDate       bool
	)

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM appdata")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &latestURL, &latestVersion, &currentURL, &currentVersion, &upToDate)
		if err != nil {
			log.Fatal(err)
		}
		output <- results.New(
			id,
			latestURL,
			version.NewVersion(latestVersion),
			currentURL,
			version.NewVersion(currentVersion),
			upToDate)
	}
	close(output)
}
