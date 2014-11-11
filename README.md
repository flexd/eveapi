#eveapi
[![GoDoc](https://godoc.org/github.com/flexd/eveapi?status.svg)](https://godoc.org/github.com/flexd/eveapi)
## EVE Online API Client
eveapi is a library that provides access to the EVE Online XML API.

This needs more work, ALPHA STATUS.
Barely anything is done, API subject to change.

## Todo


* Caching
* Everything else
* More things

## Usage

Below is an example which shows some of the calls available currently.
```go
package main

import (
    "fmt"
    "log"

    "github.com/flexd/eveapi"
)

func main() {
    key := eveapi.Key{"somekey", "somevcode"}
    charID := "93014296"
    voidCharID := "93947594"
    //api := &eveapi.API{
    //Server:    eveapi.Tranquility,
    //APIKey:    key,
    //UserAgent: "Hello",
    //Debug:     false,
    //}
    api := eveapi.Simple(key)
    serverStatus, err := api.ServerStatus()
    if err != nil {
        log.Fatalln(err)
        return
    }
    fmt.Println("Online:", serverStatus.Open, "Players:", serverStatus.OnlinePlayers)
    characters, err := api.Names2ID("Void Thought,Loryanna,Cypherous,ThisGuyDoesnOtexist")
    if err != nil {
        log.Fatalln(err)
        return
    }
    fmt.Println("Characters:", characters)
    fmt.Println("First char:", characters[0])
    accounts, err := api.CharAccountBalances(charID)
    if err != nil {
        log.Fatalln(err)
        return
    }
    fmt.Println("Current time:", accounts.CurrentTime, "Cached until:", accounts.CachedUntil)
    for _, c := range accounts.Accounts {
        fmt.Println("AccountID:", c.ID, "Key:", c.Key, "Balance:", c.Balance, "ISK")
    }
    skillqueue, err := api.SkillQueue(voidCharID)
    if err != nil {
        log.Fatalln(err)
        return
    }
    fmt.Println(skillqueue.SkillQueue[0])
}
```
