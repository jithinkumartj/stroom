package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

//To add new section, create a struct. Add values in similar format on input json
//https://stackoverflow.com/a/16466189/1645774
var GlobalConfig Configuration

type Configuration struct {
	ConfigFile string `json:"-"`
	DB         DB
}

type DB struct {
	DatabaseName string
}

func init() {
	absPath, _ := filepath.Abs("config/conf.json")
	flag.StringVar(&GlobalConfig.ConfigFile, "conf", absPath, "")
	flag.Parse()

	LoadConfig()
}

func LoadConfig() (err error) {
	rawConfig, err := ioutil.ReadFile(GlobalConfig.ConfigFile)
	if err != nil {
		log.Fatalf("File error: %v\n", err)
		return
	}

	if err = json.Unmarshal(rawConfig, &GlobalConfig); err != nil {
		log.Error("ERROR while loading config: " + err.Error())
		return
	}

	return err
}
