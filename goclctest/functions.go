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
const TestUsername = "TestUsername"

func ReadyTestConnection(t *testing.T) (net.Conn, bufio.Scanner) {
	conn := CreateTestConnection(t)
	receive := bufio.NewScanner(conn)

	receive.Scan() // Server Greeting
	receive.Scan() // Ask Username
	SendInputToServer(t, conn, TestUsername+"\n")
	receive.Scan() // User Greeting
	return conn, *receive
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
		InternalServerError(t, err)
	}
}

func UnexpectedServerReplyError(t *testing.T, want, got string) {
	t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", want, got)
}

func InternalServerError(t *testing.T, err error) {
	t.Errorf("internal server error: %v", err)
}
