package house

import (
    "gopkg.in/yaml.v2"
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
    oops = yaml.Unmarshal(raw, &f)
    if oops != nil {
        return outlet, oops
    }

    everything := f.(map[interface{}]interface{})
    buildStuff := everything["build"].(map[interface{}]interface{})
    if rawLocalBuild, ok := buildStuff["local"]; ok {
        outlet.LocalBuild = rawLocalBuild.(bool)
    } 
    rawCommands := buildStuff["commands"].([]interface{})
    buildCommands := make([]string, len(rawCommands))
    for i, rawCommand := range rawCommands {
        buildCommands[i] = fmt.Sprintf("%v", rawCommand)
    }
    outlet.BuildCommands = buildCommands

    return outlet, nil
}

func LoadConfig(source string) (HouseConfig, error) {
    h := "house.yml"
    if source == "." {
        return LoadArbitraryConfig(h)
    } else {
        return LoadArbitraryConfig(fmt.Sprintf("src/%s/%s", source, h))
    }

}
