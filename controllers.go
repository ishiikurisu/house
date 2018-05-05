package house

import (
    "errors"
    "github.com/docopt/docopt-go"
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
// used all over the place.
type Controller interface {
    Execute() (string, error)
    GetKind() ControllerKind
}

// Description of the basic controller
type BasicController struct {
    Kind ControllerKind
}

// This is a placeholder to indicate the basic controller is a controller.
// This method should not be executed.
func (controller BasicController) Execute() (string, error) {
    return "", errors.New(GetDocumentation())
}

// Will tell everyone this is a basic controller.
func (controller BasicController) GetKind() ControllerKind {
    return controller.Kind
}

// Gets the documentation for the program.
func GetDocumentation() string {
    return `House 0.6.4

Usage:
  house help
  house load [<repo>]
  house upload [(-m <message>)]
  house upload <repo> [(-m <message>)]
  house build [<repo>]
  house edit [<repo>]
  `
}

// Creates a new controller based on the provided os arguments
func Generate(args []string) Controller {
    // Parsing arguments
    usage := GetDocumentation()
    parser := &docopt.Parser {
        HelpHandler: func(err error, usage string) {
        },
    }
    options, _ := parser.ParseArgs(usage, args[1:], "0.6.4")

    // Clarifying source repository
    repo := "."
    if maybeRepo, oops := options.String("<repo>"); oops == nil {
        repo = maybeRepo
    }

    // Building controller
    if isIt, oops := options.Bool("load"); (oops == nil) && (isIt) {
        return NewLoadController(repo)
    } else if isIt, oops = options.Bool("upload"); (oops == nil) && (isIt) {
        uploader := NewUploadController(repo)
        if message, oops := options.String("<message>"); oops == nil {
            uploader.SetMessage(message)
        }
        return uploader
    } else if isIt, oops = options.Bool("build"); (oops == nil) && (isIt) {
        return NewBuildController(repo)
    } else if isIt, oops = options.Bool("edit"); (oops == nil) && (isIt) {
        return NewEditController(repo)
    }

    return BasicController {
        Kind: INVALID,
    }
}
