package server

import (
	"bufio"
	"io"
	"log"
	"net"
)

type Client struct {
	c       net.Conn
	recieve *bufio.Scanner
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
