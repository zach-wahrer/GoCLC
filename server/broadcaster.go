package server

type Broadcaster struct {
	clients map[string]*Client
	recieve chan string
}

func startBroadcaster() *Broadcaster {
	b := newBroadcaster()
	go broadcast(b)
	return b

}

func broadcast(b *Broadcaster) {
	for {
		message := <-b.recieve
		for _, client := range b.clients {
			client.Write(message)
		}
	}
}

// NewBroadcaster constructs a new Broadcaster.
func newBroadcaster() *Broadcaster {
	var b Broadcaster
	b.clients = make(map[string]*Client)
	b.recieve = make(chan string)
	return &b
}

func (b *Broadcaster) addClient(client *Client) {
	b.clients[client.name] = client
}
