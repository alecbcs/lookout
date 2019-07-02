package update

import (
	"github.com/DataDrake/cuppa/providers"
	"github.com/DataDrake/cuppa/results"
)

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
