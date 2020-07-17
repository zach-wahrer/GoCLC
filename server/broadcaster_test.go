package server

import (
	"errors"
	"goclctest"
	"strings"
	"testing"
)

var errClientNotAdded = errors.New("client not sucessfully added to broadcaster")

func TestBroadcastSendReceive(t *testing.T) {

	conn, receive := goclctest.ReadyTestConnection(t)
	defer conn.Close()

	testMessage := "Test message"
	goclctest.SendInputToServer(t, conn, testMessage+"\n")
	receive.Scan()

	if !strings.Contains(receive.Text(), testMessage) {
		t.Errorf("broadcast error - want: %s, got: %s", testMessage, receive.Text())
	}

	goclctest.SendInputToServer(t, conn, "/exit\n")
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
