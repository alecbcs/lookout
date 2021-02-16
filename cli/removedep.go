package cli

import (
	"github.com/alecbcs/lookout/ui"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
)

// RemoveDep deletes an app, dependency relationship from the database.
var RemoveDep = cmd.Sub{
	Name:  "remove-dependency",
	Alias: "rd",
	Short: "Remove an entry, dependency relationship from the database.",
	Args:  &RemoveDepArgs{},
	Run:   RemoveDepRun,
}

// RemoveDepArgs handles the specific arguments for the RemoveDep command.
type RemoveDepArgs struct {
	ID  string
	Dep string
}

// RemoveDepRun deletes an app, dependency relationship from the database.
func RemoveDepRun(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*RemoveDepArgs)
	db := database.Open(config.Global.Database.Path)
	defer db.Close()
	deps, _ := database.GetDeps(db, args.ID)
	if _, ok := deps[args.Dep]; ok {
		database.DeleteDep(db, args.ID, args.Dep)
		ui.PrintCyan(args.ID, "DEPENDENCY ["+args.Dep+"] REMOVED")
	} else {
		ui.PrintRed(args.ID, "DEPENDENCY ["+args.Dep+"] NOT FOUND")
	}
}
