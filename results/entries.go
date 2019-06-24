package results

import (
	"github.com/DataDrake/cuppa/version"
)

// Entry contains the information for a single package from the database.
type Entry struct {
	ID             string
	LatestURL      string
	LatestVersion  version.Version
	CurrentURL     string
	CurrentVersion version.Version
	OutOFdate      bool
}

// NewEntry creates a new Database Entry structure to store data in.
func NewEntry(id, latest string, latestv string, current string, currentv string, outdate bool) *Entry {
	entry := &Entry{id, latest, version.NewVersion(latestv), current, version.NewVersion(currentv), outdate}
	return entry
}
