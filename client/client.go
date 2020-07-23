// Package client provides a client for the GoCLC chat service.
package client

import (
	"log"
	"net"
)

// Connect creates a network connection to the specified address and port.
func Connect(address, port string) net.Conn {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
