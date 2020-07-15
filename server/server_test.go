package server

import (
	"bufio"
	"bytes"
	"fmt"
	"goclctest"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

const testMessage = "Test message\n"
const TestUsername = goclctest.TestUsername

func TestMain(m *testing.M) {
	go Listen(goclctest.Address, goclctest.Port)
	time.Sleep(5 * time.Millisecond)
	code := m.Run()
	os.Exit(code)
}

func TestConnectionAndServerResponse(t *testing.T) {
	conn, recieve := goclctest.CreateServerFixture(t)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "/exit\n")
	recieve.Scan()
	if recieve.Text()+"\n" != serverGoodbye {
		goclctest.UnexpectedServerReplyError(t, serverGoodbye, recieve.Text())
	}

}

func TestServerResponseForHelp(t *testing.T) {
	conn, recieve := goclctest.CreateServerFixture(t)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "/help\n")

	helpLines := len(strings.Split(helpMessage, "\n"))
	combinedReply := ""
	for i := 0; i < helpLines-1; i++ {
		recieve.Scan()
		combinedReply += recieve.Text() + "\n"
	}
	if combinedReply != helpMessage {
		goclctest.UnexpectedServerReplyError(t, helpMessage, combinedReply)
	}

	goclctest.SendInputToServer(t, conn, "/exit\n")

}

func TestServerWithEmptyInput(t *testing.T) {
	conn, _ := goclctest.CreateServerFixture(t)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "\n")
	goclctest.SendInputToServer(t, conn, "/exit\n")
}

func TestServerFixture(t *testing.T) {
	conn := goclctest.CreateTestConnection(t)
	recieve := bufio.NewScanner(conn)

	recieve.Scan()
	if recieve.Text()+"\n" != serverGreeting {
		goclctest.UnexpectedServerReplyError(t, serverGreeting, recieve.Text())
	}

	recieve.Scan()
	if recieve.Text()+"\n" != askUsername {
		goclctest.UnexpectedServerReplyError(t, askUsername, recieve.Text())
	}

	goclctest.SendInputToServer(t, conn, TestUsername+"\n")
	recieve.Scan()
	want := fmt.Sprintf("%s %s%s", userGreeting, TestUsername, userGreetingPunc)
	if recieve.Text()+"\n" != want {
		goclctest.UnexpectedServerReplyError(t, want, recieve.Text())
	}

	goclctest.SendInputToServer(t, conn, "/exit\n")

}

func TestServerLogging(t *testing.T) {
	conn, _ := goclctest.CreateServerFixture(t)
	defer conn.Close()

	var got bytes.Buffer
	log.SetOutput(&got)
	goclctest.SendInputToServer(t, conn, testMessage)
	goclctest.SendInputToServer(t, conn, "/exit\n")
	time.Sleep(5 * time.Millisecond)
	log.SetOutput(os.Stderr)

	if !strings.Contains(got.String(), testMessage) {
		t.Errorf("server didn't log message - want: %s, got: %s", testMessage, got.String())
	}
	if !strings.Contains(got.String(), TestUsername) {
		t.Errorf("server didn't log username - want %s, got: %s", TestUsername, got.String())
	}

}
