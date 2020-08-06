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
	if !strings.Contains(receive.Text()+"\n", ServerTag) {
		goclctest.UnexpectedServerReplyError(t, ServerTag, receive.Text())
	}
	if !strings.Contains(receive.Text()+"\n", ServerGoodbye) {
		goclctest.UnexpectedServerReplyError(t, ServerGoodbye, receive.Text())
	}
}

func TestServerResponseForHelp(t *testing.T) {
	conn, receive := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "/help\n")

	helpLines := len(strings.Split(HelpMessage, "\n"))
	for i := 0; i < helpLines-1; i++ {
		receive.Scan()
		if !strings.Contains(HelpMessage, receive.Text()) {
			goclctest.UnexpectedServerReplyError(t, HelpMessage, receive.Text())
		}
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
	if !strings.Contains(receive.Text()+"\n", ServerWelcome) {
		goclctest.UnexpectedServerReplyError(t, ServerWelcome, receive.Text())
	}

	receive.Scan()
	if !strings.Contains(receive.Text()+"\n", AskUsername) {
		goclctest.UnexpectedServerReplyError(t, AskUsername, receive.Text())
	}

	goclctest.SendInputToServer(t, conn, TestUsername+"\n")
	receive.Scan()
	wants := []string{ServerTag, UserGreeting, TestUsername, UserGreetingPunc}
	for _, want := range wants {
		if !strings.Contains(receive.Text()+"\n", want) {
			goclctest.UnexpectedServerReplyError(t, want, receive.Text())
		}
	}

	receive.Scan()
	wants = []string{ServerTag, TestUsername, UserAnouncement}
	for _, want := range wants {
		if !strings.Contains(receive.Text()+"\n", want) {
			goclctest.UnexpectedServerReplyError(t, want, receive.Text())
		}
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

func TestServerLoggingClientLogin(t *testing.T) {
	var got bytes.Buffer
	log.SetOutput(&got)

	conn, _ := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn.Close()
	goclctest.SendInputToServer(t, conn, "/exit\n")
	time.Sleep(5 * time.Millisecond)
	log.SetOutput(os.Stderr)

	want := fmt.Sprintf("Client login: %s", TestUsername)

	if !strings.Contains(got.String(), want) {
		t.Errorf("server didn't log user login - want %s, got: %s", want, got.String())
	}
}

func TestServerLoggingClientLogout(t *testing.T) {
	conn, _ := goclctest.ReadyTestConnection(t, TestUsername)
	defer conn.Close()

	var got bytes.Buffer
	log.SetOutput(&got)

	goclctest.SendInputToServer(t, conn, "/exit\n")
	time.Sleep(5 * time.Millisecond)
	log.SetOutput(os.Stderr)

	want := fmt.Sprintf("Client logout: %s", TestUsername)

	if !strings.Contains(got.String(), want) {
		t.Errorf("server didn't log user logout- want %s, got: %s", want, got.String())
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

	if !strings.Contains(receive.Text()+"\n", DuplicateUsername) {
		goclctest.UnexpectedServerReplyError(t, DuplicateUsername, receive.Text())
	}
}

func TestNoProhibitedCharactersInUsername(t *testing.T) {
	b := newBroadcaster()
	if err := validateUsername(prohibitedUsernameCharacters[0], b); err == nil {
		t.Errorf("server allowed prohibitied chars in username")
	}
}

func TestNoServerTagInUsername(t *testing.T) {
	b := newBroadcaster()
	if err := validateUsername(ServerTag[1:len(ServerTag)-1], b); err == nil {
		t.Errorf("server allowed reserved username")
	}
}
