package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Client struct {
	c       net.Conn
	receive *bufio.Scanner
	send    chan string
	address string
	name    string
}

func NewClient(conn net.Conn, send chan string) Client {
	client := Client{
		c:       conn,
		receive: bufio.NewScanner(conn),
		send:    send,
		address: conn.RemoteAddr().String()}
	return client
}

func (client Client) Write(input string) {
	if _, err := io.WriteString(client.c, input); err != nil {
		log.Print(err)
	}
}

func (client Client) Read() string {
	client.receive.Scan()
	return client.receive.Text()
}

func (client Client) Broadcast(message string) {
	client.send <- fmt.Sprintf("%s: %s", client.name, message)
}
