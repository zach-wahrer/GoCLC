package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// Client manages all aspects of client connections.
type Client struct {
	c       net.Conn
	receive *bufio.Scanner
	send    chan string
	address string
	name    string
}

func newClient(conn net.Conn, send chan string) Client {
	return Client{
		c:       conn,
		receive: bufio.NewScanner(conn),
		send:    send,
		address: conn.RemoteAddr().String()}
}

func (client Client) broadcast(message string) {
	client.send <- fmt.Sprintf("%s: %s", client.name, message)
}

func (client Client) read() string {
	client.receive.Scan()
	return client.receive.Text()
}

func (client Client) write(message string) {
	if _, err := io.WriteString(client.c, message); err != nil {
		log.Print(err)
	}
}
