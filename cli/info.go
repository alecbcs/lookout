package cli

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
	"github.com/gookit/color"
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
	red := color.FgRed.Render("%-20s: %s\n")
	cyan := color.FgCyan.Render("%-20s: %s\n")
	white := "%-20s: %s\n"

	id := cyan
	latestURL := white
	latestVersion := white
	currentURL := white
	currentVersion := white

	if !result.UpToDate {
		id = red
		latestURL = cyan
		latestVersion = cyan
		currentURL = red
		currentVersion = red
	}
	fmt.Printf(id, "Package ID", result.ID)
	fmt.Printf(latestURL, "LatestURL", result.LatestURL)
	fmt.Printf(currentURL, "CurrentURL", result.CurrentURL)
	fmt.Printf(latestVersion, "LatestVersion", strings.Join(result.LatestVersion, "."))
	fmt.Printf(currentVersion, "CurrentVersion", strings.Join(result.CurrentVersion, "."))
	fmt.Printf("%-20s: %s\n", "Up-To-Date", strconv.FormatBool(result.UpToDate))
	if len(result.Dependencies) > 0 {
		fmt.Printf("%-20s: %s\n", "Dependencies", strings.Join(result.Dependencies, ", "))
	}
}
