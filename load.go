package house

import "errors"

// Defines the load controller
type LoadController struct {
    Kind ControllerKind
    Source string
}

// Creates a new load controller
func NewLoadController(source string) LoadController {
    return LoadController {
        Kind: LOAD,
        Source: source,
    }
}

// Loads the git repository. Returns the standard output from the execution
// of `git pull origin master` and an error if its there.
func (controller LoadController) Execute() (string, error) {
    return "", errors.New("Maintenance mode")
}

// Defines how to get the load controller kind.
func (controller LoadController) GetKind() ControllerKind {
    return controller.Kind
}
