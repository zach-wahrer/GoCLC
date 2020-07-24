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
	client.write(runCommand("/greet"))
	getUsername(client, broadcaster)

	enrichedUserGreeting := fmt.Sprintf("%s %s %s%s", ServerTag, UserGreeting,
		client.name, UserGreetingPunc)
	client.write(enrichedUserGreeting)

	log.Printf("Client login: %s - %s", client.name, client.address)

	broadcaster.sendToAllClients(fmt.Sprintf("%s %s %s", ServerTag, client.name, UserAnouncement))
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
	client.write(runCommand("/goodbye"))
	broadcaster.removeClient(client)
	broadcaster.sendToAllClients(fmt.Sprintf("%s %s %s", ServerTag, client.name, UserDepartedAnnouncement))
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
	client.write(runCommand("/AskUsername"))
	username := client.read()
	err := validateUsername(username, broadcaster)
	for err != nil {
		client.write(fmt.Sprintf("%s", err))
		client.write(runCommand("/AskUsername"))
		username = client.read()
		err = validateUsername(username, broadcaster)
	}
	client.name = username
	broadcaster.addClient(client)
}

func logInput(client Client, input string) {
	log.Printf("%s - %s - %s", client.address, client.name, input)
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
