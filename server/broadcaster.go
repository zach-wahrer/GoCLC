package server

import "sync"

type Broadcaster struct {
	mu      sync.Mutex
	clients map[string]*Client
	receive chan string
}

// NewBroadcaster constructs a new Broadcaster.
func newBroadcaster() *Broadcaster {
	var b Broadcaster
	b.clients = make(map[string]*Client)
	b.receive = make(chan string)
	return &b
}

func startBroadcaster() *Broadcaster {
	b := newBroadcaster()
	go b.broadcast()
	return b

}

func (b *Broadcaster) broadcast() {
	for {
		message := <-b.receive
		b.sendToAll(message)
	}
}

func (b *Broadcaster) addClient(client *Client) bool {
	if b.usernameAvailable(client.name) {
		b.mu.Lock()
		b.clients[client.name] = client
		b.mu.Unlock()
		return true
	}
	return false
}

func (b *Broadcaster) removeClient(client *Client) {
	b.mu.Lock()
	delete(b.clients, client.name)
	b.mu.Unlock()
}

func (b *Broadcaster) sendToAll(message string) {
	for _, client := range b.clients {
		client.Write(message)
	}
}

func (b *Broadcaster) usernameAvailable(username string) bool {
	_, ok := b.clients[username]
	return !ok
}
