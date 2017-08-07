package main

import (
    "fmt"
    //irc "github.com/fluffle/goirc/client"
    config "github.com/huqa/goircbot/config"
)

func main() {
    //var client = irc.SimpleClient("Palli")
    c := config.LoadConfig("./bot_config.json")
    fmt.Println(c)
}
