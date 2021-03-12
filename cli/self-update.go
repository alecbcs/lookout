package cli

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/alecbcs/lookout/config"
	"github.com/inconshreveable/go-update"
	"github.com/tcnksm/go-latest"
)

// SelfUpdate checks for a new version of Lookout and updates itself
// if a newer version is found and the user agrees to update.
var SelfUpdate = cmd.Sub{
	Name:  "self-update",
	Alias: "ups",
	Short: "Have lookout update itself to the lastest version.",
	Args:  &SelfUpdateArgs{},
	Flags: &SelfUpdateFlags{},
	Run:   UpdateRun,
}

// SelfUpdateArgs handles the specific arguments for the update command.
type SelfUpdateArgs struct {
}

// SelfUpdateFlags handles the specific flags for the update command.
type SelfUpdateFlags struct {
	Yes bool `short:"y" long:"yes" desc:"If a newer version is found update without prompting the user."`
}

// UpdateRun handles the checking and self updating of Lookout.
func UpdateRun(r *cmd.Root, c *cmd.Sub) {
	fmt.Printf("Current Version: %s\n", config.Global.General.Version)

	flags := c.Flags.(*SelfUpdateFlags)
	latestVersion := &latest.GithubTag{
		Owner:      "alecbcs",
		Repository: "lookout",
	}

	res, _ := latest.Check(latestVersion, config.Global.General.Version)
	fmt.Printf("Latest Version: %s\n", res.Current)

	if res.Outdated {
		if !flags.Yes {
			fmt.Println("Would you like to update Lookout to the newest version? ([y]/n)")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.ToLower(strings.TrimSpace(input))
			if input == "n" {
				return
			}
		}
		url := "https://github.com/alecbcs/lookout/releases/download/v" + res.Current + "/lookout-v" + res.Current + "-" + runtime.GOOS + "-" + runtime.GOARCH

		doneChan := make(chan int, 1)
		wg := sync.WaitGroup{}
		wg.Add(1)

		// Display Spinner on Update.
		go SpinnerWait(doneChan, "Updating Lookout...", &wg)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		err = update.Apply(resp.Body, update.Options{})
		if err != nil {
			log.Fatal(err)
		}

		doneChan <- 0
		wg.Wait()

		fmt.Print("\rUpdating Lookout: Done!\n")
	} else {
		fmt.Println("Already Up-To-Date!")
	}
}

// Spinner is an array of the progression of the spinner.
var Spinner = []string{"|", "/", "-", "\\"}

// SpinnerWait displays the actual spinner
func SpinnerWait(done chan int, message string, wg *sync.WaitGroup) {
	ticker := time.Tick(time.Millisecond * 128)
	frameCounter := 0
	for {
		select {
		case _ = <-done:
			wg.Done()
			return
		default:
			<-ticker
			ind := frameCounter % len(Spinner)
			fmt.Printf("\r[%v] "+message, Spinner[ind])
			frameCounter++
		}
	}
}
