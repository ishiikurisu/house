package house

// #include "./benchmark/benchmark.h"
//
import "C"
import "errors"

// Executes the script in the file identified by the source string
func Execute(script string) error {
    args := []*C.char {
        C.CString("echo"),
        C.CString("hi"),
    }

    if output := C.execute(args); output == 0 {
        return nil
    } else {
        return errors.New("Not executing correctly")
    }
}
