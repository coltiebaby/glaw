package riot

import (
    "log"
    "net/http"
)
import c "vs/config"

const Url = "https://na1.api.riotgames.com/lol"

var config = c.GetConfig()
var Version = config.Version


func AddHeaders(req *http.Request) {
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
