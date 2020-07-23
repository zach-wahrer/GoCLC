package client

import (
	"bufio"
	"goclctest"
	"os"
	"server"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go server.Listen(goclctest.Address, goclctest.Port)
	time.Sleep(5 * time.Millisecond)
	code := m.Run()
	os.Exit(code)
}

func TestBasicClientConnection(t *testing.T) {
	client := connect(goclctest.Address, goclctest.Port)
	recieve := bufio.NewScanner(client)
	recieve.Scan()
	if recieve.Text()+"\n" != server.ServerGreeting {
		t.Errorf("client recieved unexpected reply - want: %s, got: %s", server.ServerGreeting, recieve.Text())
	}
	client.Close()
}
