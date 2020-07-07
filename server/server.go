package server

import (
	"fmt"
	"io"
	"log"
	"net"
)

const serverGreeting = "Welcome to the GoCLC Server!\n"

// Listen for incommimng connections
func Listen(address, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	reply := runCommand("greet")
	if _, err := io.WriteString(c, reply); err != nil {
		log.Print(err)
	}
}
