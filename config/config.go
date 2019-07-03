package config

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database datab
	Github   provider
}

type datab struct {
	Path string
}

type provider struct {
	Key string
}

var Conf *Config

func init() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(user.HomeDir, ".config", "lookout", "lookout.config")
	for true {
		_, err = toml.DecodeFile(path, &Conf)
		if os.IsNotExist(err) {
			genConf(Conf, path)
			continue
		}
		break
	}
}
