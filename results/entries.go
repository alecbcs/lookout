package results

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DataDrake/cuppa/version"
)

// Entry contains the information for a single package from the database.
type Entry struct {
	ID             string
	LatestURL      string
	LatestVersion  version.Version // Latest version is most up-to-date
	CurrentURL     string
	CurrentVersion version.Version // Current version is installed on system.
	Dependencies   []string
	UpToDate       bool
}

// PrintEntry formats a result object correctly to print
func (e Entry) PrintEntry() {
	fmt.Printf("%-20s: %s\n", "Package ID", e.ID)
	fmt.Printf("%-20s: %s\n", "Up-To-Date", strconv.FormatBool(e.UpToDate))
	fmt.Printf("%-20s: %s\n", "LatestVersion", strings.Join(e.LatestVersion, "."))
	fmt.Printf("%-20s: %s\n", "CurrentVersion", strings.Join(e.CurrentVersion, "."))
	fmt.Printf("%-20s: %s\n", "Latest", e.LatestURL)
	fmt.Printf("%-20s: %s\n", "Current", e.CurrentURL)
	if len(e.Dependencies) > 0 {
		fmt.Printf("%-20s: %s\n", "Dependencies", strings.Join(e.Dependencies, ", "))
	}
}

// New creates a new Database Entry structure to store data in.
func New(id string, latest string, latestv version.Version, current string, currentv version.Version, up2date bool) *Entry {
	entry := &Entry{
		strings.ToLower(id),
		latest,
		latestv,
		current,
		currentv,
		[]string{},
		up2date}
	return entry
}

// Patch updates an entry with non blank fields from another entry.
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
