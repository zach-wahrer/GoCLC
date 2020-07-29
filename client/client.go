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

	"github.com/jroimartin/gocui"
)

type client struct {
	remote net.Conn
	input  io.Reader
	buf    *bytes.Buffer
}

// Start manages the lifecycle of a client.
func (c client) Start() {
	defer c.remote.Close()
	go c.runUI()
	// go c.receive()
	go c.send()
	c.chat()
	time.Sleep(5 * time.Millisecond)
	os.Exit(0)
}

func NewClient(address, port string) *client {
	return &client{connect(address, port), os.Stdin, new(bytes.Buffer)}
}

func (c client) chat() {
	for {

		// if key == keyboard.KeyCtrlC || (key == keyboard.KeyEnter && c.leaveChat(c.buf.String())) {
		// 	c.buf.WriteRune('\n')
		// 	c.channel <- c.buf.Bytes()
		// 	break
		// }

		// switch key {
		// case keyboard.KeyEnter:
		// 	c.buf.WriteRune('\n')
		// 	c.channel <- c.buf.Bytes()
		// 	c.buf.Reset()
		// case keyboard.KeyBackspace, keyboard.KeyBackspace2:
		// 	count := utf8.RuneCountInString(c.buf.String())
		// 	if count > 0 {
		// 		c.buf.Truncate(count - 1)
		// 	}
		// case keyboard.KeySpace:
		// 	c.buf.WriteRune(' ')
		// default:
		// 	c.buf.WriteRune(rune)
		// }

		// fmt.Printf("\u001b[2K\u001b[1000D>%s", c.buf.String())
	}
}

func (c client) runUI() {
	ui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer ui.Close()

	ui.SetManagerFunc(layout)

	if err := ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("users", 0, 0, maxX/6, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "GoCLC\n_____________")
	}
	if v, err := g.SetView("chat", maxX/6+1, 0, maxX-1, maxY-maxY/6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Chat message")
	}
	if v, err := g.SetView("input", maxX/6+1, maxY-maxY/6+1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "type here")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (c client) receive() {
	server := bufio.NewScanner(c.remote)
	for server.Scan() {
		c.printFromServer(server.Text())
	}
}

func (c client) send() {
	for {
		// buffer := <-c.channel
		// _, err := c.remote.Write(buffer)
		// if err != nil {
		// 	log.Print(err)
		// }
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
