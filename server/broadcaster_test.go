package server

import (
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
