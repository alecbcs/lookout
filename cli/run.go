package cli

import (
	"fmt"
	"log"
	"sync"

	"github.com/gookit/color"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/database"
	"github.com/alecbcs/lookout/results"
	"github.com/alecbcs/lookout/update"
)

var Run = cmd.CMD{
	Name:  "run",
	Alias: "r",
	Short: "Run a full scan of the database to check for updates.",
	Args:  &RunArgs{},
	Run:   RunFull,
}

type RunArgs struct {
}

func RunFull(r *cmd.RootCMD, c *cmd.CMD) {
	db := database.Open("./test.db")
	input := make(chan *results.Entry, 32)
	toUpdate := make(chan *results.Entry, 32)
	var wg sync.WaitGroup
	wg.Add(1)
	go database.GetAll(db, input)
	go runWorker(&wg, input, toUpdate)
	go func() {
		wg.Wait()
		close(toUpdate)
	}()
	for app := range toUpdate {
		database.Update(db, app)
	}

}

func runWorker(wg *sync.WaitGroup, input <-chan *results.Entry, output chan<- *results.Entry) {
	for app := range input {
		result, found := update.CheckUpdate(app.CurrentURL)
		if !found {
			log.Println("Unable to find: " + app.ID)
			continue
		}
		red := color.FgRed.Render("%s\n")
		cyan := color.FgCyan.Render("%s\n")
		if result.Version.Compare(app.CurrentVersion) < 0 {

			app.LatestURL = result.Location
			app.LatestVersion = result.Version
			output <- app
			fmt.Printf("%-15s: "+red, app.ID, "New Version Found")
		} else {
			fmt.Printf("%-15s: "+cyan, app.ID, "Up-To-Date")
		}
	}
	defer wg.Done()

}
