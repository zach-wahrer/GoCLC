// Package server provides a stand-alone server for the GoCLC chat service.
package server

import (
	"fmt"
	"log"
	"net"
	"strings"
)

var prohibitedUsernameCharacters = []string{"<", ">"}

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
	broadcaster.welcomeClient(client)
	getUsername(client, broadcaster)
	broadcaster.greetUserByName(client)
	broadcaster.announceNewClient(client)
	logArrival(client)
}

func chat(client Client) {
	for {
		input := client.read()

		if ExitCommands[input] {
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
	broadcaster.sayGoodbye(client)
	broadcaster.removeClient(client)
	broadcaster.announceDepartedClient(client)
	logDepature(client)
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
	broadcaster.askUsername(client)
	username := client.read()
	err := validateUsername(username, broadcaster)
	for err != nil {
		broadcaster.sendError(client, err)
		broadcaster.askUsername(client)
		username = client.read()
		err = validateUsername(username, broadcaster)
	}
	client.name = username
	broadcaster.addClient(client)
}

func validateUsername(username string, broadcaster *Broadcaster) error {
	if strings.EqualFold(username, ServerTag[1:len(ServerTag)-1]) {
		return fmt.Errorf("Username cannot contain: %s\n", ServerTag)
	}

	for _, char := range prohibitedUsernameCharacters {
		if strings.Contains(username, char) {
			return fmt.Errorf("Username cannot contain these characters: %s\n", prohibitedUsernameCharacters)
		}
	}
	if !broadcaster.usernameAvailable(username) {
		return fmt.Errorf(DuplicateUsername)
	}
	return nil
}

func logInput(client Client, input string) {
	log.Printf("%s - %s - %s", client.address, client.name, input)
}

func logArrival(client *Client) {
	log.Printf("Client login: %s - %s", client.name, client.address)
}

func logDepature(client *Client) {
	log.Printf("Client logout: %s - %s", client.name, client.address)
}
