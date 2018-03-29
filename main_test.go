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
    if oops == nil {
        t.Error(fmt.Sprintf("Could load this repo. How? %s\n", oops.Error()))
    }

    args = []string {
        "house",
        "load",
    }
    controller = Generate(args)
    if controller.GetKind() != LOAD {
        t.Error("Wrong controller kind: shoud be load kind")
    }
    output, oops := controller.Execute()
    if oops != nil {
        t.Error(fmt.Sprintf("Couldn't load this repo. Cause: %s\n%v\n", oops.Error(), output))
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
        t.Error("Wrong controller kind: should be upload")
    }

    args = []string {
        "house",
        "upload",
        "-m",
        "whatever",
    }
    controller = Generate(args)
    if controller.GetKind() != UPLOAD {
        t.Error("Wrong controller kind: should be upload")
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
}

func TestCanGoFroAndToSomeDirs(t *testing.T) {
    cmd := NewCommander()
    cmd.GetPwd()
    cmd.Cd("main")
    cmd.GetPwd()

    _, oops := cmd.Execute()
    if oops != nil {
        t.Error("Couldn't get PWD.")
    }

    cmd.Cd("house")
    cmd.GetPwd()

    _, oops = cmd.Execute()
    if oops == nil {
        t.Error("Changing to inexistent directory.")
    }
}
