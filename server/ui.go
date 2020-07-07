package server

func runCommand(command string) string {
	switch command {
	case "greet":
		return serverGreeting
	default:
		return "goclc: error: unknown command"
	}

}
