package v4

import (
	"fmt"
)

const (
	VERSION string = `v4`
)

func BuildUriFunc(domain string) func(string) string {
	return func(endpoint string) string {
		return fmt.Sprintf("%s/%s/%s", domain, VERSION, endpoint)
	}
}
