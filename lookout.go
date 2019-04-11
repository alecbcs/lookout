package main

import (
	"fmt"

	"github.com/alecbcs/lookout/database"
	"github.com/alecbcs/lookout/update"
)

func main() {
	source := "https://github.com/DataDrake/cuppa/archive/v1.0.4.tar.gz"
	version, location := update.CheckUpdate(source)
	fmt.Println(version, location)
	database.Add("cuppa", location, version)
}
