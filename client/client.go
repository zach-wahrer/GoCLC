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
	"unicode/utf8"

	"github.com/eiannone/keyboard"
)

type client struct {
	remote  net.Conn
	input   io.Reader
	buf     *bytes.Buffer
	channel chan []byte
}

// Start manages the lifecycle of a client.
func (c client) Start() {
	defer c.remote.Close()
	go c.receive()
	go c.send()
	c.chat()
	time.Sleep(1 * time.Millisecond)
	os.Exit(0)
}

func NewClient(address, port string) *client {
	return &client{connect(address, port), os.Stdin, new(bytes.Buffer), make(chan []byte)}
}

func (c client) chat() {
	for {
		rune, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyCtrlC || (key == keyboard.KeyEnter && c.leaveChat(c.buf.String())) {
			break
		}

		switch key {
		case keyboard.KeyEnter:
			c.buf.WriteRune('\n')
			c.channel <- c.buf.Bytes()
			c.buf.Reset()
		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			count := utf8.RuneCountInString(c.buf.String())
			if count > 0 {
				c.buf.Truncate(count - 1)
			}
		case keyboard.KeySpace:
			c.buf.WriteRune(' ')
		default:
			c.buf.WriteRune(rune)
		}

		fmt.Printf("\u001b[2K\u001b[1000D>%s", c.buf.String())
	}
}

func (c client) receive() {
	server := bufio.NewScanner(c.remote)
	for server.Scan() {
		c.printFromServer(server.Text())
	}
}

func (c client) send() {
	for {
		buffer := <-c.channel
		_, err := c.remote.Write(buffer)
		if err != nil {
			log.Print(err)
		}
	}

}

func (c client) printFromServer(message string) {
	fmt.Print("\u001b[2K\u001b[1000D")
	fmt.Println(message)
	fmt.Print(c.buf.String())
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
