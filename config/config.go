package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Config struct {
	Api struct {
		Token string `yaml:"token"`
	}
	Version string `yaml:"version"`
}

func GetConfig() Config {
	config_path, err := filepath.Abs("config/config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	file_contents, _ := ioutil.ReadFile(config_path)
	config := Config{}

	err = yaml.Unmarshal(file_contents, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}
