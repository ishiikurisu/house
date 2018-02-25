package house

import "errors"

// Defines the load controller
type BuildController struct {
    Kind ControllerKind
    Source string
}

// Creates a new load controller
func NewBuildController(source string) BuildController {
    return BuildController {
        Kind: BUILD,
        Source: source,
    }
}

// Defines how to get the build controller kind.
func (controller BuildController) GetKind() ControllerKind {
    return controller.Kind
}

// Tries to run the build command in the repo's config.
// Returns the standard output from the execution of the command.
func (controller BuildController) Execute() (string, error) {
    return "", errors.New("Maintenance mode")
}
