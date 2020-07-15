package server

type Broadcaster struct {
	clients map[string]*Client
	recieve chan string
}

// NewBroadcaster constructs a new Broadcaster.
func newBroadcaster() *Broadcaster {
	var b Broadcaster
	b.clients = make(map[string]*Client)
	b.recieve = make(chan string)
	return &b
}

func startBroadcaster() *Broadcaster {
	b := newBroadcaster()
	go b.broadcast()
	return b

}

func (b *Broadcaster) broadcast() {
	for {
		message := <-b.recieve
		for _, client := range b.clients {
			client.Write(message)
		}
	}
}

func (b *Broadcaster) addClient(client *Client) {
	b.clients[client.name] = client
}
