package config

import (
	"context"
	"fmt"
	"os"
)

type Config struct {
	Token string
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

func CtxGetToken(ctx context.Context) (token string, err error) {
	if t := ctx.Value(ctxKey); t != nil {
		token = t.(string)
	} else {
		err = TokenNotSetErr
	}

	return token, err
}

func CtxSetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, ctxKey, token)
}

const (
	// Name of the environ we want to get
	TOKEN_ENV = "RIOT_API_TOKEN"
)

type confKey string

var (
	ctxKey confKey = confKey("token")
)

// Errors
var (
	TokenNotSetErr error = fmt.Errorf("%s is not set!", TOKEN_ENV)
)
