package house

import (
    "errors"
    "github.com/ishiikurisu/house/dishwasher"
)

// Defines the edit controller
type EditController struct {
    Kind ControllerKind
    Source string
}

// Creates a new load controller
func NewEditController(source string) EditController {
    return EditController {
        Kind: EDIT,
        Source: source,
    }
}

// Defines how to get the build controller kind.
func (controller EditController) GetKind() ControllerKind {
    return controller.Kind
}

// Tries to run the build command in the repo's config.
// Returns the standard output from the execution of the command.
func (controller EditController) Execute() (string, error) {
    commander := dishwasher.NewDishwasher()
    config, oops := LoadConfig(controller.Source)

    if oops != nil {
        return "", oops
    } else if config.GetEditor() == "" {
        return "", errors.New("Editor not set in configuration file")
    }

    if controller.Source != "." {
        controller.Source = "src/" + controller.Source
    }

    commander.RunCustomCommand(config.GetEditor() + " " + controller.Source)
    return commander.Execute()
}
