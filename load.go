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

// Describes the execution of the load controller
func (controller LoadController) Execute() (string, error) {
    // TODO Implement this method
    return "", errors.New("Not implemented method")
}

// Defines how to get the load controller kind
func (controller LoadController) GetKind() ControllerKind {
    return controller.Kind
}
