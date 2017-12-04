package house

import (
    "testing"
    "fmt"
)

func TestCanWriteScriptToFile(t *testing.T) {
    testScript := GenerateScriptName("test")
    scriptContent := []string {
        "echo hi",
    }
    oops := CreateScript(testScript, scriptContent)
    if oops != nil {
        t.Error("Couldn't save the script")
        return
    }
    defer DeleteScript(testScript)

    output, oops := Execute(testScript)
    if oops != nil {
        t.Error(fmt.Sprintf("Couldn't execute script with error %d", output))
    }
}

func TestCanIdentifyControllersCorrectly(t *testing.T) {
    // Load controller
    args := []string {
        "house",
        "load",
        "github.com/ishiikurisu/logeybot",
    }
    controller := Generate(args)
    if controller.GetKind() != LOAD {
        t.Error("Wrong controller kind")
    }
    _, oops := controller.Execute()
    if oops != nil {
        t.Error(fmt.Sprintf("Couldn't load this repo. Cause: %s\n", oops.Error()))
    }

    args = []string {
        "house",
        "load",
    }
    controller = Generate(args)
    if controller.GetKind() != LOAD {
        t.Error("Wrong controller kind")
    }
    _, oops = controller.Execute()
    if oops != nil {
        t.Error(fmt.Sprintf("Couldn't load this repo. Cause: %s\n", oops.Error()))
    }

    // Basic controller
    args = []string {
        "house",
        "kill",
    }
    controller = Generate(args)
    if controller.GetKind() != INVALID {
        t.Error("Are you mad, bro?")
    }

    // Upload controller
    args = []string {
        "house",
        "upload",
    }
    controller = Generate(args)
    if controller.GetKind() != UPLOAD {
        t.Error("Wrong controller kind")
    }
    if _, oops = controller.Execute(); oops != nil {
        t.Error("Why are uploading from something that is not a repo?")
    }

    // Build controller
    args = []string {
        "house",
        "build",
    }
    controller = Generate(args)
    if controller.GetKind() != BUILD {
        t.Error("Wrong controller kind: should be BuildController")
    }
    if _, oops = controller.Execute(); oops != nil {
        t.Error(fmt.Sprintf("Couldn't build itself: %v\n", oops))
    }
}

func TestCanGoFroAndToSomeDirs(t *testing.T) {
    dir := "pokemon"
    correctAnswers := []string {
        "cd src",
        "cd pokemon",
    }
    givenAnswers := GoTo(dir)
    for i, correctAnswer := range correctAnswers {
        if correctAnswer != givenAnswers[i] {
            t.Error(fmt.Sprintf("This answer is not correct: %s", correctAnswer))
        }
    }

    dir = "github.com/ishiikurisu/house"
    correctAnswers = []string {
        "cd src",
        "cd github.com",
        "cd ishiikurisu",
        "cd house",
    }
    givenAnswers = GoTo(dir)
    for i, correctAnswer := range correctAnswers {
        if correctAnswer != givenAnswers[i] {
            t.Error(fmt.Sprintf("This answer is not correct: %s", correctAnswer))
        }
    }

    dir = "github.com/ishiikurisu/house"
    givenAnswers = GoFrom(dir)
    if 4 != len(givenAnswers) {
        t.Error(fmt.Sprintf("This answer is not correct: expected 4, got %d", len(givenAnswers)))
    }
}

func TestBuildControllerCanLoadConfigurationFile(t *testing.T) {
    json := `{
        "build": {
            "local": true,
            "commands": [
                "./configure",
                "make"
            ]
        }
    }`
    jsonFile := "testhouse.json"
    CreateScript(jsonFile, []string{json})
    defer DeleteScript(jsonFile)
    config, oops := LoadArbitraryConfig(jsonFile)

    if oops != nil {
        t.Error("Couldn't load configuration file")
        return
    }

    // Checking for local build procedure
    if !config.IsLocal() {
        t.Error("this is a local build procedure!")
    }

    // Checking expected commands
    expected := []string {
        "./configure",
        "make",
    }
    extracted := config.BuildCommands
    for i, _ := range expected {
        if expected[i] != extracted[i] {
            t.Error("Commands were not laoded correctly")
            break
        }
    }
}
