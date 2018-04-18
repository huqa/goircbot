package modules

import (
    "testing"
    "reflect"
    "fmt"
    irc "github.com/fluffle/goirc/client"
)

type MyTestModule struct {
    command string
}

func (m *MyTestModule) Init() {
    fmt.Println("Initialize")
}

func (m *MyTestModule) Run(conn *irc.Conn, line *irc.Line) {
    fmt.Println("Run command")
}

func (m *MyTestModule) Command() string {
    return m.command
}

func TestNewBotModules(t *testing.T) {
    b := NewBotModules("!")
    if reflect.TypeOf(*b) != reflect.TypeOf((*BotModules)(nil)).Elem() {
        t.Error("Expected BotModules, got ", b)
    }
}

func TestAddRunnableModule(t *testing.T) {
    b := NewBotModules("!")
    m := &MyTestModule{command: "Testing"}
    b.AddRunnableModule(m)
    if len(b.modules) != 1 {
        t.Error("Expected modules length to be 1, got ", len(b.modules))
    }
    RunnableModuleType := reflect.TypeOf((*RunnableModule)(nil)).Elem()
    for _,module := range b.modules {
        if reflect.TypeOf(&module).Implements(RunnableModuleType) {
            t.Error("Expected RunnableModule, got ", module)
        }
    }
}

func TestInitModules(t *testing.T) {
    b := NewBotModules("!")
    m := &MyTestModule{command: "Testing"}
    //b.AddRunnableModule(m)
    b.InitModules()
}

