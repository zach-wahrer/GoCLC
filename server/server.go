package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

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
	login(client)
	chat(client)

	client.Write(runCommand("/goodbye"))
}

func login(client Client) {
	client.Write(runCommand("/greet"))
	client.Write(runCommand("/askUsername"))
	addUser(client.Read())
	client.Write(fmt.Sprintf("%s %s%s", userGreeting, client.recieve.Text(), userGreetingPunc))
}

func chat(client Client) {
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
}
