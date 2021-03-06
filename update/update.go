package update

import (
	"net/http"
	"regexp"
	"strings"

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
		match := p.Match(archive)
		if len(match) == 0 {
			continue
		}

		// GitHub will never work without a token
		if p.String() == "GitHub" && cuppa.Global.Github.Key == "" {
			log.Fatal("A GitHub token is required in your $HOME/.config/lookout/lookout.config")
		}

		// Pull the latest (non-beta) release from repository.
		r, err := p.Latest(match)
		if err != nil {
			continue
		}
		if r != nil {
			found = true
			matchLink(archive, r)
			return r, found
		}
	}
	// Return an empty result + found status if not found.
	return r, found
}

func matchLink(archive string, result *results.Result) {
	if archive != "" {
		vexp := regexp.MustCompile(`([0-9]{1,4}[.])+[0-9,a-d]{1,4}`)
		updatedLink := vexp.ReplaceAllString(archive, strings.Join(result.Version, "."))

		resp, err := http.Head(updatedLink)
		if err != nil {
			return
		}
		if resp.StatusCode != http.StatusOK {
			return
		}
		result.Location = updatedLink
	}
}
