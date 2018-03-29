package house

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "fmt"
)

type HouseConfig struct {
    LocalBuild bool
    BuildCommands []string
    Editor string
}

func NewHouseConfig() HouseConfig {
    return HouseConfig {
        LocalBuild: false,
        BuildCommands: []string {},
        Editor: "",
    }
}

func (h HouseConfig) IsLocal() bool {
    return h.LocalBuild
}

func (h HouseConfig) GetCommands() []string {
    return h.BuildCommands
}

func (h HouseConfig) GetEditor() string {
    return h.Editor
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

    // Getting build parameters
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

    // Getting edit parameters
    if rawEditStuff, ok := everything["edit"]; ok {
        editStuff := rawEditStuff.(map[interface{}]interface{})
        editor := editStuff["editor"].(string)
        outlet.Editor = editor
    }

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
