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
	color   string
}

func newClient(conn net.Conn, send chan string) Client {
	return Client{
		c:       conn,
		receive: bufio.NewScanner(conn),
		send:    send,
		address: conn.RemoteAddr().String(),
		color:   randomColor()}

}

func (client Client) broadcast(message string) {
	client.send <- client.wrapMessage(message)
}

func (client Client) read() string {
	client.receive.Scan()
	return client.receive.Text()
}

func (client Client) write(message string) error {
	if _, err := io.WriteString(client.c, message); err != nil {
		err := fmt.Errorf("%s disconnected unexpectedly", client.address)
		log.Print(err)
		return err
	}
	return nil
}

func (client Client) wrapMessage(message string) string {
	return fmt.Sprintf("%s<%s> %s%s", client.color, client.name, colorReset, message)
}
