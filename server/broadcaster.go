package server

import (
	"fmt"
	"sync"
)

// Broadcaster manages connections with clients and sends messages to them.
type Broadcaster struct {
	mu      sync.Mutex
	clients map[string]*Client
	receive chan string
}

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

func (b *Broadcaster) sendToOne(client *Client, message string) {
	client.write(message)
}

func (b *Broadcaster) sendToAll(message string) {
	for _, client := range b.clients {
		if err := client.write(message); err != nil {
			b.removeClient(client)
		}
	}
}

func (b *Broadcaster) greetUserByName(client *Client) {
	b.sendToOne(client, fmt.Sprintf("%s %s %s%s", ServerTag, UserGreeting,
		client.name, UserGreetingPunc))
}

func (b *Broadcaster) announceNewClient(client *Client) {
	b.sendToAll(fmt.Sprintf("%s %s %s", ServerTag, client.name, UserAnouncement))
}

func (b *Broadcaster) sayGoodbye(client *Client) {
	b.sendToOne(client, ServerTag+" "+runCommand("/goodbye"))
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

func (b *Broadcaster) usernameAvailable(username string) bool {
	_, ok := b.clients[username]
	return !ok
}
