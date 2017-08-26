package main

import (
    "os"
    "fmt"
    "github.com/ishiikurisu/house"
)

func main() {
    controller := house.Generate(os.Args)
    output, oops := controller.Execute()

    fmt.Println("---")
    if oops == nil {
        fmt.Println(output)
    } else {
        fmt.Printf("Oops: %s\n", oops)
    }
    fmt.Println("...")
}
