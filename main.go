package main

import (
    "fmt"
    "g-law/riot"
)


func main() {
    fmt.Println("hello world")
    // s := riot.GetSummonerByName("Oscillation")
    c := riot.GetFreeChampion(1)
    // fmt.Printf("%s\n", s)
    fmt.Printf("%s\n", c)
}
