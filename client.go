package glaw

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"text/template"
)

type Client struct {
	client *http.Client
	token  string
}

type Option interface {
	apply(*Client) (*Client, error)
}

func NewClient(token string, opts ...Option) (c *Client, err error) {
	c = &Client{
		client: http.DefaultClient,
		token:  token,
	}

	for _, opt := range opts {
		if c, err := opt.apply(c); err != nil {
			return c, err
		}
	}

	return c, err
}

func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	req.Header.Add("X-Riot-Token", c.token)
	resp, err = c.client.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func ProcessResponse(resp *http.Response, to interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(to)
}

type Request struct {
	Method  string
	Domain  string
	Version Version
	Region  Region
	Uri     string
	Body    io.Reader
}

func (r Request) URL() string {
	t := template.Must(template.New(`url`).Parse(partial))
	b := &strings.Builder{}

	t.Execute(b, r)
	return b.String()
}

func (r Request) NewHttpRequestWithCtx(ctx context.Context) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, r.Method, r.URL(), r.Body)
}

func (r Request) NewHttpRequest() *http.Request {
	ctx := context.Background()

	return r.NewHttpRequestWithCtx(ctx)
}

func NewRequest(method string, region Region, version Version) {
	return Request{
		Method:  method,
		Version: version,
		Region:  region,
	}
}

type Version string

const (
	V3 Version = `v3`
	V4 Version = `v4`
)

const partial = `"https://{{.Region.Base()}}/lol/{{.Domain}}/{{.Version}}/{{.Uri}}`
