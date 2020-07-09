package server

import (
	"bufio"
	"fmt"
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

	client.Write(runCommand("/greet"))
	client.Write(runCommand("/askUsername"))
	addUser(client.Read())
	client.Write(fmt.Sprintf("%s %s%s", userGreeting, client.recieve.Text(), userGreetingPunc))

	for {
		command := client.Read()

		if command == "/quit" || command == "/exit" || command == "/q" {
			break
		}

		if command[0] == '/' {
			client.Write(runCommand(command))
		} else {
			// Todo - Send out chat message
		}
	}

	client.Write(runCommand("/goodbye"))
}
