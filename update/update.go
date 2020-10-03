package update

import (
	cuppa "github.com/DataDrake/cuppa/config"
	"github.com/DataDrake/cuppa/providers"
	"github.com/DataDrake/cuppa/results"
	"github.com/alecbcs/lookout/config"

	"log"
)

func init() {
	// Port Lookout configuration to CUPPA
	cuppa.Global.Github.Key = config.Global.Github.Key
}

// CheckUpdate checks a given URL for the latest available release.
func CheckUpdate(archive string) (*results.Result, bool) {
	// define a CUPPA result
	var r *results.Result
	found := false
	// Iterate through all available providers that CUPPA supports.
	for _, p := range providers.All() {
		name := p.Match(archive)
		if name == "" {
			continue
		}

		// GitHub will never work without a token
		if p.Name() == "GitHub" && cuppa.Global.Github.Key == "" {
			log.Fatal("A GitHub token is required in your $HOME/.config/lookout/lookout.config")
		}

		// Pull the latest (non-beta) release from repository.
		r, s := p.Latest(name)
		if s != results.OK {
			continue
		}
		found = true
		return r, found
	}
	// Return an empty result + found status if not found.
	return r, found
}
