package house

import (
    "fmt"
    "os/exec"
    "os"
    "strings"
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


// Adds a new action to the end of the commander actions.
func (cmdr* Commander) Append(action func() (string, error)) {
    cmdr.Actions = append(cmdr.Actions, action)
}

// Gets the current working directory on which the commander is running.
func (cmdr *Commander) GetPwd() {
    cmdr.Append(func() (string, error) {
        cmd := exec.Command("pwd")
        output, oops := cmd.Output()
        return string(output), oops
    })
}

// Changes the current working directory.
func (cmdr *Commander) Cd(where string) {
    cmdr.Append(func() (string, error) {
        oops := os.Chdir(where)
        return "", oops
    })
}

// Executes an arbitrary command.
func (cmdr *Commander) RunCustomCommand(custom string) {
    cmdr.Append(func() (string, error) {
        pieces := strings.Split(custom, " ")
        cmd := exec.Command(pieces[0])
        cmd.Args = pieces
        output, oops := cmd.Output()
        return string(output), oops
    })
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
