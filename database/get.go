package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/alecbcs/cuppa/version"
	"github.com/alecbcs/lookout/results"
)

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
	depMap, _ := GetDeps(db, id)
	deps := make([]string, len(depMap))
	counter := 0
	for dep := range depMap {
		deps[counter] = dep
		counter++
	}
	result = results.New(
		id,
		latestURL,
		version.NewVersion(latestVersion),
		currentURL,
		version.NewVersion(currentVersion),
		upToDate)
	result.Dependencies = deps
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

	entries := make([]*results.Entry, 0)
	for rows.Next() {
		err = rows.Scan(&id, &latestURL, &latestVersion, &currentURL, &currentVersion, &upToDate)
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, results.New(
			id,
			latestURL,
			version.NewVersion(latestVersion),
			currentURL,
			version.NewVersion(currentVersion),
			upToDate))
	}
	for _, entry := range entries {
		output <- entry
	}
	close(output)
}
