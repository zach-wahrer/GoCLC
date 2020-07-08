package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const serverGreeting = "Welcome to the GoCLC Server!\n"
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

func handleConn(c net.Conn) {
	defer c.Close()
	reply := runCommand("/greet")
	if _, err := io.WriteString(c, reply); err != nil {
		log.Print(err)
	}
	input := bufio.NewScanner(c)
	for input.Scan() {
		command := input.Text()
		if command[0] == '/' {
			if command == "/quit" || command == "/exit" || command == "/q" {
				break
			} else {
				if _, err := io.WriteString(c, runCommand(command)); err != nil {
					log.Print(err)
				}
			}
		}
	}

	if _, err := io.WriteString(c, serverGoodbye); err != nil {
		log.Print(err)
	}
}
