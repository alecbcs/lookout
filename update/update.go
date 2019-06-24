package update

import (
	"github.com/DataDrake/cuppa/providers"
	"github.com/DataDrake/cuppa/results"
)

// CheckUpdate checks a given URL for the latest available release.
func CheckUpdate(archive string) (*results.Result, bool) {
	var r *results.Result
	found := false
	for _, p := range providers.All() {
		name := p.Match(archive)
		if name == "" {
			continue
		}
		r, s := p.Latest(name)
		if s != results.OK {
			continue
		}
		found = true
		return r, found
	}
	return r, found
}
