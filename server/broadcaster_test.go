package server

import (
	"goclctest"
	"strings"
	"testing"
)

func TestBroadcastSendReceive(t *testing.T) {

	conn, recieve := goclctest.CreateServerFixture(t)
	defer conn.Close()

	testMessage := "Test message"
	goclctest.SendInputToServer(t, conn, testMessage+"\n")
	recieve.Scan()

	if !strings.Contains(recieve.Text(), testMessage) {
		t.Errorf("broadcast error - want: %s, got: %s", testMessage, recieve.Text())
	}

	goclctest.SendInputToServer(t, conn, "/exit\n")
}
