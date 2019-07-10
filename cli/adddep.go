package cli

import (
	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
	"github.com/alecbcs/lookout/ui"
)

// AddDep imports a new app, dependency relationship into the database.
var AddDep = cmd.CMD{
	Name:  "add-dependency",
	Alias: "ad",
	Short: "Add an entry, dependency relationship to the database",
	Args:  &AddDepArgs{},
	Run:   AddDepRun,
}

// AddDepArgs handles the specific arguments for the AddDep command.
type AddDepArgs struct {
	ID  string
	Dep string
}

// AddDepRun imports a new app, dependency relationship.
// It will also check for duplicate dependencies before continuing.
func AddDepRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*AddDepArgs)
	db := database.Open(config.Global.Database.Path)
	deps, _ := database.GetDeps(db, args.ID)
	if _, ok := deps[args.Dep]; !ok {
		database.InsertDep(db, args.ID, args.Dep)
	}
	ui.PrintCyan(args.ID, "DEPENDENCY ["+args.Dep+"] ADDED")

}
