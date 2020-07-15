package server

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestBroadcastRecieve(t *testing.T) {
	var got bytes.Buffer
	log.SetOutput(&got)

	send := make(chan string)
	recieve := make(chan string)
	go startBroadcaster(send, recieve)

	testMessage := "Test message"
	recieve <- testMessage

	log.SetOutput(os.Stderr)

	if !strings.Contains(got.String(), testMessage) {
		t.Errorf("broadcast error - want: %s, got: %s", testMessage, got.String())
	}
}
