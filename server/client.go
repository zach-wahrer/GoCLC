package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Client struct {
	c         net.Conn
	recieve   *bufio.Scanner
	broadcast chan string
	name      string
}

func (client Client) Write(input string) {
	if _, err := io.WriteString(client.c, input); err != nil {
		log.Print(err)
	}
}

func (client Client) Read() string {
	client.recieve.Scan()
	return client.recieve.Text()
}

func (client Client) Broadcast(message string) {
	client.broadcast <- fmt.Sprintf("%s: %s", client.name, message)
}
