package config

import (
    "os"
    "testing"
)

const (
    KEY string = "SOME_KEY"
)

func set() {
    os.Setenv(TOKEN_ENV, KEY)
}

func cleanup() {
    os.Unsetenv(TOKEN_ENV)
}

func TestConfigFromEnv(t *testing.T) {
    c := NewConfig()
    set()
    defer cleanup()

    err := c.FromEnv()
    if err != nil {
        t.Errorf("Got an unexpected error: %s", err)
    }

    if c.Token != KEY {
        t.Error("Failed to get token...")
    }
}

func TestConfigFromEnvError(t *testing.T) {
    c := NewConfig()
    cleanup()

    err := c.FromEnv()
    if err != TokenNotSetErr {
        t.Errorf("Expected TokenNotSetErr but got: %s", err)
    }
}
