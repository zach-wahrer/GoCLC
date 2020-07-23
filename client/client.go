// Package client provides a client for the GoCLC chat service.
package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	conn net.Conn
}

// Start creates a new client connection and runs client functions
func Start(address, port string) {
	c := newClient(address, port)
	go c.receive()
	for {

	}
}

func (c client) receive() {
	server := bufio.NewScanner(c.conn)
	for server.Scan() {
		fmt.Println(server.Text())
	}
}

func newClient(address, port string) *client {
	return &client{connect(address, port)}
}

func connect(address, port string) net.Conn {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
