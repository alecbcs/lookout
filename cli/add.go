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

// Add imports a new app entry into the database.
var Add = cmd.CMD{
	Name:  "add",
	Alias: "a",
	Short: "Add an entry to the database",
	Args:  &AddArgs{},
	Run:   AddRun,
}

// AddArgs handles the specifc arguments for the add command.
type AddArgs struct {
	ID      string
	Version string
	URL     string
}

// AddRun imports a new app entry into the database.
// Will also check if the application is up-to-date before entering into the database.
func AddRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*AddArgs)
	db := database.Open("./test.db")
	result, found := update.CheckUpdate(args.URL)
	if !found {
		log.Fatal("Unable to find " + args.ID + " " + args.URL)
	}
	// Create new entry struct
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
