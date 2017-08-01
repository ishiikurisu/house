package house

import (
    "testing"
    "os"
    "fmt"
)

func TestCanWriteScriptToFile(t *testing.T) {
    testScript := "test.sh"
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
    fmt.Printf("TODO Implement controllers")
}
