package glaw

import (
	"net/http"
	"net/url"
)

type ApiClient interface {
	NewRequest(string) ApiRequest
	ChangeRegion(Region)
	Get(ApiRequest) (*http.Response, error)
}

type ApiRequest interface {
	AddParameter(string, string)
	SetParameters(url.Values)
	Encode() string
	Uri() string
}
