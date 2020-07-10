// Package server provides a stand-alone server for the GoCLC chat service.
package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Listen checks for connections on the given address and port and then
// spawns a goroutine to handle each client separately via helper functions.
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
	login(client)
	chat(client)
	logout(client)
}

func login(client Client) {
	client.Write(runCommand("/greet"))
	client.Write(runCommand("/askUsername"))
	client.name = client.Read()
	addUser(client.name)
	fullUserGreeting := fmt.Sprintf("%s %s%s", userGreeting,
		client.name, userGreetingPunc)
	client.Write(fullUserGreeting)
}

func chat(client Client) {
	for {
		command := client.Read()

		if exitCommands[command] {
			break
		}

		if command == "" {
			continue
		} else if command[0] == '/' {
			client.Write(runCommand(command))
		} else {
			// Todo - Send out chat message
		}
	}
}

func logout(client Client) {
	client.Write(runCommand("/goodbye"))
}
