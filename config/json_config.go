package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type BotConfig struct {
    Nick string `json:"nick"`
    Ident string `json:"ident"`
    Realname string `json:"realname"`
    Version string `json:"version"`
    QuitMessage string `json:"quitMessage"`
}

func (c BotConfig) ToString() string {
    return toJson(c)
}

func toJson(c interface{}) string {
    bytes, err := json.Marshal(c)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    return string(bytes)
}

func LoadConfig(filePath string) BotConfig {
    raw, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var config BotConfig
    json.Unmarshal(raw, &config)
    return config
}
