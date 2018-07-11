package dishwasher

import "testing"

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
  result, _ := ReplaceParameters(params, command)
  if command != result {
	t.Error("Dishwasher is seeing stuff where there isnt")
  }
}
