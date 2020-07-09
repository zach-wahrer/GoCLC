package server

import (
	"bufio"
	"net"
)

type Client struct {
	c       net.Conn
	recieve *bufio.Scanner
}
