package main

import (
    "os"
    "fmt"
    "github.com/ishiikurisu/house"
)

func main() {
    controller := house.Generate(os.Args)

    _, oops := controller.Execute()
    if oops != nil {
        fmt.Printf("... # %s\n", oops)
    }
}
