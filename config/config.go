package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	key string = "GLAW_CONFIG"
)

type Config struct {
	Api struct {
		Token string `yaml:"token"`
	}
	Version  string `yaml:"version"`
	LogLevel bool   `yaml:"log-level"`
}

func GetConfig() Config {
	fp := os.Getenv(key)
	log.Println("Getting fp:", fp)
	config_path, err := filepath.Abs(fp)
	if err != nil {
		log.fatal(err)
	}

	file_contents, _ := ioutil.ReadFile(config_path)
	config := Config{}

	err = yaml.Unmarshal(file_contents, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	lvl, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.fatal(err)
	}
	log.SetLevel(lvl)
	return config
}
