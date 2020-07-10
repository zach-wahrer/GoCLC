// Package goclctest shares functions needed for GoCLC testing.
package goclctest

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"testing"
)

const Address = "localhost"
const Port = "9000"

func CreateServerFixture(t *testing.T) (net.Conn, bufio.Scanner) {
	conn := CreateTestConnection(t)
	recieve := bufio.NewScanner(conn)

	recieve.Scan() // Server Greeting
	recieve.Scan() // Ask Username
	SendInputToServer(t, conn, "TestUsername\n")
	recieve.Scan() // User Greeting
	return conn, *recieve
}

func CreateTestConnection(t *testing.T) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", Address, Port))
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func SendInputToServer(t *testing.T, conn net.Conn, input string) {
	if _, err := io.WriteString(conn, input); err != nil {
		UnexpectedServerError(t, err)
	}
}

func UnexpectedServerReplyError(t *testing.T, want, got string) {
	t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", want, got)
}

func UnexpectedServerError(t *testing.T, err error) {
	t.Errorf("unexpected server error: %v", err)
}
