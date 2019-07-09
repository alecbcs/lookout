package update

import (
	"github.com/DataDrake/cuppa/version"
)

// UpToDate returns a boolean to describe is the application is up-to-date.
func UpToDate(a version.Version, b version.Version) bool {
	return a.Compare(b) >= 0
}
