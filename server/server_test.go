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
	conn, recieve := createServerFixture(t)
	defer conn.Close()
	sendInputToServer(t, conn, "/exit\n")
	recieve.Scan()
	if recieve.Text()+"\n" != serverGoodbye {
		unexpectedServerReplyError(t, serverGoodbye, recieve.Text())
	}

}

func TestServerResponseForHelp(t *testing.T) {
	conn, recieve := createServerFixture(t)
	defer conn.Close()
	sendInputToServer(t, conn, "/help\n")

	helpLines := len(strings.Split(helpMessage, "\n"))
	combinedReply := ""
	for i := 0; i < helpLines-1; i++ {
		recieve.Scan()
		combinedReply += recieve.Text() + "\n"
	}
	if combinedReply != helpMessage {
		unexpectedServerReplyError(t, helpMessage, combinedReply)
	}

	sendInputToServer(t, conn, "/exit\n")

}

func TestServerFixture(t *testing.T) {
	conn := createTestConnection(t)
	recieve := bufio.NewScanner(conn)

	recieve.Scan()
	if recieve.Text()+"\n" != serverGreeting {
		unexpectedServerReplyError(t, serverGreeting, recieve.Text())
	}

	recieve.Scan()
	if recieve.Text()+"\n" != askUsername {
		unexpectedServerReplyError(t, askUsername, recieve.Text())
	}

	sendInputToServer(t, conn, "TestUsername\n")
	recieve.Scan()
	want := fmt.Sprintf("%s TestUsername%s", userGreeting, userGreetingPunc)
	if recieve.Text()+"\n" != want {
		unexpectedServerReplyError(t, want, recieve.Text())
	}

	sendInputToServer(t, conn, "/exit\n")

}

func createServerFixture(t *testing.T) (net.Conn, bufio.Scanner) {
	conn := createTestConnection(t)
	recieve := bufio.NewScanner(conn)

	recieve.Scan() // Server Greeting
	recieve.Scan() // Ask Username
	sendInputToServer(t, conn, "TestUsername\n")
	recieve.Scan() // User Greeting
	return conn, *recieve
}

func createTestConnection(t *testing.T) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

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
