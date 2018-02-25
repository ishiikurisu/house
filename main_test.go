package house

import (
    "testing"
    "fmt"
)

func TestCanIdentifyControllersCorrectly(t *testing.T) {
    // Load controller
    args := []string {
        "house",
        "load",
        "github.com/ishiikurisu/logeybot",
    }
    controller := Generate(args)
    if controller.GetKind() != LOAD {
        t.Error("Wrong controller kind")
    }
    _, oops := controller.Execute()
    if oops != nil {
        t.Error(fmt.Sprintf("Couldn't load this repo. Cause: %s\n", oops.Error()))
    }

    args = []string {
        "house",
        "load",
    }
    controller = Generate(args)
    if controller.GetKind() != LOAD {
        t.Error("Wrong controller kind")
    }
    _, oops = controller.Execute()
    if oops != nil {
        t.Error(fmt.Sprintf("Couldn't load this repo. Cause: %s\n", oops.Error()))
    }

    // Basic controller
    args = []string {
        "house",
        "kill",
    }
    controller = Generate(args)
    if controller.GetKind() != INVALID {
        t.Error("Are you mad, bro?")
    }

    // Upload controller
    args = []string {
        "house",
        "upload",
    }
    controller = Generate(args)
    if controller.GetKind() != UPLOAD {
        t.Error("Wrong controller kind")
    }
    if _, oops = controller.Execute(); oops != nil {
        t.Error("Why are uploading from something that is not a repo?")
    }

    // Build controller
    args = []string {
        "house",
        "build",
    }
    controller = Generate(args)
    if controller.GetKind() != BUILD {
        t.Error("Wrong controller kind: should be BuildController")
    }
    if _, oops = controller.Execute(); oops != nil {
        t.Error(fmt.Sprintf("Couldn't build itself: %v\n", oops))
    }
}

func TestCanGoFroAndToSomeDirs(t *testing.T) {
    cmd := NewCommander()
    cmd.GetPwd()
    cmd.Cd("main")
    cmd.GetPwd()

    output, oops := cmd.Execute()
    if oops != nil {
        t.Error("Couldn't get PWD.")
    } else {
        fmt.Println(output)
    }

    cmd.Cd("house")
    cmd.GetPwd()

    output, oops = cmd.Execute()
    if oops == nil {
        t.Error("Changing to inexistent directory.")
    } else {
        fmt.Println(oops)
    }
}
