package main

import (
	"bytes"
	"fmt"

	salticidae "github.com/masonforest/go-salticidae-native"
)

type MyNet struct {
	Name    string
	Network *salticidae.Network
}

func NewMyNet(name string, address string) MyNet {
	network := salticidae.NewNetwork(address)
	fmt.Printf("%s accepted, waiting for greetings.\n", name)
	network.OnAccept = func() {
		fmt.Printf("%s connected, sending hello.\n", name)
		network.Send(MsgHello{
			Name: name,
			Text: "Hello there!",
		}, 0)
	}
	network.RegisterOp(0, func(b []byte) {
		msg := DecodeHello(bytes.NewReader(b))
		fmt.Printf("[%s] %s says %s\n", name, msg.Name, msg.Text)
		network.Send(MsgAck{}, 1)
	})
	network.RegisterOp(1, func(b []byte) {
		fmt.Printf("[%s] the peer knows\n", name)
	})
	return MyNet{
		Name:    name,
		Network: network,
	}
}

func setupNetwork() {
	alicesAddr := "127.0.0.1:12345"
	bobsAddr := "127.0.0.1:12346"
	alice := NewMyNet("alice", ":12345")
	bob := NewMyNet("bob", ":12346")
	go alice.Network.Start()
	go bob.Network.Start()
	go alice.Network.Connect(bobsAddr)
	bob.Network.Connect(alicesAddr)

}
func main() {
	setupNetwork()
}
