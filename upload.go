package house

import "fmt"

// Defines the upload controller
type UploadController struct {
    Kind ControllerKind
    Source string
}

// Creates a new upload controller
func NewUploadController(source string) UploadController {
    return UploadController {
        Kind: UPLOAD,
        Source: source,
    }
}

// Uploads the git repository. Returns the standard output from the execution
// of the upload sequence
func (controller UploadController) Execute() (string, error) {
    commands := make([]string, 0)
    script := GenerateScriptName("upload")

    // Preparing script
    if controller.Source != "." {
        moreCommands := GoTo(controller.Source)
        for _, command := range moreCommands {
            commands = append(commands, command)
        }
    }

    commands = append(commands, "git add -A")
    if controller.Source == "." {
        commands = append(commands, fmt.Sprintf("git checkout %s", script))
    }
    commands = append(commands, "git commit")
    commands = append(commands, "git push origin master")

    if controller.Source != "." {
        moreCommands := GoFrom(controller.Source)
        for _, command := range moreCommands {
            commands = append(commands, command)
        }
    }

    // Executing script
    CreateScript(script, commands)
    defer DeleteScript(script)
    return Execute(script)

}

// Defines how to get the upload controller kind.
func (controller UploadController) GetKind() ControllerKind {
    return controller.Kind
}
