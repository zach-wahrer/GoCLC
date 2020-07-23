// Package client provides a client for the GoCLC chat service.
package client

import (
	"log"
	"net"
)

func connect(address, port string) net.Conn {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
