package main
import "testing"

func TestCommandAliases(t *testing.T) {
	for _, command := range AllCommands {
		if _, ok := CommandAliases[command]; !ok {
			t.Errorf("Command const %d missing from CommandAliases", command)
		}
	}
}
