package main

import (
    "fmt"
    irc "github.com/fluffle/goirc/client"
    config "github.com/huqa/goircbot/config"
)

func main() {
    //var client = irc.SimpleClient("Palli")
    c := config.LoadConfig("./bot_config.json")
    //fmt.Println(c.ToString())
    //fmt.Println(c.Nick)
    cfg := irc.Client(createIrcConfiguration(c))

    quit := make(chan bool)
    cfg.HandleFunc(irc.DISCONNECTED,
        func(conn *irc.Conn, line *irc.Line) { quit <- true })
    // Tell client to connect.
    if err := cfg.Connect(); err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }
    <-quit
}

func createIrcConfiguration(c config.BotConfig) *irc.Config {
    cfg := irc.NewConfig(c.Nick)
    cfg.SSL = false
    cfg.Server = c.Server
    return cfg
}
