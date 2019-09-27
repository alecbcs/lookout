package cli

import (
	"log"

	"github.com/alecbcs/lookout/ui"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
)

// Upgrade modifies an entry in the database to make it the latest version.
var Upgrade = cmd.CMD{
	Name:  "upgrade",
	Alias: "up",
	Short: "Set an entry to the latest version possible.",
	Args:  &UpgradeArgs{},
	Run:   UpgradeRun,
}

// UpgradeArgs handles the specific arguments for the upgrade command.
type UpgradeArgs struct {
	ID string
}

// UpgradeRun will patch an existing database entry with the latest information.
func UpgradeRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*UpgradeArgs)
	db := database.Open(config.Global.Database.Path)
	defer db.Close()
	entry, err := database.Get(db, args.ID)
	if err != nil {
		log.Fatal("Unable to locate: " + args.ID)
	}
	entry.CurrentURL = entry.LatestURL
	entry.CurrentVersion = entry.LatestVersion
	entry.UpToDate = true
	database.Update(db, entry)
	ui.PrintCyan(entry.ID, "UPDATED")
}
