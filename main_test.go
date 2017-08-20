package house

import (
    "testing"
    "os"
    "fmt"
)

func TestCanWriteScriptToFile(t *testing.T) {
    testScript := "test.sh"
    scriptContent := "echo hi\n"

    if GetOS() == "win32" {
        testScript = "test.bat"
    }

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
    args := []string {
        "house",
        "load",
        "github.com/ishiikurisu/logeybot",
    }
    controller := Generate(args)
    if controller.GetKind() != LOAD {
        t.Error("Wrong controller kind")
    }

    args = []string {
        "house",
        "kill",
    }
    controller = Generate(args)
    if controller.GetKind() != INVALID {
        t.Error("Are you mad, bro?")
    }
}
