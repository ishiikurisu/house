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
    for _, command := range config.BuildCommands {
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
