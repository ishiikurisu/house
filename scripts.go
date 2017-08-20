package house

// #include "./benchmark/benchmark.h"
import "C"
import "errors"

// Executes the script in the file identified by the source string
func Execute(script string) (int, error) {
    args := "sh " + script
    if GetOS() == "win32" {
        args = "cmd /C " + script
    } else if GetOS() == "nope" {
        return -1, errors.New("Unknown OS")
    }

    oops := errors.New("Not executing correctly")
    output := int(C.execute(C.CString(args)))
    if output == 0 {
        oops = nil
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
