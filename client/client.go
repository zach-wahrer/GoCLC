// Package client provides a client for the GoCLC chat service.
package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"server"
	"strings"
	"time"
)

type client struct {
	remote net.Conn
	input  io.Reader
}

// Start creates a new client connection and runs client functions
func (c client) Start() {
	defer c.remote.Close()
	go c.receive()
	c.chat()
	os.Exit(0)
}

func NewClient(address, port string) *client {
	return &client{connect(address, port), os.Stdin}
}

func (c client) chat() {
	reader := bufio.NewReader(c.input)
	for {
		input, _ := reader.ReadString('\n')
		if _, err := io.WriteString(c.remote, input); err != nil {
			log.Print(err)
		}
		if server.ExitCommands[strings.TrimSuffix(input, "\n")] {
			time.Sleep(5 * time.Millisecond)
			break
		}
	}
}

func (c client) receive() {
	server := bufio.NewScanner(c.remote)
	for server.Scan() {
		fmt.Println(server.Text())
	}
}

func (c client) send(message string) {
	if _, err := io.WriteString(c.remote, message); err != nil {
		log.Print(err)
	}
}

func connect(address, port string) net.Conn {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
