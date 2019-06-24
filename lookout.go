package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/alecbcs/lookout/database"
	"github.com/alecbcs/lookout/results"
	"github.com/alecbcs/lookout/update"
)

func main() {
	db := database.Open("./test.db")
	source := "https://github.com/DataDrake/cuppa/archive/v1.0.4.tar.gz"
	result, found := update.CheckUpdate(source)
	if !found {
		fmt.Println("Not Found")
	}
	entry := results.NewEntry("cuppa", result.Location, strings.Join(result.Version, "."), result.Location, strings.Join(result.Version, "."), true)
	fmt.Println(entry.ID, entry.LatestURL)
	//database.Add(db, entry)
	returned, error := database.Get(db, "cuppa")
	if error != "" {
		log.Fatal(error)
	}
	fmt.Println(returned.CurrentVersion)
}
