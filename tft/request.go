package tft

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/coltiebaby/glaw"
)

const partial = "https://%s/tft/%s/%s/%s"

func NewRequest(method, domain, uri string, region glaw.Region, version glaw.Version) Request {
	return Request{
		Method:  method,
		Domain:  domain,
		Version: version,
		Region:  region,
		Uri:     uri,
	}
}

type Request struct {
	Method  string
	Domain  string
	Version glaw.Version
	Region  glaw.Region
	Uri     string
	Body    io.Reader
}

func (r Request) URL() string {
	return fmt.Sprintf(partial, r.Region.Base(), r.Domain, r.Version, r.Uri)
}

func (r Request) GetRegion() glaw.Region {
	return r.Region
}

func (r Request) NewHttpRequest(ctx context.Context) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, r.Method, r.URL(), r.Body)
}
