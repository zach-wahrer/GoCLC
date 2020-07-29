package client

import (
	"bufio"
	"bytes"
	"goclctest"
	"io"
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

// func TestAdvancedClientReceive(t *testing.T) {
// 	out := captureStdout(func() {
// 		c := NewClient(goclctest.Address, goclctest.Port)
// 		c.Start()
// 	})
// 	if !strings.Contains(out, server.ServerGreeting) {
// 		t.Errorf("client received unexpected reply - want: %s, got: %s", server.ServerGreeting, out)
// 	}
// 	if !strings.Contains(out, server.AskUsername) {
// 		t.Errorf("client received unexpected reply - want: %s, got: %s", server.AskUsername, out)
// 	}
// }
//
// func TestClientSend(t *testing.T) {
// 	out := captureStdout(func() {
// 		c := NewClient(goclctest.Address, goclctest.Port)
// 		c.input = bytes.NewBufferString(goclctest.TestUsername + "\n" + "/exit")
// 		go c.Start()
// 	})
// 	if !strings.Contains(out, server.UserGreeting) {
// 		t.Errorf("client received unexpected reply - want: %s, got: %s", server.UserGreeting+" "+goclctest.TestUsername, out)
// 	}
// }

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	go f()
	time.Sleep(5 * time.Millisecond)
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC
	return out
}
