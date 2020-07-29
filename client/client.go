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

	"github.com/jroimartin/gocui"
)

type client struct {
	remote net.Conn
	input  io.Reader
	buf    *bytes.Buffer
}

func NewClient(address, port string) *client {
	return &client{connect(address, port), os.Stdin, new(bytes.Buffer)}
}

// Start manages the lifecycle of a client.
func (c client) Start() {
	defer c.remote.Close()

	ui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer ui.Close()

	ui.SetManagerFunc(c.layout)

	if err := ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (c client) layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()
	if v, err := ui.SetView("users", 0, 0, maxX/6, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "GoCLC\n_____________")
	}
	if v, err := ui.SetView("receive", maxX/6+1, 0, maxX-1, maxY-maxY/6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = true
		v.Wrap = true
		go c.receive(ui, v)
	}
	if v, err := ui.SetView("input", maxX/6+1, maxY-maxY/6+1, maxX-1, maxY-1); err != nil {
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

func (c client) receive(ui *gocui.Gui, v *gocui.View) {
	server := bufio.NewScanner(c.remote)
	for server.Scan() {
		fmt.Fprintln(v, server.Text())
		ui.Update(func(ui *gocui.Gui) error {
			return c.updateView(ui, "receive")
		})
	}
}

func (c client) updateView(ui *gocui.Gui, viewName string) error {

	_, err := ui.View(viewName)
	if err != nil {
		return err
	}
	return nil
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
