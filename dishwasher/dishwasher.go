package dishwasher

import (
    "fmt"
    "os/exec"
    "os"
    "strings"
    "errors"
)

// The Dishwasher will execute the required actions for a house tool.
type Dishwasher struct {
    // This is the list of actions that will be called and cleaned after every
    // execution.
    Actions []func() (string, error)
}

// Creates an empty Dishwasher
func NewDishwasher() Dishwasher {
    return Dishwasher {
        Actions: make([]func() (string, error), 0),
    }
}


// Adds a new action to the end of the Dishwasher actions.
func (machine* Dishwasher) Append(action func() (string, error)) {
    machine.Actions = append(machine.Actions, action)
}

// Gets the current working directory on which the Dishwasher is running.
func (machine *Dishwasher) GetPwd() {
    machine.Append(func() (string, error) {
        cmd := exec.Command("pwd")
        output, oops := cmd.CombinedOutput()
        return string(output), oops
    })
}

// Changes the current working directory.
func (machine *Dishwasher) Cd(where string) {
    machine.Append(func() (string, error) {
        oops := os.Chdir(where)
        return "", oops
    })
}

func (machine *Dishwasher) Commit(message string) {
    machine.Append(func() (string, error) {
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
func (machine *Dishwasher) RunCustomCommand(custom string) {
    // IDEA When splitting the string, consider stuff inside "" as one piece
    machine.Append(func() (string, error) {
        pieces := strings.Split(custom, " ")
        cmd := exec.Command(pieces[0])
        cmd.Args = pieces
        cmd.Stdin = os.Stdin
        output, oops := cmd.CombinedOutput()
        return string(output), oops
    })
}

func (machine *Dishwasher) Execute() (string, error) {
    var oops error = nil
    var outlet string = ""

    for i, action := range machine.Actions {
        output, smallOops := action()
        outlet = fmt.Sprintf("%s%s", outlet, string(output))
        if smallOops != nil {
            oops = errors.New(fmt.Sprintf("Check step %d\n", i+1))
            break
        }
    }

    if oops == nil {
        machine.Actions = make([]func() (string, error), 0)
    }

    return outlet, oops
}
