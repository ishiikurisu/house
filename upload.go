package house

import (
    "github.com/ishiikurisu/house/dishwasher"
    "fmt"
)

// Defines the upload controller
type UploadController struct {
    Kind ControllerKind
    Source string
    Message string
    Remote string
    Branch string
}

// Creates a new upload controller
func NewUploadController(source string) UploadController {
    return UploadController {
        Kind: UPLOAD,
        Source: source,
        Message: "",
        Remote: "origin",
        Branch: "master",
    }
}

// Sets a new message for the upload controller
func (controller *UploadController) SetMessage(message string) {
    controller.Message = message
}

// Sets the remote target for the upload controller
func (controller *UploadController) SetTarget(remote string) {
    controller.Remote = remote
}

func (controller *UploadController) SetBranch(branch string) {
    controller.Branch = branch
}

// Generates a
func (controller *UploadController) GeneratePushCommand() string {
    return fmt.Sprintf("git push %s %s", controller.Remote, controller.Branch)
}

// Uploads the git repository. Returns the standard output from the execution
// of the upload sequence
func (controller UploadController) Execute() (string, error) {
    commander := dishwasher.NewDishwasher()

    if controller.Source != "." {
        commander.Cd("src")
        commander.Cd(controller.Source)
    }

    commander.Branch(controller.Branch)
    commander.RunCustomCommand("git add -A")
    commander.Commit(controller.Message)
    commander.RunCustomCommand(controller.GeneratePushCommand())

    return commander.Execute()
}

// Defines how to get the upload controller kind.
func (controller UploadController) GetKind() ControllerKind {
    return controller.Kind
}
