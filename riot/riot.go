package riot

import (
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
)
import c "vs/config"

const Url = "https://na1.api.riotgames.com/lol"

var config = c.GetConfig()
var Version = config.Version


func add_headers(req *http.Request) {
    req.Header.Add("X-Riot-Token", config.Api.Token)
}

// TODO: Maybe move these over to another thing
// Not sure if this is the best practice
var Client = &http.Client{
//    CheckRedirect: redirectPolicyFunc,
}

type defaultapi func() ([]byte, error)

func DefaultOutputHandler(fn defaultapi) (http.HandlerFunc) {
    var handler = func(w http.ResponseWriter, r *http.Request) {
        body, err := fn()
        if err != nil {
            log.Fatalf("DefaultOutputHandler found an err: %v", err)
        }

        w.Write(body)
    }

    return http.HandlerFunc(handler)
}

func GetData(http_method string, uri string) ([]byte, error) {
    var err error

    url := fmt.Sprintf("%s%s", Url, uri)

    req, _ := http.NewRequest("GET", url, nil)
    add_headers(req)
    resp, _ := Client.Do(req)

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    return body, err
}
