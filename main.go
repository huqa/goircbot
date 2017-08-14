package main

import (
    "fmt"
    config "github.com/huqa/goircbot/config"
)

func main() {

    fmt.Println("Hello there!")
    //fmt.Printf("List of channels: %v\n", config.IrcChannels)

    c := config.LoadConfig("./bot_config.json")

    fmt.Printf("List of channels: %v\n", c.Channels)

    // create bot
    bot, err := CreateBot(c)

    //quit := make(chan bool)
    //conn.HandleFunc(irc.DISCONNECTED,
    //    func(conn *irc.Conn, line *irc.Line) { quit <- true })
    // Tell client to connect.
    //if err := conn.Connect(); err != nil {
    //    fmt.Printf("Connection error: %s\n", err.Error())
    //}
    //<-quit
    if err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }
    fmt.Println(bot)
}

//func createIRCConfiguration(c config.BotConfig) *irc.Config {
//    cfg := irc.NewConfig(c.Nick, c.Ident, c.Realname)
//    cfg.QuitMessage = c.QuitMessage
//    cfg.Version = c.Version
//    cfg.SSL = false
//    cfg.Server = c.Server
//    return cfg
//}
