package house

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
)

type HouseConfig struct {
    LocalBuild bool
    BuildCommands []string
}

func NewHouseConfig() HouseConfig {
    return HouseConfig {
        LocalBuild: false,
        BuildCommands: []string {},
    }
}

func (h HouseConfig) IsLocal() bool {
    return h.LocalBuild
}

func (h HouseConfig) GetCommands() []string {
    return h.BuildCommands
}

func LoadArbitraryConfig(source string) (HouseConfig, error) {
    outlet := NewHouseConfig()
    raw, oops := ioutil.ReadFile(source)
    if oops != nil {
        return outlet, oops
    }

    var f interface{}
    oops = json.Unmarshal(raw, &f)
    if oops != nil {
        return outlet, oops
    } else {
        buildStuff := f.(map[string]interface{})["build"].(map[string]interface{})
        outlet.LocalBuild = buildStuff["local"].(bool)
        rawCommands := buildStuff["commands"].([]interface{})
        buildCommands := make([]string, len(rawCommands))
        for i, rawCommand := range rawCommands {
            buildCommands[i] = fmt.Sprintf("%v", rawCommand)
        }
        outlet.BuildCommands = buildCommands
    }

    return outlet, nil
}
