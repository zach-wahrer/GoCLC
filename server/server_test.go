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
	conn := createTestConnection(t)
	defer conn.Close()

	reply := bufio.NewScanner(conn)

	reply.Scan()
	if reply.Text()+"\n" != serverGreeting {
		unexpectedServerReplyError(t, serverGreeting, reply.Text())
	}

	sendInputToServer(t, conn, "/exit\n")

	reply.Scan()
	if reply.Text()+"\n" != serverGoodbye {
		unexpectedServerReplyError(t, serverGoodbye, reply.Text())
	}

}

func TestServerResponseForHelp(t *testing.T) {
	conn := createTestConnection(t)
	defer conn.Close()

	reply := bufio.NewScanner(conn)
	reply.Scan() // Skip welcome message

	sendInputToServer(t, conn, "/help\n")

	helpLines := len(strings.Split(helpMessage, "\n"))
	combinedReply := ""
	for i := 0; i < helpLines-1; i++ {
		reply.Scan()
		combinedReply += reply.Text() + "\n"
	}

	if combinedReply != helpMessage {
		unexpectedServerReplyError(t, helpMessage, combinedReply)
	}

	sendInputToServer(t, conn, "/exit\n")

}

// func TestServerGetUsername(t *testing.T) {
// 	conn := createTesetConnection(t)
// 	defer conn.Close()
//
// 	reply := bufio.NewScanner(conn)
// 	reply.Scan() // Skip welcome message
//
// 	if _, err := io.WriteString(conn, "TestUser\n"); err != nil {
//
// 	}
// }

func sendInputToServer(t *testing.T, conn net.Conn, input string) {
	if _, err := io.WriteString(conn, input); err != nil {
		unexpectedServerError(t, err)
	}
}

func unexpectedServerReplyError(t *testing.T, want, got string) {
	t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", want, got)
}

func unexpectedServerError(t *testing.T, err error) {
	t.Errorf("unexpected server error: %v", err)
}

func createTestConnection(t *testing.T) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
	if err != nil {
		t.Fatal(err)
	}
	return conn
}
