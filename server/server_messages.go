package server

const ServerColor = "\033[96m"
const ServerTag = "<Server>"
const ServerGreeting = "Welcome to the GoCLC Server!\n"

const AskUsername = "Please enter a username:\n"
const DuplicateUsername = "That username is already taken.\n"
const UserGreeting = "Hello,"
const UserGreetingPunc = "!\n"
const UserAnouncement = "has entered the chat.\n"
const ServerGoodbye = "Goodbye!\n"
const UserDepartedAnnouncement = "has left the chat.\n"

const HelpMessage = "Available Commands:\n" +
	"/greet - show server welcome message\n" +
	"/exit - leave the chat server\n" +
	"/help - prints this help message\n"

const UnknownCommandError = "unknown command\n"
