package cli

import (
	"log"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
)

// Info gets an entry from the database and displays the relevant information.
var Info = cmd.CMD{
	Name:  "info",
	Alias: "in",
	Short: "Displays the information for a package in the database",
	Args:  &InfoArgs{},
	Run:   InfoRun,
}

// InfoArgs handles the search specific arguments passed to the command.
type InfoArgs struct {
	ID string
}

// InfoRun searched the database and displays the information.
func InfoRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*InfoArgs)
	db := database.Open(config.Global.Database.Path)
	defer db.Close()
	result, err := database.Get(db, args.ID)
	if err != nil {
		log.Fatal("Unable to locate: " + args.ID)
	}
	result.PrintEntry()
}
