package house

// #include "./benchmark/benchmark.h"
import "C"
import "errors"

// Executes the script in the file identified by the source string
func Execute(script string) (int, error) {
    args := C.CString("sh " + script)
    oops := errors.New("Not executing correctly")
    output := int(C.execute(args))

    if output == 0 {
        oops = nil
    }

    return output, oops
}
