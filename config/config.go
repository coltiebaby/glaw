package config

import (
	"fmt"
	"os"
)

type Config struct {
	Token string `yaml:"api.token"`
}

func NewConfig() *Config {
	return &Config{}
}

// Fetches the config info from the environ.
// Returns an error if the token is not set.
func (c *Config) FromEnv() error {
	var ok bool
	if c.Token, ok = os.LookupEnv(TOKEN_ENV); !ok {
		return TokenNotSetErr
	}

	return nil
}

const (
    // Name of the environ we want to get
	TOKEN_ENV = "RIOT_API_TOKEN"
)

// General Error passed back if the token is not set
var TokenNotSetErr error = fmt.Errorf("%s is not set!", TOKEN_ENV)
