package main

import (
    "os"
    "fmt"
    "github.com/ishiikurisu/house"
)

func main() {
    controller := house.Generate(os.Args)
    output, oops := controller.Execute()

    if oops == nil {
        if len(output) > 0 {
            fmt.Println(output)
        }
    } else {
        fmt.Println("---")
        fmt.Printf("Oops: %s\n%s\n", oops, output)
        fmt.Println("...")
    }
}
