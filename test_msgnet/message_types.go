package main

import (
	"bytes"

	salticidae "github.com/masonforest/go-salticidae-native"
)

type MsgAck struct{}

func (msg MsgAck) Encode() []byte {
	b := make([]byte, 0)

	return b
}

func DecodeAck(r *bytes.Reader) MsgAck {
	return MsgAck{}
}

type MsgHello struct {
	Name string
	Text string
}

func (msg MsgHello) Encode() []byte {
	b := make([]byte, 0)
	b = append(b, salticidae.WriteString(msg.Name)...)
	b = append(b, []byte(msg.Text)...)

	return b
}

func DecodeHello(r *bytes.Reader) MsgHello {
	return MsgHello{
		Name: salticidae.ReadString(r),
		Text: string(salticidae.ReadToEnd(r)),
	}
}
