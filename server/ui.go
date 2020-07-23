package server

var exitCommands = map[string]bool{"/quit": true, "/exit": true, "/q": true}

func runCommand(command string) string {
	switch command {
	case "/greet":
		return ServerGreeting
	case "/AskUsername":
		return AskUsername
	case "/DuplicateUsername":
		return DuplicateUsername
	case "/goodbye":
		return ServerGoodbye
	case "/help":
		return HelpMessage
	default:
		return UnknownCommandError
	}
}
