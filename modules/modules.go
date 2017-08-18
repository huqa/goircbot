package modules

import (
    "strings"
    irc "github.com/fluffle/goirc/client"
)

var enabledModules = []RunnableModule{
    NewHelloModule(),
}

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
    for _,em := range enabledModules {
        m.AddRunnableModule(em)
        em.Init()
    }
}

func (m *BotModules) RunModule(conn *irc.Conn, line *irc.Line) {
    //fmt.Println("line.Text() ", line.Text())
    //fmt.Println("line.Public() ", line.Public())
    //fmt.Println("line.Target() ", line.Target())
    var ircLine = line.Text()
    if !strings.HasPrefix(ircLine, m.Prefix) {
        return
    }
    withoutPrefix := strings.Replace(ircLine, m.Prefix, "", 1)
    command := strings.Split(withoutPrefix, " ")[0]
    for _,module := range m.modules {
        if module.Command() == command {
            module.Run(conn, line)
        }
    }
}

