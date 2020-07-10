// Package main starts a GoCLC server on the local machine
package main

import (
	"server"
)

func main() {
	server.Listen("localhost", "8000")
}
