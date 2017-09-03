package house

// Defines the load controller
type UploadController struct {
    Kind ControllerKind
    Source string
}

// Creates a new load controller
func NewUploadController(source string) UploadController {
    return UploadController {
        Kind: UPLOAD,
        Source: source,
    }
}

// Loads the git repository. Returns the standard output from the execution
// of `git pull origin master` and an error if its there.
func (controller UploadController) Execute() (string, error) {
    commands := make([]string, 0)

    // Preparing script
    if controller.Source != "." {
        moreCommands := GoTo(controller.Source)
        for _, command := range moreCommands {
            commands = append(commands, command)
        }
    }

    commands = append(commands, "git add -A")
    commands = append(commands, "git commit")
    commands = append(commands, "git push origin master")

    if controller.Source != "." {
        moreCommands := GoFrom(controller.Source)
        for _, command := range moreCommands {
            commands = append(commands, command)
        }
    }

    // Executing script
    script := GenerateScriptName("upload")
    CreateScript(script, commands)
    defer DeleteScript(script)
    return Execute(script)

}

// Defines how to get the load controller kind.
func (controller UploadController) GetKind() ControllerKind {
    return controller.Kind
}
