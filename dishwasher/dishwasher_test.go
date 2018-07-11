package dishwasher

import "testing"
import "fmt"

func TestCanGoFroAndToSomeDirs(t *testing.T) {
    cmd := NewDishwasher()
    cmd.GetPwd()
    cmd.Cd("..")
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

    cmd.Cd("dishwasher")
    cmd.Execute()
}

func TestCanCheckForDirectoriesExistence(t *testing.T) {
    cmd := NewDishwasher()
    cmd.MkDir("test")
    if _, oops := cmd.Execute(); oops != nil {
        t.Error("Could not create test directory")
    }
}

func TestDishwasherCanReplaceCommands(t *testing.T) {
  params := make(map[string]string)

  // testing command without parameters
  command := "echo 'no one cares'"
  desired := command
  result, _ := ReplaceParameters(params, command)
  if command != result {
	t.Error("Dishwasher is seeing stuff where there isnt")
  }

  // testing command with single parameter
  params["where"] = "/path/to/something"
  command = "python main.py @where"
  desired = "python main.py /path/to/something"
  result, _ = ReplaceParameters(params, command)
  if result != desired {
	msg := fmt.Sprintf("Dishwasher couldnt replace a single variable\n%s\n",
					   result)
	t.Error(msg)
  }

  // TODO Test with multiple parameters
  // TODO Test with unknown variables
  // what if the command should contain a string with a @
  // that is not a house variable?
}
