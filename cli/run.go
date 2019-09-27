package cli

import (
	"runtime"
	"sync"

	"github.com/alecbcs/lookout/ui"

	"github.com/DataDrake/cli-ng/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
	"github.com/alecbcs/lookout/results"
	"github.com/alecbcs/lookout/update"
)

// Run executes a full database scan to see if any applictions are out of date.
var Run = cmd.CMD{
	Name:  "run",
	Alias: "r",
	Short: "Run a full scan of the database to check for updates.",
	Args:  &RunArgs{},
	Run:   RunFull,
}

// RunArgs handles arguments passed to the run command.
type RunArgs struct {
}

// RunFull creates all nessisary go routines for the scan to run.
func RunFull(r *cmd.RootCMD, c *cmd.CMD) {
	db := database.Open(config.Global.Database.Path)
	// Create the input chan to store database enties.
	input := make(chan *results.Entry)
	// Create the toUpdate chan to store updated apps to write back to database.
	// Buffer chan to improve performance.
	toUpdate := make(chan *results.Entry, 32)
	workers := genNumWorkers()
	// Create Sync WaitGroup to watch goroutines so that we don't close the channels early.
	var wg sync.WaitGroup
	wg.Add(workers)
	// Get all will read db entries and put in queue for workers.
	go database.GetAll(db, input)
	for i := 0; i < workers; i++ {
		go runWorker(&wg, input, toUpdate)
	}
	// Create a go routine to wait till all workers are finished before closing channel.
	go func() {
		wg.Wait()
		close(toUpdate)
	}()
	// Update all db entires that are out-of-date.
	for app := range toUpdate {
		database.Update(db, app)
	}
	defer db.Close()
}

// Generate the number of worker processes to optimize efficiency.
// Subtract 2 from the number of cores because of the main thread and the GetAll function.
func genNumWorkers() int {
	if runtime.NumCPU() > 2 {
		return runtime.NumCPU() - 2
	}
	return 1
}

// runWorker defines an update worker process.
func runWorker(wg *sync.WaitGroup, input <-chan *results.Entry, output chan<- *results.Entry) {
	// Pull the next app off the queue channel.
	for app := range input {
		result, found := update.CheckUpdate(app.CurrentURL)
		if !found {
			ui.PrintRed(app.ID, "Not Found")
			continue
		}
		// If the latest version does not match the database version, mark out-od-date.
		if !update.UpToDate(result.Version, app.CurrentVersion) {
			app.LatestURL = result.Location
			app.LatestVersion = result.Version
			output <- app
			ui.PrintRed(app.ID, "New Version Found")
		} else {
			ui.PrintCyan(app.ID, "Up-To-Date")
		}
	}
	// Tell WaitGroup that go routine has finished.
	defer wg.Done()

}
