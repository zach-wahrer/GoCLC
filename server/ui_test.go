package server

import (
	"testing"
)

func TestCommandGreet(t *testing.T) {
	got := runCommand("/greet")
	if got != serverGreeting {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", serverGreeting, got)
	}
}

func TestCommandUnknown(t *testing.T) {
	got := runCommand("/gibberish")
	if got != unknownCommandError {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", unknownCommandError, got)
	}
}
