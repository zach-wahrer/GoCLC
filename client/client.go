// Package client provides a client for the GoCLC chat service.
package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"server"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
)

type client struct {
	remote   net.Conn
	ui       *gocui.Gui
}

func NewClient(address, port string) *client {
	ui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	return &client{connect(address, port), ui}
}

// Start manages the lifecycle of a client.
func (c client) Start() {
	defer c.remote.Close()
	defer c.ui.Close()

	c.createUILayout()
	c.createKeybindings()

	if err := c.ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (c client) createUILayout() {
	c.ui.SetManagerFunc(c.layout)
}

func (c client) createKeybindings() {
	if err := c.ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, c.quit); err != nil {
		log.Panicln(err)
	}
	if err := c.ui.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, c.sendInput); err != nil {
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
		v.Editable = true
		v.Autoscroll = true
		v.Wrap = true
		ui.SetCurrentView("input")

	}
	return nil
}

func (c client) quit(ui *gocui.Gui, v *gocui.View) error {
	if _, err := c.remote.Write([]byte("/quit\n")); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	return gocui.ErrQuit
}

func (c client) receive(ui *gocui.Gui, v *gocui.View) {
	server := bufio.NewScanner(c.remote)
	for server.Scan() {
		fmt.Fprintln(v, server.Text())
		ui.Update(func(ui *gocui.Gui) error {
			return c.updateView("receive")
		})
	}
}

func (c client) sendInput(ui *gocui.Gui, v *gocui.View) error {
	if c.leaveChat(v.Buffer()) {
		return c.quit(ui, v)
	}
	if _, err := c.remote.Write([]byte(v.Buffer())); err != nil {
		return err
	}
	v.Clear()
	v.SetCursor(0, 0)
	return nil
}

func (c client) updateView(viewName string) error {
	_, err := c.ui.View(viewName)
	if err != nil {
		return err
	}
	return nil
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
