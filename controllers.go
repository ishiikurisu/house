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

type ControllerConfiguration struct {
    Help bool
    Get bool
    Load bool
    Upload bool
    Edit bool
    Build bool
    Execute bool
    A bool
    M bool
    R bool
    B bool
    Repo string
    Message string
    Arguments []string
    Remote string
    Branch string
}

// Gets the documentation for the program.
func GetDocumentation() string {
    return `House 0.8.4

Usage:
  house help
  house get <repo>
  house edit [<repo>]
  house load [(-b <branch>)]
  house load <repo> [(-b <branch>)]
  house upload [(-m <message>)] [(-r <remote>)] [(-b <branch>)]
  house upload <repo> [(-m <message>)] [(-r <remote>)] [(-b <branch>)]
  house build [<repo>] [(-a <arguments>...)]
  house execute [<repo>] [(-a <arguments>...)]
  `
}

// Creates a new controller based on the provided os arguments
func ParseConfiguration(args []string) ControllerConfiguration {
    var config ControllerConfiguration
    usage := GetDocumentation()
    parser := &docopt.Parser {
        HelpHandler: func(err error, usage string) {
        },
    }
    options, _ := parser.ParseArgs(usage, args[1:], "0.8.4")
    options.Bind(&config)
    return config
}

func GenerateController(config ControllerConfiguration) Controller {
    // Clarifying source repository
    repo := "."
    if len(config.Repo) > 1 {
      repo = config.Repo
    }

    // Building controller
    if config.Load {
        loadController := NewLoadController(repo)
        if config.B {
            loadController.SetBranch(config.Branch)
        }
        return loadController
    } else if config.Upload {
        uploader := NewUploadController(repo)
        if config.M {
            uploader.SetMessage(config.Message)
        }
        if config.R {
            uploader.SetTarget(config.Remote)
        }
        if config.B {
            uploader.SetBranch(config.Branch)
        }
        return uploader
    } else if config.Build {
        return GenerateBuildTool(config)
    } else if config.Edit {
        return NewEditController(repo)
  	} else if config.Get {
  	  	return NewGetController(repo)
  	} else if config.Execute {
        return GenerateExecuteTool(config)
    }

    return BasicController {
        Kind: INVALID,
    }
}

// Creates an ExecuteController
func GenerateExecuteTool(config ControllerConfiguration) ExecuteController {
    repo := "."
    if len(config.Repo) > 1 {
        repo = config.Repo
    }
    controller := NewExecuteController(repo)
    if len(config.Arguments) > 0 {
        controller.ParseArguments(config.Arguments)
    }
    return controller
}

// Creates a BuildController
func GenerateBuildTool(config ControllerConfiguration) BuildController {
  	repo := "."
    if len(config.Repo) > 1 {
        repo = config.Repo
    }
    controller := NewBuildController(repo)
    if len(config.Arguments) > 0 {
        controller.ParseArguments(config.Arguments)
    }
    return controller
}

// The main House function: parses configurations and generates a controller
func Generate(args []string) Controller {
    config := ParseConfiguration(args)
    return GenerateController(config)
}
