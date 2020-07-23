// Package client provides a client for the GoCLC chat service.
package client

import (
	"bufio"
	"log"
	"net"
)

func StartClient(address, port string) {
	conn := connect(address, port)
	defer conn.Close()
	go receiver(conn)
}

func connect(address, port string) net.Conn {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func receiver(conn net.Conn) bufio.Scanner {
	return *bufio.NewScanner(conn)
}
