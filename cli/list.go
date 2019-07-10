package cli

import (
	"fmt"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
	"github.com/alecbcs/lookout/results"
)

// List shows all the apps in the database.
var List = cmd.CMD{
	Name:  "list",
	Alias: "l",
	Short: "List all of the applications in the database.",
	Args:  &ListArgs{},
	Run:   ListRun,
}

// ListArgs handles the specific arguments passed to list.
// For real though its a place holder.
type ListArgs struct {
}

// ListRun searches the database and lists all apps in it.
func ListRun(r *cmd.RootCMD, c *cmd.CMD) {
	db := database.Open(config.Global.Database.Path)
	apps := make(chan *results.Entry, 3)
	go database.GetAll(db, apps)
	for app := range apps {
		fmt.Println(app.ID)
	}
}
