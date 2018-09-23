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

    // This is a command that is ran after every custom command, so more complex
    // actions can be taken into consideration.
    SideEffect func(string, error)
}

// Creates an empty Dishwasher
// By default, it prints the output to the standard output.
func NewDishwasher() Dishwasher {
    return Dishwasher {
        Actions: make([]func() (string, error), 0),
        SideEffect: func(s string, e error) {
            if len(s) > 0 {
                fmt.Print(s)
            }
            if e != nil {
                fmt.Printf("%s\n", e)
            }
        },
    }
}


// Adds a new action to the end of the Dishwasher actions.
func (machine* Dishwasher) Append(action func() (string, error)) {
    machine.Actions = append(machine.Actions, action)
}

// Sets a new side effect to the dishwasher
func (machine *Dishwasher) SetSideEffect(action func(string, error)) {
    machine.SideEffect = action
}

// Gets the current working directory on which the Dishwasher is running.
func (machine *Dishwasher) GetPwd() {
    machine.Append(func() (string, error) {
        cmd := exec.Command("pwd")
        output, oops := cmd.CombinedOutput()
        machine.SideEffect("", oops)
        return string(output), oops
    })
}

// Changes the current working directory.
func (machine *Dishwasher) Cd(where string) {
    machine.Append(func() (string, error) {
        oops := os.Chdir(where)
        machine.SideEffect("", oops)
        return "", oops
    })
}

// Creates a new directory
func (machine *Dishwasher) MkDir(where string) {
    machine.Append(func() (string, error) {
        return fmt.Sprintf("mkdir %s", where), os.MkdirAll(where, os.ModeDir)
    })
}

// Performs a git commit in the current working directory
func (machine *Dishwasher) Commit(message string) {
    machine.Append(func() (string, error) {
        cmd := exec.Command("git", "commit")
        if len(message) > 0 {
            cmd = exec.Command("git", "commit", "-m", message)
        }
        cmd.Stdin = os.Stdin
        output, oops := cmd.CombinedOutput()
        machine.SideEffect(string(output), oops)
        return string(output), oops
    })
}

// This function executes a command and returns its standard output and
// error messages
func RunCommand(custom string) (string, error) {
    custom = strings.TrimSpace(strings.TrimSuffix(custom, "$"))
    pieces := strings.Split(custom, " ")
    cmd := exec.Command(pieces[0])
    cmd.Args = pieces
    cmd.Stdin = os.Stdin
    output, oops := cmd.CombinedOutput()
    return string(output), oops
}

// Executes an arbitrary command.
func (machine *Dishwasher) RunCustomCommand(custom string) {
    machine.Append(func() (string, error) {
        var output string = ""
        var oops error = nil

        custom = strings.TrimSpace(custom)
        if custom[len(custom)-1] == '$' {
            go RunCommand(custom)
        } else {
            output, oops = RunCommand(custom)
            machine.SideEffect(output, oops)
        }

        return output, oops
    })
}

func (machine *Dishwasher) Execute() (string, error) {
    var oops error = nil
    var outlet string = ""

    for i, action := range machine.Actions {
        output, smallOops := action()
        outlet = fmt.Sprintf("%s%s", outlet, string(output))
        if smallOops != nil {
		  	outlet = fmt.Sprintf("%s\n%s", outlet, smallOops)
            oops = errors.New(fmt.Sprintf("Check step %d", i+1))
            break
        }
    }

    if oops == nil {
        machine.Actions = make([]func() (string, error), 0)
    }

    return outlet, oops
}
