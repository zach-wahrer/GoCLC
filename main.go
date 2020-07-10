// Package main starts a GoCLC server on the local machine
package main

import (
	"flag"
	"log"
	"server"
)

func main() {
	args := getArgs()
	log.Printf("starting GoCLC server on %s:%s\n", *args["-address"], *args["-port"])
	server.Listen(*args["-address"], *args["-port"])
}

func getArgs() map[string]*string {
	address := flag.String("address", "localhost", "address to run server on")
	port := flag.String("port", "8000", "port to run server on")
	flag.Parse()
	return map[string]*string{"-address": address, "-port": port}
}
