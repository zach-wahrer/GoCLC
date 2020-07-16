package server

type Broadcaster struct {
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
		b.clients[client.name] = client
		return true
	}
	return false
}

func (b *Broadcaster) removeClient(client *Client) {
	delete(b.clients, client.name)
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
