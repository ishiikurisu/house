package house

import (
    "errors"
)

// The commander will execute the required actions for a house tool.
type Commander struct {
    // This is the list
    Actions []func()

    // IDEA Call the commander dishwasher
}

// Creates an empty commander
func NewCommander() Commander {
    return Commander {
        Actions: make([]func(), 0),
    }
}

func (cmd *Commander) GetPwd() {
    // TODO Implement me!
}

func (cmd *Commander) Execute() (string, error) {
    // TODO Implement me!
    return "", errors.New("Not implemented yet")
}
