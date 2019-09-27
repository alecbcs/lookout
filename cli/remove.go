package cli

import (
	"log"

	"github.com/alecbcs/lookout/ui"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
)

// Remove deletes an entry from the application database.
var Remove = cmd.CMD{
	Name:  "remove",
	Alias: "rm",
	Short: "Remove an entry from the database.",
	Args:  &RemoveArgs{},
	Run:   RemoveRun,
}

// RemoveArgs handles the arguments passed to the remove command.
type RemoveArgs struct {
	ID string
}

//RemoveRun executes a remove statments from the SQL database.
func RemoveRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*RemoveArgs)
	db := database.Open(config.Global.Database.Path)
	defer db.Close()
	err := database.Delete(db, args.ID)
	if err != nil {
		log.Fatal(err)
	}
	ui.PrintRed(args.ID, "REMOVED")
}
