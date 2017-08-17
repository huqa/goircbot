package modules

import (
    "fmt"
    irc "github.com/fluffle/goirc/client"
)

type RunnableModule interface {
    Init()
    Run(conn *irc.Conn, line *irc.Line)
    Command() string
}

type BotModules struct {
    Prefix string
    modules []RunnableModule
}

func NewBotModules(prefix string) *BotModules {
    return &BotModules{Prefix: prefix}
}

func (m *BotModules) AddRunnableModule(module RunnableModule) {
    m.modules = append(m.modules, module)
}

func (m *BotModules) InitModules() {
    for _,module := range m.modules {
        module.Init()
    }
}

func (m *BotModules) RunModule(conn *irc.Conn, line *irc.Line) {
    fmt.Println("line.Text() ", line.Text())
    fmt.Println("line.Public() ", line.Public())
    fmt.Println("line.Target() ", line.Target())
    //for _,module := range m.modules {
    //    cmd := module.Command()
        
        // Todo do something with command
    //}
}
