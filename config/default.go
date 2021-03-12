package config

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// defaultConf defines the default values for Lookout's configuration.
func defaultConf() Config {
	result := Config{
		General: general{
			// Configuration version number. If a field is added or changed
			// in this default, the version must be changed to tell the app
			// to rebuild the users config files.
			Version: appVersion,
		},
		Database: database{
			// This is the path to the app database.
			// By default it is placed in the same dir as the config.
			Path: filepath.Join(filepath.Dir(path), "apps.db"),
		},
		Github: github{
			// In order to reliably fetch Github projects, an API key
			// must be added here by the user. See CUPPA documentation.
			Key: "",
		},
	}
	return result
}

// genConf encodes the values of the Config stuct back into a TOML file.
func genConf(conf Config) {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(conf)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(buf.Bytes())
}

// reloadConf imports the users config onto a default config and then rewrites
// the configuration file.
func reloadConf() {
	result := defaultConf()
	readConf(&result)
	result.General.Version = appVersion
	genConf(result)
}
