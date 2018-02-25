package house

import (
    "fmt"
    "os/exec"
)

// The commander will execute the required actions for a house tool.
type Commander struct {
    // This is the list of actions that will be called and cleaned after every
    // execution.
    Actions []func() (string, error)

    // IDEA Call the commander dishwasher
}

// Creates an empty commander
func NewCommander() Commander {
    return Commander {
        Actions: make([]func() (string, error), 0),
    }
}

func (cmdr *Commander) GetPwd() {
    getPwd := func() (string, error) {
        cmd := exec.Command("pwd")
        output, oops := cmd.Output()
        return string(output), oops
    }

    cmdr.Actions = append(cmdr.Actions, getPwd)
}

func (cmdr *Commander) Execute() (string, error) {
    var oops error = nil
    var outlet string = ""

    for _, action := range cmdr.Actions {
        output, smallOops := action()
        outlet = fmt.Sprintf("%s%s", outlet, string(output))
        if smallOops != nil {
            oops = smallOops
            break
        }
    }

    if oops == nil {
        cmdr.Actions = make([]func() (string, error), 0)
    }

    return outlet, oops
}
