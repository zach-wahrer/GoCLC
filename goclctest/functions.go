// Package goclctest shares functions needed for GoCLC testing.
package goclctest

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"testing"
)

// Address is a testing constant for use in goclctest and goclcserver
const Address = "localhost"

// Port is a testing constant for use in goclctest and goclcserver
const Port = "9000"

// TestUsername is a testing constant for use in goclctest and goclcserver
const TestUsername = "TestUsername"

// ReadyTestConnection creates a test connection to server, then logs in.
func ReadyTestConnection(t *testing.T) (net.Conn, bufio.Scanner) {
	conn := CreateTestConnection(t)
	receive := bufio.NewScanner(conn)

	receive.Scan() // Server Greeting
	receive.Scan() // Ask Username
	SendInputToServer(t, conn, TestUsername+"\n")
	receive.Scan() // User Greeting
	return conn, *receive
}

// CreateTestConnection returns a test connection to the specified address/port.
func CreateTestConnection(t *testing.T) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", Address, Port))
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

// SendInputToServer writes a string to a connection.
func SendInputToServer(t *testing.T, conn net.Conn, input string) {
	if _, err := io.WriteString(conn, input); err != nil {
		InternalServerError(t, err)
	}
}

// UnexpectedServerReplyError creates a testing error.
func UnexpectedServerReplyError(t *testing.T, want, got string) {
	t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", want, got)
}

// InternalServerError creates a testing error.
func InternalServerError(t *testing.T, err error) {
	t.Errorf("internal server error: %v", err)
}
