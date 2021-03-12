package config

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config defines the configuration struct for importing settings from TOML.
type Config struct {
	General  general
	Database database
	Github   github
}

// general defines the substruct about general application settings.
type general struct {
	Version string
}

// database defines database specific config settings.
type database struct {
	Path string
}

// github defines provider specific config settings.
type github struct {
	Key string
}

var (
	// Global is the configuration struct for the application.
	Global     Config
	appVersion string
	path       string
)

// initialize the app config system. If a config doesn't exist, create one.
// If the config is out of date read the current config and rebuild with new fields.
func init() {
	// Determine the current user to build expected file path.
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// Create expected config path.
	path = filepath.Join(user.HomeDir, ".config", "lookout", "lookout.config")
	readConf(&Global)
	// If the configuration version has changed update the config to the new
	// format while keeping the user's preferences.
	if Global.General.Version != appVersion {
		reloadConf()
		readConf(&Global)
	}
}

// Read the config or create a new one if it doesn't exist.
func readConf(conf *Config) {
	_, err := toml.DecodeFile(path, &conf)
	if os.IsNotExist(err) {
		genConf(defaultConf())
		readConf(conf)
	}
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
}
