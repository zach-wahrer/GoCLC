package server

import (
	"errors"
	"goclctest"
	"strings"
	"testing"
)

func TestBroadcastSendReceive(t *testing.T) {

	conn, receive := goclctest.CreateServerFixture(t)
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
		goclctest.InternalServerError(t, errors.New("client not sucessfully added to broadcaster"))
	}
}
