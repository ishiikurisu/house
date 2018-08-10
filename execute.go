package house

import (
    // "github.com/ishiikurisu/house/dishwasher"
    "errors"
)

// Defines the execute controller
type ExecuteController struct {
    Kind ControllerKind
    Source string
    Arguments map[string]string
}

// Creates a new execute controller
func NewExecuteController(source string) ExecuteController {
    return ExecuteController {
        Kind: EXECUTE,
        Source: source,
        Arguments: make(map[string]string),
    }
}

// Parse a string to set new variables
func (controller *ExecuteController) ParseArguments(inlet []string) {
    key := ""
    for _, it := range inlet {
        if len(key) == 0 {
            key = it[1:]
        } else {
            controller.Arguments[key] = it
            key = ""
        }
    }
}

// Defines how to get the execute controller kind.
func (controller ExecuteController) GetKind() ControllerKind {
    return controller.Kind
}

// Tries to run the build command in the repo's config.
// Returns the standard output from the execution of the command.
func (controller ExecuteController) Execute() (string, error) {
    // TODO Implement me when you found out how to get arguments in Docopt
  	// commander := dishwasher.NewDishwasher()
    outlet := ""
    return outlet, errors.New("Not implemented yet")
}
