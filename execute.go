package house

import (
    "github.com/ishiikurisu/house/dishwasher"
    "errors"
    "fmt"
    "strconv"
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

// Defines how to get the execute controller kind.
func (controller ExecuteController) GetKind() ControllerKind {
    return controller.Kind
}

// Tries to run the execute command in the repo's config.
// Returns the standard output from the execution of the command.
func (controller ExecuteController) Execute() (string, error) {
  	commander := dishwasher.NewDishwasher()
    config, oops := LoadConfig(controller.Source)
    outlet := ""
    offset := 0

    if oops != nil {
        return outlet, oops
    }

    if config.LocalExecution && (controller.Source != ".") {
        offset = 2
        commander.Cd("src")
        commander.Cd(controller.Source)
    }
    for _, rawCommand := range config.ExecutionCommands {
        command, oops := dishwasher.ReplaceParameters(controller.Arguments, rawCommand)
        if oops != nil {
            return outlet, oops
        }
        commander.RunCustomCommand(command)
    }

    outlet, oops = commander.Execute()
    if oops != nil {
        s := fmt.Sprintf("%s", oops)[len("check step "):]
        wrongStep, _ := strconv.Atoi(s)
        wrongStep -= offset
        oops = errors.New(fmt.Sprintf("Check step %d", wrongStep))
    }

    return outlet, oops
}
