// Package server provides a stand-alone server for the GoCLC chat service.
package server

import (
	"fmt"
	"log"
	"net"
)

// Listen checks for chat connections on the given address and port.
func Listen(address, port string) {
	broadcaster := startBroadcaster()

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
		go handleConn(conn, broadcaster)
	}
}

func handleConn(conn net.Conn, broadcaster *Broadcaster) {
	defer conn.Close()
	client := newClient(conn, broadcaster.receive)
	login(&client, broadcaster)
	chat(client)
	logout(&client, broadcaster)
}

func login(client *Client, broadcaster *Broadcaster) {
	client.write(runCommand("/greet"))
	getUsername(client, broadcaster)

	enrichedUserGreeting := fmt.Sprintf("%s %s%s", userGreeting,
		client.name, userGreetingPunc)
	client.write(enrichedUserGreeting)

	log.Printf("Client login: %s - %s", client.name, client.address)

	broadcaster.sendToAllClients(fmt.Sprintf("%s %s", client.name, userAnouncement))
}

func chat(client Client) {
	for {
		input := client.read()

		if exitCommands[input] {
			logInput(client, input)
			break
		}

		if input == "" {
			continue
		}

		handleInput(client, input)
	}
}

func logout(client *Client, broadcaster *Broadcaster) {
	client.write(runCommand("/goodbye"))
	broadcaster.removeClient(client)
	log.Printf("Client logout: %s - %s", client.name, client.address)
}

func handleInput(client Client, input string) {
	if input[0] == '/' {
		client.write(runCommand(input))
	} else {
		client.broadcast(input + "\n")
	}
	logInput(client, input)
}

func getUsername(client *Client, broadcaster *Broadcaster) {
	client.write(runCommand("/askUsername"))
	client.name = client.read()
	for ok := broadcaster.addClient(client); !ok; ok = broadcaster.addClient(client) {
		client.write(runCommand("/duplicateUsername"))
		client.write(runCommand("/askUsername"))
		client.name = client.read()
	}
}

func logInput(client Client, input string) {
	log.Printf("%s - %s - %s", client.address, client.name, input)
}
