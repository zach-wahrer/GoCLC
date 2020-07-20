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
	conn, receive := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "/exit\n")
	receive.Scan()
	if receive.Text()+"\n" != serverGoodbye {
		goclctest.UnexpectedServerReplyError(t, serverGoodbye, receive.Text())
	}
}

func TestServerResponseForHelp(t *testing.T) {
	conn, receive := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "/help\n")

	helpLines := len(strings.Split(helpMessage, "\n"))
	combinedReply := ""
	for i := 0; i < helpLines-1; i++ {
		receive.Scan()
		combinedReply += receive.Text() + "\n"
	}
	if combinedReply != helpMessage {
		goclctest.UnexpectedServerReplyError(t, helpMessage, combinedReply)
	}

	goclctest.SendInputToServer(t, conn, "/exit\n")
}

func TestServerWithEmptyInput(t *testing.T) {
	conn, _ := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "\n")
	goclctest.SendInputToServer(t, conn, "/exit\n")
}

func TestServerFixture(t *testing.T) {
	conn := goclctest.CreateTestConnection(t)
	receive := bufio.NewScanner(conn)

	receive.Scan()
	if receive.Text()+"\n" != serverGreeting {
		goclctest.UnexpectedServerReplyError(t, serverGreeting, receive.Text())
	}

	receive.Scan()
	if receive.Text()+"\n" != askUsername {
		goclctest.UnexpectedServerReplyError(t, askUsername, receive.Text())
	}

	goclctest.SendInputToServer(t, conn, TestUsername+"\n")
	receive.Scan()
	want := fmt.Sprintf("%s %s%s", userGreeting, TestUsername, userGreetingPunc)
	if receive.Text()+"\n" != want {
		goclctest.UnexpectedServerReplyError(t, want, receive.Text())
	}

	goclctest.SendInputToServer(t, conn, "/exit\n")
}

func TestServerLoggingClientMessages(t *testing.T) {
	conn, _ := goclctest.ReadyTestConnection(t, TestUsername)
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

func TestServerLoggingClientEntry(t *testing.T) {
	var got bytes.Buffer
	log.SetOutput(&got)

	conn, _ := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "/exit\n")
	time.Sleep(5 * time.Millisecond)
	log.SetOutput(os.Stderr)

	want := fmt.Sprintf("Client login: %s", TestUsername)

	if !strings.Contains(got.String(), want) {
		t.Errorf("server didn't log new user - want %s, got: %s", want, got.String())
	}
}

func TestNoDuplicateUsersAllowed(t *testing.T) {
	conn1, _ := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn1.Close()
	conn2 := goclctest.CreateTestConnection(t)
	defer conn2.Close()
	receive := bufio.NewScanner(conn2)

	receive.Scan() // Server Greeting
	receive.Scan() // Ask Username
	goclctest.SendInputToServer(t, conn2, TestUsername+"\n")
	receive.Scan()

	if receive.Text()+"\n" != duplicateUsername {
		goclctest.UnexpectedServerReplyError(t, duplicateUsername, receive.Text())
	}

}
