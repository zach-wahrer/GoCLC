package server

var exitCommands = map[string]bool{"/quit": true, "/exit": true, "/q": true}

func runCommand(command string) string {
	switch command {
	case "/greet":
		return serverGreeting
	case "/askUsername":
		return askUsername
	case "/goodbye":
		return serverGoodbye
	case "/help":
		return helpMessage
	default:
		return unknownCommandError
	}
}
