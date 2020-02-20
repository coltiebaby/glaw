package glaw

import (
    "context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/coltiebaby/glaw/errors"
	"github.com/coltiebaby/glaw/ratelimit"
)

type Client struct {
    client http.Client
    token string
}

type Option interface {
    apply(*Client) (*Client, error)
}

func NewClient(opts ...Option) (c *Client, err error) {
    c := &Client{}

    for _, opt := range opts {
        if c, err := opt.apply(c); err != nil {
            return c, err
        }
    }

    return c, err
}

func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	// req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	// if err != nil {
	// 	return resp, err
	// }

	req.Header.Add("X-Riot-Token", token)
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

// /lol/platform/v3/champion-rotations
type Request struct {
    Method string
    Domain string
    Version Version
    Region Region
    Uri string
    Body io.Reader
}

func (r Request) NewHttpRequestWithCtx(ctx context.Context) *http.Request {
    template := `https://%s/lol/%s/%s/%s`
    u := fmt.Sprintf(template, r.Region.Base(), Domain, Version, Uri)

    return http.NewRequestWithContext(ctx, r.Method, u, r.Body)
}

func (r Request) NewHttpRequest() *http.Request {
    ctx := context.Background()

    return r.NewHttpRequestWithCtx(ctx)
}

func NewRequest(method string, region Region, version Version) {
   return Request {
       Method: method,
       Version: version,
       Region: region,
    }
}

type Version string

const (
    V3 Version = `v3`
    V4 Version = `v4`
)
