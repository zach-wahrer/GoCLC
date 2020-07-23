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
	conn   net.Conn
	reader io.Reader
}

// Start creates a new client connection and runs client functions
func (c client) Start() {
	defer c.conn.Close()
	go c.receive()

	for {
		if _, err := io.Copy(c.conn, c.reader); err != nil {
			log.Print(err)
		}
	}
}

func NewClient(address, port string) *client {
	return &client{connect(address, port), os.Stdin}
}

func (c client) receive() {
	server := bufio.NewScanner(c.conn)
	for server.Scan() {
		fmt.Println(server.Text())
	}
}

func (c client) send(message string) {
	if _, err := io.WriteString(c.conn, message); err != nil {
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
