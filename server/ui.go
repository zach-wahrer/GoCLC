package server

var ExitCommands = map[string]bool{"/quit": true, "/exit": true, "/q": true}

func runCommand(command string) string {
	switch command {
	case "/welcome":
		return ServerWelcome
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
