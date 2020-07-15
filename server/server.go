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
	login(&client)
	chat(client)
	logout(client)
}

func login(client *Client) {
	client.Write(runCommand("/greet"))
	client.Write(runCommand("/askUsername"))
	client.name = client.Read()
	enrichedUserGreeting := fmt.Sprintf("%s %s%s", userGreeting,
		client.name, userGreetingPunc)
	client.Write(enrichedUserGreeting)
}

func chat(client Client) {
	for {
		input := client.Read()

		if exitCommands[input] {
			break
		}

		if input == "" {
			continue
		} else if input[0] == '/' {
			client.Write(runCommand(input))
		} else {
			client.Broadcast(input)
		}
	}
}

func logout(client Client) {
	client.Write(runCommand("/goodbye"))
}
