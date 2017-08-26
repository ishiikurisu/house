package house

import "errors"

// Defines the load controller
type LoadController struct {
    Kind ControllerKind
    Source string
}

// Creates a new load controller
func NewLoadController(source string) LoadController {
    return LoadController {
        Kind: LOAD,
        Source: source,
    }
}

// Loads the git repository. Returns the standard output from the execution
// of `git pull origin master` and an error if its there.
func (controller LoadController) Execute() (string, error) {
    output := ""
    oops := errors.New("Not implemented method")
    // commands := make([]string, 0)
    //
    // // Preparing script
    // if controller.Source != "." {
    //     moreCommands := GoTo(controller.Source)
    //     for _, command = range moreCommands {
    //         commands = append(commands, command)
    //     }
    // }
    //
    // commands = append(commands, "git pull origin master")
    //
    // if controller.Source != "." {
    //     moreCommands := GoFrom(controller.Source)
    //     for _, command = range moreCommands {
    //         commands = append(commands, command)
    //     }
    // }
    //
    // // Executing script

    return output, oops
}

// Defines how to get the load controller kind.
func (controller LoadController) GetKind() ControllerKind {
    return controller.Kind
}
