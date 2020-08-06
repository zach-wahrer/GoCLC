package server

import (
	"testing"
)

func TestCommandWelcome(t *testing.T) {
	got := runCommand("/welcome")
	if got != ServerWelcome {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", ServerWelcome, got)
	}
}

func TestCommandAskUsername(t *testing.T) {
	got := runCommand("/AskUsername")
	if got != AskUsername {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", AskUsername, got)
	}
}

func TestCommandGoodbye(t *testing.T) {
	got := runCommand("/goodbye")
	if got != ServerGoodbye {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", ServerGoodbye, got)
	}
}

func TestCommandHelp(t *testing.T) {
	got := runCommand("/help")
	if got != HelpMessage {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", HelpMessage, got)
	}
}

func TestCommandUnknown(t *testing.T) {
	got := runCommand("/gibberish")
	if got != UnknownCommandError {
		t.Errorf("unexpected command response: want \"%s\", got \"%s\"", UnknownCommandError, got)
	}
}
