package server

func runCommand(command string) string {
	switch command {
	case "/greet":
		return serverGreeting
	case "/help":
		return helpMessage
	default:
		return unknownCommandError
	}
}
