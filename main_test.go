package house

import (
    "testing"
    "os"
    "fmt"
)

func TestCanWriteScriptToFile(t *testing.T) {
    testScript := GenerateScriptName("test")
    scriptContent := "echo hi\n"
    fp, _ := os.Create(testScript)
    fp.WriteString(scriptContent)
    fp.Close()

    output, oops := Execute(testScript)
    if oops != nil {
        t.Error(fmt.Sprintf("Couldn't execute script with error %d", output))
    }

    os.Remove(testScript)
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
