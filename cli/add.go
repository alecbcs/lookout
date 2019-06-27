package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/DataDrake/cuppa/version"
	"github.com/alecbcs/lookout/database"
	"github.com/alecbcs/lookout/results"
	"github.com/alecbcs/lookout/update"
)

var Add = cmd.CMD {
	Name:  "add",
	Alias: "a",
	Short: "Add an entry to the database",
	Args:  &AddArgs{},
	Run:   AddRun,
}

type AddArgs struct {
	ID      string
	Version string
	URL     string
}

func AddRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*AddArgs)
	db := database.Open("./test.db")
	result, found := update.CheckUpdate(args.URL)
	if !found {
		log.Fatal("Unable to find " + args.ID + " " + args.URL)
	}
	entry := results.New(
		args.ID,
		result.Location,
		result.Version,
		args.URL,
		version.NewVersion(args.Version),
		args.Version == strings.Join(result.Version, "."))
	database.Add(db, entry)
	fmt.Println(entry.ID, entry.LatestURL, entry.UpToDate)
}
