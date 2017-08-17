package main

import (
    "fmt"
    config "github.com/huqa/goircbot/config"
    irc "github.com/fluffle/goirc/client"
    modules "github.com/huqa/goircbot/modules"
)

type IrcBot struct {
    owner       string
    connection  *irc.Conn
    config      *irc.Config
    channels    []string
    modules     *modules.BotModules
    removers map[string]irc.Remover
}

var handlers = map[string]func(*IrcBot) irc.HandlerFunc{
    irc.PRIVMSG:        (*IrcBot).privmsg,
    irc.CONNECTED:      (*IrcBot).connected,
    irc.DISCONNECTED:   (*IrcBot).disconnected,
}

// Creates a new IRC bot with a configuration parameter
func CreateBot(c config.BotConfig) (*IrcBot, error) {
    cfg := createIRCConfiguration(c)
    bot := &IrcBot{
        config:     cfg,
        channels:   c.Channels,
    }
    bot.connection = irc.Client(cfg)
    bot.removers = make(map[string]irc.Remover)
    // Add event handlers
    for event, handler := range handlers {
        // HandleFunc returns a handler remover in case we need to remove them
        bot.removers[event] = bot.connection.HandleFunc(event, handler(bot))
    }

    // TODO configfy prefix
    bot.modules = modules.NewBotModules("!")
    bot.modules.InitModules()

    // TODO add runnable modules
    // Connect to server
    if err := bot.connection.Connect(); err != nil {
        return nil, err
    }
    return bot, nil
}

func createIRCConfiguration(c config.BotConfig) *irc.Config {
    cfg := irc.NewConfig(c.Nick, c.Ident, c.Realname)
    cfg.QuitMessage = c.QuitMessage
    cfg.Version = c.Version
    cfg.SSL = false
    cfg.Server = c.Server
    return cfg
}

// Do something on server connect
func (c *IrcBot) connected() irc.HandlerFunc {
    return func(conn *irc.Conn, line *irc.Line) {
        for _,channel := range c.channels {
            c.connection.Join(channel)
        }
    }
}

func (c *IrcBot) disconnected() irc.HandlerFunc {
    return func(conn *irc.Conn, line *irc.Line) {
        // Rejoin server on disconnect
        conn.Connect()
    }
}

func (c *IrcBot) privmsg() irc.HandlerFunc {
    return func(conn *irc.Conn, line *irc.Line) {
        // TODO do commands
        fmt.Println(line)
        c.modules.RunModule(conn, line)
    }
} 

