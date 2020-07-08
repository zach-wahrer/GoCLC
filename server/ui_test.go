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

func TestCommandGoodbye(t *testing.T) {
	got := runCommand("/goodbye")
	if got != serverGoodbye {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", serverGoodbye, got)
	}
}

func TestCommandHelp(t *testing.T) {
	got := runCommand("/help")
	if got != helpMessage {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", helpMessage, got)
	}
}

func TestCommandUnknown(t *testing.T) {
	got := runCommand("/gibberish")
	if got != unknownCommandError {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", unknownCommandError, got)
	}
}
