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

    // Basic controller
    args = []string {
        "house",
        "kill",
    }
    controller = Generate(args)
    if controller.GetKind() != INVALID {
        t.Error("Are you mad, bro?")
    }
}

func TestCanGoFroAndToSomeDirs(t *testing.T) {
    dir := "pokemon"
    correctAnswers := []string {
        "pokemon",
    }
    givenAnswers := GoTo(dir)
    for i, correctAnswer := range correctAnswers {
        if correctAnswer != givenAnswers[i] {
            t.Error(fmt.Sprintf("This answer is not correct: %s", correctAnswer))
        }
    }

    dir = "github.com/ishiikurisu/house"
    correctAnswers = []string {
        "github.com",
        "ishiikurisu",
        "house",
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
