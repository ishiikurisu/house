package house

// #include "./benchmark/benchmark.h"
//
import "C"
import "errors"
import "fmt"

// Executes the script in the file identified by the source string
func Execute(script string) (int, error) {
    shell := "sh"
    if os := C.get_os(); os == C.WINDOWS_OS {
      shell = "cmd /B "
    }

    args := C.CString(fmt.Sprintf("%s %s", shell, script))
    oops := errors.New("Not executing correctly")
    output := int(C.sysexec(args))
    if output == 0 {
        oops = nil
    }

    return output, oops
}
