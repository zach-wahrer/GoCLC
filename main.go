package main

import (
	"server"
)

func main() {
	server.Listen("localhost", "8000")
}
