package house

import (
  "fmt"
  "strings"
  "github.com/ishiikurisu/house/dishwasher"
)

// Defines the load controller
type GetController struct {
    Kind ControllerKind
    Source string
}

// Creates a new load controller
func NewGetController(source string) GetController {
    return GetController {
        Kind: GET,
        Source: source,
    }
}

// Defines how to get the load controller kind.
func (controller GetController) GetKind() ControllerKind {
    return controller.Kind
}

// Clones a remote git repository. Returns the standard output from the
// execution of `git clone <source>` and an error if its there.
func (controller GetController) Execute() (string, error) {
    machine := dishwasher.NewDishwasher()
    folders := strings.Split(controller.Source, "/")
    folder := fmt.Sprintf("src/%s", strings.Join(folders[0:len(folders)-1], "/"))
    machine.MkDir(folder)
    machine.Cd(folder)
    machine.RunCustomCommand(fmt.Sprintf("git clone https://%s", controller.Source))
    return machine.Execute()
}
