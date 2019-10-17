package config

import (
	"context"
	"os"
	"testing"
)

const (
	env string = `RIOT_API_TOKEN`
	KEY string = "SOME_KEY"
)

func set() {
	os.Setenv(env, KEY)
}

func cleanup() {
	os.Unsetenv(env)
}
func TestConfigFromEnvError(t *testing.T) {
	c, err := FromEnv()
	if c != nil || err != TokenNotSetErr {
		t.Error("Token should not been set")
	}
}

func TestConfigFromEnv(t *testing.T) {
	set()
	defer cleanup()

	c, err := FromEnv()

	if c.Token != KEY || err != nil {
		t.Error("Failed to get token...")
	}
}

// func TestFromEnvError(t *testing.T) {
// 	cleanup()
//
// 	FromEnv()
//     err := recover()
// 	if err != TokenNotSetErr {
// 		t.Errorf("Expected TokenNotSetErr but got: %s", err)
// 	}
// }

func TestCtxToken(t *testing.T) {
	ctx := context.Background()
	token, err := CtxGetToken(ctx)
	if err != TokenNotSetErr {
		t.Errorf("Expected TokenNotSetErr but got: %s", err)
	}

	ctx = CtxSetToken(ctx, KEY)
	token, err = CtxGetToken(ctx)
	if err != nil {
		t.Errorf("Got an unexpected error: %s", err)
	}

	if token != KEY {
		t.Error("Failed to get token...")
	}
}
