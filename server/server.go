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
	send := make(chan string)
	recieve := make(chan string)
	go startBroadcaster(send, recieve)

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
		go handleConn(conn, recieve)
	}
}

func handleConn(conn net.Conn, broadcast chan string) {
	defer conn.Close()
	var client = Client{c: conn, recieve: bufio.NewScanner(conn), broadcast: broadcast}
	client.name = login(client)
	chat(client)
	logout(client)
}

func login(client Client) string {
	client.Write(runCommand("/greet"))
	client.Write(runCommand("/askUsername"))
	name := client.Read()
	enrichedUserGreeting := fmt.Sprintf("%s %s%s", userGreeting,
		name, userGreetingPunc)
	client.Write(enrichedUserGreeting)
	return name
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
			client.Broadcast(command)
		}
	}
}

func logout(client Client) {
	client.Write(runCommand("/goodbye"))
}
