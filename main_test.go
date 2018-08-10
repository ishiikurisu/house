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
    config := ParseConfiguration(args)
    controller := GenerateController(config)
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
    config = ParseConfiguration(args)
    controller = GenerateController(config)
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
    config = ParseConfiguration(args)
    controller = GenerateController(config)
    if controller.GetKind() != INVALID {
        t.Error("Are you mad, bro?")
    }

    // Upload controller
    args = []string {
        "house",
        "upload",
    }
    config = ParseConfiguration(args)
    controller = GenerateController(config)
    if controller.GetKind() != UPLOAD {
        t.Error("Wrong controller kind: should be upload")
    }

    args = []string {
        "house",
        "upload",
        "-m",
        "whatever",
    }
    config = ParseConfiguration(args)
    controller = GenerateController(config)
    if controller.GetKind() != UPLOAD {
        t.Error("Wrong controller kind: should be upload")
    }

    // Build controller
    args = []string {
        "house",
        "build",
    }
    config = ParseConfiguration(args)
    controller = GenerateController(config)
    if controller.GetKind() != BUILD {
        t.Error("Wrong controller kind: should be BuildController")
    }

  	// Get Controller
  	args = []string {
        "house",
        "get",
        "github.com/ishiikurisu/house",
	   }
    config = ParseConfiguration(args)
    controller = GenerateController(config)
  	if controller.GetKind() != GET {
        t.Error("Wrong controller kind: should be GetController")
    }
}

func TestCanExecuteCommandsWithExecuteTool(t *testing.T) {
    args := []string {
        "house",
        "execute",
        "github.com/ishiikurisu/house",
        "-a",
        "@where",
        "Aus Deutschland:",
        "@what",
        "guten morgen Joe",
    }
    config := ParseConfiguration(args)
    controller := GenerateExecuteTool(config)
    where, ok := controller.Arguments["where"]
    if !ok {
        t.Error("Wasn't able to parse all variables correctly")
        return
    }
    if where != "Aus Deutschland:" {
        t.Error("Parsed variable incorrectly")
        return
    }

    where, ok = controller.Arguments["what"]
    if !ok {
        t.Error("Wasn't able to parse all variables correctly")
        return
    }
    if where != "guten morgen Joe" {
        t.Error("Parsed variable incorrectly")
        return
    }
}
