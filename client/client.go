// Package client provides a client for the GoCLC chat service.
package client

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"server"
	"strings"
	"time"
)

type client struct {
	remote net.Conn
	input  io.Reader
	buf    *bytes.Buffer
}

// Start manages the lifecycle of a client.
func (c client) Start() {
	defer c.remote.Close()
	go c.receive()
	c.chat()
	os.Exit(0)
}

func NewClient(address, port string) *client {
	return &client{connect(address, port), os.Stdin, new(bytes.Buffer)}
}

func (c client) chat() {
	reader := bufio.NewReader(c.input)
	for {
		rune, _, err := reader.ReadRune()
		if err != nil {
			log.Print(err)
		}
		c.buf.WriteRune(rune)
		if rune == '\n' {
			if c.leaveChat(c.buf.String()) {
				time.Sleep(5 * time.Millisecond)
				break
			}
			c.send(c.buf.String())
			c.buf.Reset()
		}

	}
}

func (c client) receive() {
	server := bufio.NewScanner(c.remote)
	for server.Scan() {
		fmt.Println(server.Text())
	}
}

func (c client) send(message string) {
	if _, err := io.WriteString(c.remote, message); err != nil {
		log.Print(err)
	}
}

func (c client) leaveChat(input string) bool {
	return server.ExitCommands[strings.TrimSuffix(input, "\n")]
}

func connect(address, port string) net.Conn {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
