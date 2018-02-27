package house

import "fmt"

// Defines the upload controller
type UploadController struct {
    Kind ControllerKind
    Source string
    Message string
}

// Creates a new upload controller
func NewUploadController(source string) UploadController {
    return UploadController {
        Kind: UPLOAD,
        Source: source,
        Message: "",
    }
}

// Sets a new message for the upload controller
func (controller *UploadController) SetMessage(message string) {
    controller.Message = message
}

// Uploads the git repository. Returns the standard output from the execution
// of the upload sequence
func (controller UploadController) Execute() (string, error) {
    commander := NewCommander()

    if controller.Source != "." {
        commander.Cd("src")
        commander.Cd(controller.Source)
    }

    commit := "git commit"
    if len(controller.Message) == 0 {
        commit = fmt.Sprintf("%s -m \"%s\"", commit, controller.Message)
    }

    commander.RunCustomCommand("git add -A")
    commander.RunCustomCommand(commit)
    commander.RunCustomCommand("git push origin master")

    return commander.Execute()
}

// Defines how to get the upload controller kind.
func (controller UploadController) GetKind() ControllerKind {
    return controller.Kind
}
