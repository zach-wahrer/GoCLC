package server

import (
	"fmt"
	"log"
	"os/exec"
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

func (b *Broadcaster) welcomeClient(client *Client) {
	b.sendToOne(client, runCommand("/welcome"))
}

func (b *Broadcaster) askUsername(client *Client) {
	b.sendToOne(client, runCommand("/AskUsername"))
}

func (b *Broadcaster) sendError(client *Client, err error) {
	b.sendToOne(client, wrapServerMessage(fmt.Sprintf("%s", err)))
}

func (b *Broadcaster) greetUserByName(client *Client) {
	message := fmt.Sprintf("%s %s%s", UserGreeting, client.name, UserGreetingPunc)
	b.sendToOne(client, wrapServerMessage(message))
}

func (b *Broadcaster) announceNewClient(client *Client) {
	message := fmt.Sprintf("%s %s", client.name, UserAnouncement)
	b.sendToAll(wrapServerMessage(message))
}

func (b *Broadcaster) announceDepartedClient(client *Client) {
	message := fmt.Sprintf("%s %s %s", ServerTag, client.name, UserDepartedAnnouncement)
	b.sendToAll(wrapServerMessage(message))
}

func (b *Broadcaster) sayGoodbye(client *Client) {
	b.sendToOne(client, wrapServerMessage(runCommand("/goodbye")))
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

func wrapServerMessage(message string) string {
	return fmt.Sprintf("%s%s%s%s %s", ServerColor, colorBold, ServerTag, colorReset, message)
}

func createASCIIArt(message, font string) string {
	out, err := exec.Command("python3", "ascii_art", font, message).Output()
	if err != nil {
		log.Print(fmt.Sprintf("ASCII art not generated properly: %v", err))
		return ""
	}
	return string(out)
}
