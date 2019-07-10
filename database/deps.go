package database

import (
	"database/sql"
	"log"
)

// ImportDeps maps an app's dependencies to the app.
// Unlike AddDep this will also remove old dependencies.
// This was we can keep track of package relationships.
func ImportDeps(db *sql.DB, id string, new []string) {
	deps, _ := GetDeps(db, id)
	leftover, _ := GetDeps(db, id)
	for _, dependency := range new {
		if _, ok := deps[dependency]; !ok {
			InsertDep(db, id, dependency)
			continue
		}
		delete(leftover, dependency)
	}
	for key, value := range leftover {
		DeleteDep(db, value, key)
	}
}

// DeleteDep removes a dependency, package relationship from the database.
func DeleteDep(db *sql.DB, id string, dep string) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(
		"DELETE FROM deps WHERE id = ? AND dependency = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(id, dep)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertDep adds a dependency, package relationship to the database.
func InsertDep(db *sql.DB, id string, dep string) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(
		"INSERT INTO deps(" +
			"id," +
			"dependency" +
			") VALUES (?,?);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(id, dep)
	if err != nil {
		log.Fatal(err)
	}
}

// GetDeps returns a map of all a packages dependent packages.
func GetDeps(db *sql.DB, id string) (result map[string]string, found bool) {
	var dependency string
	result = make(map[string]string)
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM deps WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &dependency)
		if err != nil {
			log.Fatal(err)
		}
		result[dependency] = id
	}
	return result, len(result) > 0
}
