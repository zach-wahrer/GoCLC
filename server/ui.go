package server

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

func addUser(username string) {

}
