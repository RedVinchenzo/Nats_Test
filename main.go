package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

func main() {
	sendHello(1)
	sendHello(2)
}

func sendHello(i int) {
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Publish("foo", []byte("Hello "+i))
}

func receiveMessages() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
}
