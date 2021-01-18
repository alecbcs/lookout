package update

import (
	cuppa "github.com/DataDrake/cuppa/config"
	"github.com/DataDrake/cuppa/providers"
	"github.com/DataDrake/cuppa/results"

	"log"
)

// Init sets the GitHub token in CUPPA from the GitHub Token in Lookout.
func Init(key string) {
	// Port Lookout configuration to CUPPA
	cuppa.Global.Github.Key = key
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
		r, err := p.Latest(name)
		if err != nil {
			continue
		}
		found = true
		return r, found
	}
	// Return an empty result + found status if not found.
	return r, found
}
