package server

import (
	"errors"
	"goclctest"
	"testing"
)

func TestClientAddedToBroadcaster(t *testing.T) {
	b := NewBroadcaster()
	b.addClient(Client{name: "TestClient"})
	if client := b.clients["TestClient"]; client.name != "TestClient" {
		goclctest.InternalServerError(t, errors.New("client not sucessfully added to broadcaster"))
	}
}
