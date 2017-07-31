package house

import (
    "testing"
    "os"
)

func TestCanWriteScriptToFile(t *testing.T) {
    testScript := "test.sh"
    scriptContent := "echo hi\n"

    fp, _ := os.Create(testScript)
    fp.WriteString(scriptContent)
    fp.Close()

    output, oops := Execute(testScript)
    if oops != nil {
        t.Error("Couldn't execute script")
    }
    if output != "hi" {
        t.Error("something went wrong in the script execution")
    }

    os.Remove(testScript)
}
