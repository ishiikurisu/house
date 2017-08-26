package house

// #include "./benchmark/benchmark.h"
import "C"
import "errors"
import "os/exec"
import "strings"
import "os"
import "fmt"

// Executes the script in the file identified by the source string
func Execute(script string) (string, error) {
    cmd := exec.Command("sh", script)
    output := ""

    if GetOS() == "win32" {
        cmd = exec.Command("cmd", "/C", script)
    } else if GetOS() == "nope" {
        return output, errors.New("Unknown OS")
    }

    outlet, oops := cmd.Output()
    if oops == nil {
        output = string(outlet)
    }

    return output, oops
}

// Discovers on which OS this program is running
func GetOS() string {
    osCode := int(C.get_os())
    if osCode == C.WINDOWS_OS {
        return "win32"
    } else if osCode == C.LINUX_OS {
        return "linux"
    } else {
        return "nope"
    }
}

// Generates a script name based on an arbitrary name
func GenerateScriptName(script string) string {
    output := script + ".sh"

    if GetOS() == "win32" {
        output = script + ".bat"
    }

    return output
}

// Saves a script into a file. Returns an error if it can't even open the file.
func CreateScript(where string, what []string) error {
    fp, oops := os.Create(where)
    if oops == nil {
        defer fp.Close()
    } else {
        return oops
    }

    for _, line := range what {
        fp.WriteString(fmt.Sprintf("%s\n", line))
    }

    return nil
}

// Deletes a script file.
func DeleteScript(where string) {
    os.Remove(where)
}

// Separates a string into its directory parts
func GoTo(inlet string) []string {
    return strings.Split(inlet, "/")
}

// Creates a list of `cd ..` to go back from a directory
func GoFrom(inlet string) []string {
    limit := len(GoTo(inlet))
    outlet := make([]string, limit+1)

    for i := 0; i <= limit; i++ {
        outlet[i] = "cd .."
    }

    return outlet
}
