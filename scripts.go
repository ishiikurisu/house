package house

// #include "./benchmark/benchmark.h"
import "C"
import "errors"
import "os/exec"

// Executes the script in the file identified by the source string
func Execute(script string) (int, error) {
    cmd := exec.Command("sh", script)
    output := 0

    if GetOS() == "win32" {
        cmd = exec.Command("cmd", "/C", script)
    } else if GetOS() == "nope" {
        return -1, errors.New("Unknown OS")
    }

    _, oops := cmd.Output()
    if oops != nil {
        output = 2
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
