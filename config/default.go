package config

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func genConf(conf *Config, path string) {
	conf = &Config{
		Database: datab{
			Path: filepath.Join(filepath.Dir(path), "apps.db"),
		},
		Github: provider{
			Key: "",
		},
	}

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
	f.Write(buf.Bytes())
}
