package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const serverGreeting = "Welcome to the GoCLC Server!\n"
const askUsername = "Please enter a username:\n"
const userGreeting = "Hello,"
const userGreetingPunc = "!\n"
const serverGoodbye = "Goodbye!\n"
const helpMessage = "Available Commands:\n" +
	"/greet - show server welcome message\n" +
	"/help - prints this help message\n"
const unknownCommandError = "unknown command"

// Listen for incoming connections
func Listen(address, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	var client = Client{c: conn, recieve: bufio.NewScanner(conn)}

	writeToRemote(conn, runCommand("/greet"))

	writeToRemote(conn, runCommand("/askUsername"))
	client.recieve.Scan()
	addUser(client.recieve.Text())
	writeToRemote(conn, fmt.Sprintf("%s %s%s", userGreeting, client.recieve.Text(), userGreetingPunc))

	for client.recieve.Scan() {

		command := client.recieve.Text()
		if command == "/quit" || command == "/exit" || command == "/q" {
			break
		}

		if command[0] == '/' {
			writeToRemote(conn, runCommand(command))
		} else {
			// Todo - Send out chat message
		}
	}

	writeToRemote(conn, runCommand("/goodbye"))
}

func writeToRemote(conn net.Conn, input string) {
	if _, err := io.WriteString(conn, input); err != nil {
		log.Print(err)
	}
}
