package house

import (
    "errors"
)

// Defining controller kinds
type ControllerKind int

const (
    INVALID ControllerKind = iota
    LOAD
    UPLOAD
    GET
    EDIT
    BUILD
)

// The basic controller to be
type Controller interface {
    Execute() (string, error)
    GetKind() ControllerKind
}

// Creates a new controller based on the provided os arguments
func Generate(args []string) Controller {
    if args[1] == "load" {
        src := "."
        if len(args) >= 3 {
            src = args[2]
        }
        return NewLoadController(src)
    } else {
        return BasicController {
            Kind: INVALID,
        }
    }
}

// Description of the basic controller
type BasicController struct {
    Kind ControllerKind
}

// This is a placeholder to indicate the basic controller is a controller.
// This method should not be executed.
func (controller BasicController) Execute() (string, error) {
    return "", errors.New("Basic controllers shouldn't execute")
}

// Will tell everyone this is a basic controller.
func (controller BasicController) GetKind() ControllerKind {
    return controller.Kind
}
