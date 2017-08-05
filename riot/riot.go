package riot

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"

    "github.com/julienschmidt/httprouter"
)
import c "vs/config"

const Url = "https://na1.api.riotgames.com/lol"

var config = c.GetConfig()
var Version = config.Version

type apiNoParams func() ([]byte, error)
type apiParams func(*httprouter.Params) ([]byte, error)

func add_headers(req *http.Request) {
    req.Header.Add("X-Riot-Token", config.Api.Token)
}

// TODO: Maybe move these over to another thing
// Not sure if this is the best practice
var Client = &http.Client{
//    CheckRedirect: redirectPolicyFunc,
}

func BuildUrls(router *httprouter.Router) {
    // Bread and butter of the package
    mastery_init(router)
    summoner_init(router)
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

func noParams(fn apiNoParams) (httprouter.Handle) {
    return func(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
        body, err := fn()
        if err != nil {
            log.Fatalf("DefaultOutputHandler found an err: %v", err)
        }

        w.Write(body)
    }
}

func hasParams(fn apiParams) (httprouter.Handle) {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        body, err := fn(&ps)
        if err != nil {
            log.Fatalf("DefaultOutputHandler found an err: %v", err)
        }

        w.Write(body)
    }
}
