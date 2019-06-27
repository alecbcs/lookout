package cli

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/database"
)

var Search = cmd.CMD{
	Name:  "search",
	Alias: "s",
	Short: "Search for a package in database",
	Args:  &SearchArgs{},
	Run:   SearchRun,
}

type SearchArgs struct {
	ID string
}

func SearchRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*SearchArgs)
	db := database.Open("./test.db")
	result, err := database.Get(db, args.ID)
	if err != nil {
		log.Fatal("Unable to locate: " + args.ID)
	}
	fmt.Println("ID:             " + result.ID)
	fmt.Println("LatestURL:      " + result.LatestURL)
	fmt.Println("LatestVersion:  " + strings.Join(result.LatestVersion, "."))
	fmt.Println("CurrentURL:     " + result.CurrentURL)
	fmt.Println("CurrentVersion: " + strings.Join(result.CurrentVersion, "."))
	fmt.Println("Up-To-Date:     " + strconv.FormatBool(result.UpToDate))
}
