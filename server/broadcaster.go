package server

type Broadcaster struct {
	clients map[string]Client
}

// NewBroadcaster constructs a new Broadcaster.
func NewBroadcaster() *Broadcaster {
	var b Broadcaster
	b.clients = make(map[string]Client)
	return &b
}

func (b Broadcaster) addClient(client Client) {
	b.clients[client.name] = client
}
