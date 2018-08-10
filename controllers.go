package house

import (
    "errors"
    "github.com/docopt/docopt-go"
)

/* #########################
   # CONTROLLER DEFINITION #
   ######################### */

// Defining controller kinds
type ControllerKind int

const (
    INVALID ControllerKind = iota
    LOAD
    UPLOAD
    GET
    EDIT
    BUILD
    EXECUTE
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

/* ##########
   # DOCOPT #
   ########## */

// Gets the documentation for the program.
func GetDocumentation() string {
    return `House b0.8.0

Usage:
  house help
  house get <repo>
  house load [<repo>]
  house upload [(-m <message>)]
  house upload <repo> [(-m <message>)]
  house build [<repo>]
  house edit [<repo>]
  house execute [<repo>] [(-a <arguments>...)]
  `
}

// Creates a new controller based on the provided os arguments
func Generate(args []string) Controller {
    // Listing all possible configurations
    var config struct {
        Help bool
        Get bool
        Load bool
        Upload bool
        Edit bool
        Build bool
        Execute bool
        A bool
        M bool
        Repo string
        Message string
        Arguments []string
    }

    // Parsing arguments
    usage := GetDocumentation()
    parser := &docopt.Parser {
        HelpHandler: func(err error, usage string) {
        },
    }
    options, _ := parser.ParseArgs(usage, args[1:], "b0.8.0")
    options.Bind(&config)

    // Clarifying source repository
    repo := "."
    if len(config.Repo) > 1 {
        repo = config.Repo
    }

    // Building controller
    if config.Load {
        return NewLoadController(repo)
    } else if config.Upload {
        uploader := NewUploadController(repo)
        if message := config.Message; len(message) > 0 {
            uploader.SetMessage(message)
        }
        return uploader
    } else if config.Build {
        return NewBuildController(repo)
    } else if config.Edit {
        return NewEditController(repo)
  	} else if config.Get {
  	  	return NewGetController(repo)
  	} else if config.Execute {
        return GenerateExecuteTool(args)
    }

    return BasicController {
        Kind: INVALID,
    }
}

// Creates an ExecuteController
// TODO Discover why I am not working
func GenerateExecuteTool(args []string) ExecuteController {
    // Listing all possible configurations
    var config struct {
        Help bool
        Get bool
        Load bool
        Upload bool
        Edit bool
        Build bool
        Execute bool
        A bool
        M bool
        Repo string
        Message string
        Arguments []string
    }

    // Parsing arguments
    usage := GetDocumentation()
    parser := &docopt.Parser {
        HelpHandler: func(err error, usage string) {
        },
    }
    options, _ := parser.ParseArgs(usage, args[1:], "b0.8.0")
    options.Bind(&config)

    // Generating controller
    repo := "."
    if len(config.Repo) > 1 {
        repo = config.Repo
    }
    controller := NewExecuteController(repo)
    if len(config.Arguments) > 0 {
        controller.ParseArguments(config.Arguments[0])
    }
    return controller
}
