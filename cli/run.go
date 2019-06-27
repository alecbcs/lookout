package cli

import (
	"log"
	"sync"

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
		if result.Location != app.LatestURL {
			app.LatestURL = result.Location
			app.LatestVersion = result.Version
			output <- app
		}
	}
	defer wg.Done()

}
