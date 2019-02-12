package house

import "github.com/ishiikurisu/house/dishwasher"
import "fmt"

// Defines the load controller
type LoadController struct {
    Kind ControllerKind
    Source string
    Branch string
}

// Creates a new load controller
func NewLoadController(source string) LoadController {
    return LoadController {
        Kind: LOAD,
        Source: source,
        Branch: "master",
    }
}

func (controller *LoadController) SetBranch(branch string) {
  controller.Branch = branch
}

// Loads the git repository. Returns the standard output from the execution
// of `git pull origin master` and an error if its there.
func (controller LoadController) Execute() (string, error) {
    commander := dishwasher.NewDishwasher()

    if controller.Source != "." {
        commander.Cd("src")
        commander.Cd(controller.Source)
    }

    commander.RunCustomCommand(fmt.Sprintf("git checkout %s", controller.Branch))
    commander.RunCustomCommand(fmt.Sprintf("git pull origin %s", controller.Branch))

    return commander.Execute()
}

// Defines how to get the load controller kind.
func (controller LoadController) GetKind() ControllerKind {
    return controller.Kind
}
