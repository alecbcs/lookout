package update

import (
	"fmt"

	"github.com/DataDrake/cuppa/providers"
	"github.com/DataDrake/cuppa/results"
)

// CheckUpdate checks a given URL for the latest available release.
func CheckUpdate(archive string) {

	for _, p := range providers.All() {
		name := p.Match(archive)
		if name == "" {
			continue
		}
		r, s := p.Latest(name)
		if s != results.OK {
			continue
		}
		loc := r.Location
		ver := r.Version
		fmt.Println(loc)
		fmt.Println(ver)

	}
}
