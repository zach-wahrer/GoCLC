// Package main starts a GoCLC server on the local machine
package main

import (
	"client"
	"flag"
	"log"
	"server"
)

func main() {
	args := getArgs()
	if *args["-server"] == "true" {
		log.Printf("starting GoCLC server on %s:%s\n", *args["-address"], *args["-port"])
		server.Listen(*args["-address"], *args["-port"])
	} else {
		c := client.NewClient(*args["-address"], *args["-port"])
		c.Start()
	}

}

func getArgs() map[string]*string {
	address := flag.String("address", "localhost", "address to connect to or run server on")
	port := flag.String("port", "8000", "port to connect to or run server on")
	server := flag.String("server", "false", "start in server mode")
	flag.Parse()
	return map[string]*string{"-address": address, "-port": port, "-server": server}
}
