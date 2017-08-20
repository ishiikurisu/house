package house

import (
    "errors"
)

// Useless function to say hi.
func SayHi() string {
    return "hi"
}

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
    Execute(string) (string, error)
    GetKind() ControllerKind
}

// Creates a new controller based on the provided os arguments
func Generate(args []string) Controller {
    controller := BasicController {
        Kind: INVALID,
    }

    if args[1] == "load" {
        controller.Kind = LOAD
    }

    return controller
}

// Description of the basic controller
type BasicController struct {
    Kind ControllerKind
}

func (controller BasicController) Execute(input string) (string, error) {
    return input, errors.New("Basic controllers shouldn't execute")
}

func (controller BasicController) GetKind() ControllerKind {
    return controller.Kind
}
