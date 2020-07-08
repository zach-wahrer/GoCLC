package server

func runCommand(command string) string {
	switch command {
	case "/greet":
		return serverGreeting
	case "/goodbye":
		return serverGoodbye
	case "/help":
		return helpMessage
	default:
		return unknownCommandError
	}
}
