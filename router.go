package main

import (
    "log"
    "net/http"

    "github.com/julienschmidt/httprouter"

    "vs/riot"
)

func main() {
    router := httprouter.New()
    // riot.BuildUrls(router)

    log.Println("Listening...")
    http.ListenAndServe(":3000", router)
}
