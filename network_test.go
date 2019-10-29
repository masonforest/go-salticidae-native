package salticidae

import (
	"bytes"
	"reflect"
	"testing"
)

type MsgAck struct {
	opcode uint8
}

type MsgHello struct {
	opcode uint8
	Name   string
	Text   string
}

func (msg MsgAck) Encode() []byte {
	return make([]byte, 0)
}

func (msg MsgHello) Encode() []byte {
	b := make([]byte, 0)
	b = append(b, WriteString(msg.Name)...)
	b = append(b, []byte(msg.Text)...)

	return b
}

func (msg MsgHello) Opcode() uint8 {
	return msg.opcode
}

func (msg MsgAck) Opcode() uint8 {
	return msg.opcode
}

func Decode(r *bytes.Reader) Encodable {
	ReadUInt32(r)
	opcode := ReadUInt8(r)
	ReadBytes(r, 4)
	ReadUInt32(r)
	return MsgHello{
		opcode: opcode,
		Name:   ReadString(r),
		Text:   string(ReadToEnd(r)),
	}
}

func TestNetworkMessageRoundTrip(t *testing.T) {
	original := MsgHello{
		opcode: 0,
		Name:   "alice",
		Text:   "Hello there!",
	}
	encoded := Encode(original, 0)
	decoded := Decode(bytes.NewReader(encoded))
	if !reflect.DeepEqual(original, decoded) {
		t.Errorf("expected %+v; got %+v", original, decoded)
	}
}
