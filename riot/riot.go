package riot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/coltiebaby/g-law/config"
)

var Client = &http.Client{}
var c = config.GetConfig()

type RiotRequest struct {
	Type    string
	Uri     string
	Version string
	Params  map[string]string
}

func (rr RiotRequest) GetData() *http.Request {
	values := url.Values{}
	for k, v := range rr.Params {
		values.Add(k, v)
	}
	u := &url.URL{
		Scheme:   "https",
		Host:     "na1.api.riotgames.com",
		Path:     fmt.Sprintf("lol/%s/%s/%s", rr.Type, c.Version, rr.Uri),
		RawQuery: values.Encode(),
	}

	fmt.Println("%s -- %s", rr.Type, u)

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("X-Riot-Token", c.Api.Token)
	resp, _ := Client.Do(req)

	return resp
}
