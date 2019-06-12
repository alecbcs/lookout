package update

import (
	"github.com/DataDrake/cuppa/providers"
	"github.com/DataDrake/cuppa/results"
	"strings"
)

// CheckUpdate checks a given URL for the latest available release.
func CheckUpdate(archive string) (version string, location string) {
	for _, p := range providers.All() {
		name := p.Match(archive)
		if name == "" {
			continue
		}
		r, s := p.Latest(name)
		if s != results.OK {
			continue
		}
		test := strings.Join(r.Version, ".")
		return test, r.Location
	}
	return "", ""
}
