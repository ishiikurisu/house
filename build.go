package house

import (
  "github.com/ishiikurisu/house/dishwasher"
  "errors"
  "fmt"
  "strconv"
)

// Defines the load controller
type BuildController struct {
    Kind ControllerKind
    Source string
  	Arguments map[string]string
}

// Creates a new load controller
func NewBuildController(source string) BuildController {
    return BuildController {
        Kind: BUILD,
        Source: source,
	  	Arguments: make(map[string]string),
    }
}

// Defines how to get the build controller kind.
func (controller BuildController) GetKind() ControllerKind {
    return controller.Kind
}


// Parse a string to set new variables
func (controller *BuildController) ParseArguments(inlet []string) {
    key := ""
    for _, it := range inlet {
        if len(key) == 0 {
            maybe := it
            if maybe[0] == '@' {
                maybe = maybe[1:]
            }
            key = maybe
        } else {
            controller.Arguments[key] = it
            key = ""
        }
    }
}


// Tries to run the build command in the repo's config.
// Returns the standard output from the execution of the command.
func (controller BuildController) Execute() (string, error) {
  	commander := dishwasher.NewDishwasher()
    config, oops := LoadConfig(controller.Source)
  	offset := 0

    if oops != nil {
        return "", oops
    }

    if config.IsLocal() && (controller.Source != ".") {
	  	offset = 2
        commander.Cd("src")
        commander.Cd(controller.Source)
    }
    for _, rawCommand := range config.BuildCommands {
	  	command, oops := dishwasher.ReplaceParameters(controller.Arguments, rawCommand)
        if oops != nil {
            return "", oops
        }
        commander.RunCustomCommand(command)
    }

  	outlet, oops := commander.Execute()
	if oops != nil {
	  	s := fmt.Sprintf("%s", oops)[len("check step "):]
	  	wrongStep, _ := strconv.Atoi(s)
		wrongStep -= offset
	  	oops = errors.New(fmt.Sprintf("Check step %d", wrongStep))
	}
  	return outlet, oops
}
