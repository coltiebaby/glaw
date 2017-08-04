package mastery

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "vs/riot"
)

func MasteryAllChampions() ([]byte, error) {
    // All Champion Mastery
    // Ask Riot for all the summoners champions based on mastery score.

    // TODO: Make an actual error here
    var err error

    uri := "/champion-mastery/%s/champion-masteries/by-summoner/%s"
    uri = fmt.Sprintf(uri, riot.Version, "28747969")
    url := fmt.Sprintf("%s%s", riot.Url, uri)

    req, _ := http.NewRequest("GET", url, nil)
    riot.AddHeaders(req)
    resp, _ := riot.Client.Do(req)

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    return body, err
}
