package v3

const (
	VERSION string = `v3`
)

func BuildUriFunc(domain string) func(string) string {
	return func(endpoint string) string {
		return fmt.Sprintf("%s/%s/%s", domain, VERSION, endpoint)
	}
}
