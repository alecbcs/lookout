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

// Search gets an entry from the database and displays the relevant information.
var Search = cmd.CMD{
	Name:  "search",
	Alias: "s",
	Short: "Search for a package in database",
	Args:  &SearchArgs{},
	Run:   SearchRun,
}

// SearchArgs handles the search specific arguments passed to the command.
type SearchArgs struct {
	ID string
}

// SearchRun searched the database and displays the information.
func SearchRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*SearchArgs)
	db := database.Open(config.Global.Database.Path)
	result, err := database.Get(db, args.ID)
	if err != nil {
		log.Fatal("Unable to locate: " + args.ID)
	}
	red := color.FgRed.Render("%-20s: %s\n")
	green := color.FgGreen.Render("%-20s: %s\n")
	white := "%-20s: %s\n"

	id := green
	latestURL := white
	latestVersion := white
	currentURL := white
	currentVersion := white

	if !result.UpToDate {
		id = red
		latestURL = green
		latestVersion = green
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
		fmt.Printf("%-20s: [%s]\n", "Dependencies", strings.Join(result.Dependencies, ", "))
	}
}
