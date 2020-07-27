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
}

// Start manages the lifecycle of a client.
func (c client) Start() {
	defer c.remote.Close()
	go c.receive()
	c.chat()
	os.Exit(0)
}

func NewClient(address, port string) *client {
	return &client{connect(address, port), os.Stdin}
}

func (c client) chat() {
	reader := bufio.NewReader(c.input)
	input := new(bytes.Buffer)
	for {
		rune, _, err := reader.ReadRune()
		if err != nil {
			log.Print(err)
		}
		input.WriteRune(rune)
		if rune == '\n' {
			if c.leaveChat(input.String()) {
				time.Sleep(5 * time.Millisecond)
				break
			}
			c.send(input.String())
			input.Reset()
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
