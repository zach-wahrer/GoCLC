package server

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
)

const testAddress = "localhost"
const testPort = "8000"

func TestMain(m *testing.M) {
	go Listen(testAddress, testPort)
	code := m.Run()
	os.Exit(code)
}

func TestConnectionAndServerResponse(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	reply := new(bytes.Buffer)
	if _, err := io.Copy(reply, conn); err != nil {
		t.Errorf("unexpected server error: %v", err)
	}
	want := serverGreeting
	if reply.String() != want {
		t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", want, reply)
	}

}
