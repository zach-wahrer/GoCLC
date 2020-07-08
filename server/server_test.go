package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

const testAddress = "localhost"
const testPort = "8000"

func TestMain(m *testing.M) {
	go Listen(testAddress, testPort)
	time.Sleep(5 * time.Millisecond)
	code := m.Run()
	os.Exit(code)
}

func TestConnectionAndServerResponse(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	reply := bufio.NewScanner(conn)

	reply.Scan()
	want := serverGreeting
	if reply.Text()+"\n" != want {
		t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", want, reply.Text())
	}

	if _, err := io.WriteString(conn, "/exit\n"); err != nil {
		t.Errorf("unexpected server error: %v", err)
	}

	reply.Scan()
	want = serverGoodbye
	if reply.Text()+"\n" != want {
		t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", want, reply.Text())
	}

}

func TestServerResponseForHelp(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	reply := bufio.NewScanner(conn)
	reply.Scan() // Skip welcome message

	if _, err := io.WriteString(conn, "/help\n"); err != nil {
		t.Errorf("unexpected server error: %v", err)
	}

	helpLines := len(strings.Split(helpMessage, "\n"))
	combinedReply := ""
	for i := 0; i < helpLines-1; i++ {
		reply.Scan()
		combinedReply += reply.Text() + "\n"
	}

	want := helpMessage
	if combinedReply != want {
		t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", want, reply.Text())
	}
	if _, err := io.WriteString(conn, "/exit\n"); err != nil {
		t.Errorf("unexpected server error: %v", err)
	}

}
