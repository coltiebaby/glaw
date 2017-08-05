// Declare this file is the main package.
package main

// Import libraries
import (
    // "fmt"
    // "io/ioutil"
    "log"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

import m "vs/champions/mastery"

type apiNoParams func() ([]byte, error)
type apiParams func(*httprouter.Params) ([]byte, error)

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

func main() {
    router := httprouter.New()

    // Creating the routes
    router.GET("/summoner/:summoner_id/champion/masteries", hasParams(m.MasteryAllChampions))
    //router.GET("/summoner/:summoner_id/champion/:champion_id/mastery", hasParams(m.MasteryGetChampion))
    //router.GET("/summoner/:summoner_id/champion/mastery/sum", hasParams(m.MasterySummonerScore))

    log.Println("Listening...")
    http.ListenAndServe(":3000", router)
}
