// Package client provides a client for the GoCLC chat service.
package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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

}

func NewClient(address, port string) *client {
	return &client{connect(address, port), os.Stdin}
}

func (c client) chat() {
	for {
		if _, err := io.Copy(c.remote, c.input); err != nil {
			log.Print(err)
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
