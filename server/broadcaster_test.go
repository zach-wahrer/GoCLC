package server

import (
	"bytes"
	"errors"
	"goclctest"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

var errClientNotAdded = errors.New("client not sucessfully added to broadcaster")

func TestBroadcastSendReceive(t *testing.T) {
	conn, receive := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn.Close()

	testMessage := "Test message"
	goclctest.SendInputToServer(t, conn, testMessage+"\n")
	receive.Scan()

	if !strings.Contains(receive.Text(), testMessage) {
		t.Errorf("broadcast error - want: %s, got: %s", testMessage, receive.Text())
	}
	goclctest.SendInputToServer(t, conn, "/exit\n")
}

func TestUnexpectedDisconnectedClient(t *testing.T) {
	conn1, _ := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn1.Close()
	conn2, _ := goclctest.ReadyTestConnection(t, "AnotherTestuser")
	conn2.Close()

	var got bytes.Buffer
	log.SetOutput(&got)
	goclctest.SendInputToServer(t, conn1, "Test message\n")
	goclctest.SendInputToServer(t, conn1, "Another test message\n")
	goclctest.SendInputToServer(t, conn1, "/exit\n")
	time.Sleep(5 * time.Millisecond)
	log.SetOutput(os.Stderr)

	err := "disconnected unexpectedly"
	if !strings.Contains(got.String(), err) {
		t.Errorf("server didn't disconnect user - want: %s, got: %s", err, got.String())
	}
	if strings.Contains(got.String(), "write: broken pipe") {
		t.Errorf("server wrote to broken pipe - want: %s, got: %s", err, got.String())
	}

}

func TestClientAddedToBroadcaster(t *testing.T) {
	b := newBroadcaster()
	b.addClient(&Client{name: "TestClient"})
	if _, ok := b.clients["TestClient"]; !ok {
		goclctest.InternalServerError(t, errClientNotAdded)
	}
}

func TestClientRemovedFromBroadcaster(t *testing.T) {
	b := newBroadcaster()
	client := Client{name: "TestClient"}
	b.addClient(&client)
	b.removeClient(&client)
	if _, ok := b.clients["TestClient"]; ok {
		goclctest.InternalServerError(t, errors.New("client not removed from broadcaster"))
	}
}

func TestAttemptDuplicateUsernameForBroadcaster(t *testing.T) {
	b := newBroadcaster()
	if ok := b.addClient(&Client{name: "TestClient"}); !ok {
		goclctest.InternalServerError(t, errClientNotAdded)
	}
	if ok := b.addClient(&Client{name: "AnotherTester"}); !ok {
		goclctest.InternalServerError(t, errClientNotAdded)
	}
	if ok := b.addClient(&Client{name: "TestClient"}); ok {
		goclctest.InternalServerError(t, errors.New("duplicate client added to broadcaster"))
	}
}

func TestASCIIArtGeneration(t *testing.T) {
	want := "| |_ ___ ___| |_\n|  _| -_|_ -|  _|\n|_| |___|___|_|"
	got := createASCIIArt("rectangles", "test")
	for _, line := range strings.Split(want, "\n") {
		if !strings.Contains(got, line) {
			goclctest.InternalServerError(t, errors.New("ascii art not properly generated"))
		}
	}
}

func TestASCIIArtGenerationFailure(t *testing.T) {
	want := ""
	got := createASCIIArt("1", "test")
	if got != want {
		goclctest.InternalServerError(t, errors.New("ascii art returned a string when it should be empty"))
	}
}
