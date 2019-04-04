package main

import (
	"github.com/alecbcs/lookout/update"
)

func main() {
	update.CheckUpdate("https://github.com/DataDrake/cuppa/archive/v1.0.4.tar.gz")
}
