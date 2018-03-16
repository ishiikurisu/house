package house

import (
    "fmt"
    "os/exec"
    "os"
    "strings"
    "errors"
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
        output, oops := cmd.CombinedOutput()
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

func (cmdr *Commander) Commit(message string) {
    cmdr.Append(func() (string, error) {
        cmd := exec.Command("git", "commit")
        if len(message) > 0 {
            cmd = exec.Command("git", "commit", "-m", message)
        }
        cmd.Stdin = os.Stdin
        output, oops := cmd.CombinedOutput()
        return string(output), oops
    })
}

// Executes an arbitrary command.
func (cmdr *Commander) RunCustomCommand(custom string) {
    // IDEA When splitting the string, consider stuff inside "" as one piece
    cmdr.Append(func() (string, error) {
        pieces := strings.Split(custom, " ")
        cmd := exec.Command(pieces[0])
        cmd.Args = pieces
        cmd.Stdin = os.Stdin
        output, oops := cmd.CombinedOutput()
        return string(output), oops
    })
}

func (cmdr *Commander) Execute() (string, error) {
    var oops error = nil
    var outlet string = ""

    for i, action := range cmdr.Actions {
        output, smallOops := action()
        outlet = fmt.Sprintf("%s%s", outlet, string(output))
        if smallOops != nil {
            oops = errors.New(fmt.Sprintf("Check step %d\n", i+1))
            break
        }
    }

    if oops == nil {
        cmdr.Actions = make([]func() (string, error), 0)
    }

    return outlet, oops
}
