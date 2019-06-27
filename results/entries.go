package results

import (
	"strings"

	"github.com/DataDrake/cuppa/version"
)

// Entry contains the information for a single package from the database.
type Entry struct {
	ID             string
	LatestURL      string
	LatestVersion  version.Version
	CurrentURL     string
	CurrentVersion version.Version
	UpToDate       bool
}

// New creates a new Database Entry structure to store data in.
func New(id string, latest string, latestv version.Version, current string, currentv version.Version, up2date bool) *Entry {
	entry := &Entry{
		strings.ToLower(id),
		latest,
		latestv,
		current,
		currentv,
		up2date}
	return entry
}

// Patch updates an entry with anothers information.
func Patch(full *Entry, diff *Entry) {
	if diff.CurrentURL != "" {
		full.CurrentURL = diff.CurrentURL
		full.CurrentVersion = diff.CurrentVersion
	}
	if diff.LatestURL != "" {
		full.LatestURL = diff.LatestURL
		full.LatestVersion = diff.LatestVersion
	}
	full.UpToDate = diff.UpToDate
}
