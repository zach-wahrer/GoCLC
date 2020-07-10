package main

import (
	"fmt"
	"goclctest"
	"os"
	"testing"
)

func TestCommandLineFlags(t *testing.T) {
	os.Args = append(os.Args, fmt.Sprintf("-address=%s", goclctest.Address))
	os.Args = append(os.Args, fmt.Sprintf("-port=%s", goclctest.Port))
	go main()
	conn := goclctest.CreateTestConnection(t)
	goclctest.SendInputToServer(t, conn, "TestUserName")
	goclctest.SendInputToServer(t, conn, "/exit")
}
