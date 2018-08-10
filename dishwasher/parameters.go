package dishwasher

import (
  "fmt"
  "strings"
)

// Takes a string describing a parameterized template command and returns
// the resulting command depending on the provided parameters from the map.
// The template string defines each variable as a name preceded by an "at"
// symbol. The parameters map should contain the variable names as keys
// (without the "at" symbol) while the keys are desired corresponding values.
func ReplaceParameters(parameters map[string]string, template string) (string, error) {
  outlet := template

  for parameter, value := range parameters {
	variable := fmt.Sprintf("@%s", parameter)
	outlet = strings.Replace(outlet, variable, value, -1)
  }

  return outlet, nil
}
