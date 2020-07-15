package server

import "log"

type Broadcaster struct {
	clients map[string]Client
	send    chan string
	recieve chan string
}

func startBroadcaster(send, recieve chan string) {
	b := newBroadcaster()
	b.send = send
	b.recieve = recieve

	for {
		select {
		case message := <-b.recieve:
			log.Print(message)
		}
	}

}

// NewBroadcaster constructs a new Broadcaster.
func newBroadcaster() *Broadcaster {
	var b Broadcaster
	b.clients = make(map[string]Client)
	return &b
}
