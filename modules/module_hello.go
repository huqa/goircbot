package modules

import (
    irc "github.com/fluffle/goirc/client"
)

const callCommand = "hello"

type HelloModule struct {
    command string
}

func (m *HelloModule) Init() {}

func (m *HelloModule) Run(conn *irc.Conn, line *irc.Line) {
    conn.Privmsg(line.Target(), "Moronääs :DD")
}

func (m *HelloModule) Command() string {
    return m.command
}

func NewHelloModule() *HelloModule {
    return &HelloModule{command: callCommand}
}
