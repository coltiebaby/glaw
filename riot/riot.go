package riot

import (
    "fmt"
    "net/http"
    "net/url"
    "g-law/config"

    "io/ioutil"
)

var Client = &http.Client{}
var c = config.GetConfig()

type RiotRequest struct {
    Type string
    Uri  string
    Params map[string]string
}

func (rr RiotRequest) GetData() ([]byte) {
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
    res, _ := Client.Do(req)

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    fmt.Printf("%s\n", body)

    return body
}
