// Declare this file is the main package.
package main

// Import libraries
import (
    // "fmt"
    // "io/ioutil"
    "log"
    "net/http"
    "vs/riot"
)

import m "vs/champions/mastery"

func main() {
    mux := http.NewServeMux()

    // Creating the routes
    mux.Handle("/champions/mastery/all", riot.DefaultOutputHandler(m.MasteryAllChampions))

    log.Println("Listening...")
    http.ListenAndServe(":3000", mux)
}
