package house

import (
  "fmt"
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

  	// Preparing folder structure
	// TODO Extract required folders from `source` string
  	// github.com/ishiikurisu/house => [github.com, ishiikurisu, house]


	// Executing
  	machine.Cd("src")
  	// TODO cd and mkdir required dir to make this work
  	// Do not forget to NOT change to the last dir, just make it
    machine.RunCustomCommand(fmt.Sprintf("git clone %s", controller.Source))

    return machine.Execute()
}


