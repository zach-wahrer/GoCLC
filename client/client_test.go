package client

import (
	"bufio"

	"goclctest"
	"os"
	"server"
	"strings"
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
	conn := connect(goclctest.Address, goclctest.Port)
	conn.Close()
}

func TestBasicClientReceive(t *testing.T) {
	conn := connect(goclctest.Address, goclctest.Port)
	receive := bufio.NewScanner(conn)
	receive.Scan()
	if !strings.Contains(receive.Text()+"\n", server.ServerGreeting) {
		t.Errorf("client received unexpected reply - want: %s, got: %s", server.ServerGreeting, receive.Text())
	}
	conn.Close()
}
