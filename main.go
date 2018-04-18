package main

import (
    "fmt"
    "os"
    "os/signal"
    "flag"
    config "github.com/huqa/goircbot/config"
)

func main() {

    fmt.Println("Hello there!")
    configPathPtr := flag.String("config", "./bot_config.json", "config file path")
    flag.Parse()

    c := config.LoadConfig(*configPathPtr)

    fmt.Printf("List of channels: %v\n", c.Channels)

    // create bot
    bot, err := CreateBot(c)
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt)

    <-sig
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
