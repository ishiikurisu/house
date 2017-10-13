package house

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
    config, oops := LoadConfig(controller.Source)
    if oops != nil {
        return "", oops
    }

    var commands []string
    var command string
    if config.IsLocal() {
        moreCommands := GoTo(controller.Source)
        for _, command = range moreCommands {
            commands = append(commands, command)
        }
    }
    for _, command = range config.BuildCommands {
        commands = append(commands, command)
    }
    if config.IsLocal() {
        moreCommands := GoFrom(controller.Source)
        for _, command = range moreCommands {
            commands = append(commands, command)
        }
    }

    script := GenerateScriptName("build")
    oops = CreateScript(script, commands)
    defer DeleteScript(script)
    if oops != nil {
        return "", oops
    } else {
        return Execute(script)
    }
}
