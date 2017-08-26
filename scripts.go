package house

// #include "./benchmark/benchmark.h"
import "C"
import "errors"
import "os/exec"
import "strings"

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

// Separates a string into its directory parts
func GoTo(inlet string) []string {
    return strings.Split(inlet, "/")
}
