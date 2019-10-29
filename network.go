package salticidae

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"net"
)

const MAGIC = 0

type Encodable interface {
	Encode() []byte
}

type Network struct {
	Listener net.Listener
	Conn     net.Conn
	OnAccept func()
	Name     string
	Ops      map[uint8]func(d []byte)
}

type Header struct {
	Magic    uint32
	Opcode   uint8
	Length   uint32
	CheckSum []byte
}

func (n *Network) RegisterOp(opCode uint8, f func(d []byte)) {
	n.Ops[opCode] = f
}

func Encode(e Encodable, opCode uint8) []byte {
	payload := e.Encode()
	b := WriteUInt32(MAGIC)
	b = append(b, opCode)
	b = append(b, WriteUInt32(uint32(len(payload)))...)
	b = append(b, CheckSum(payload)...)
	b = append(b, payload...)

	return b
}

func CheckSum(b []byte) []byte {
	h := sha1.New()
	h.Write(b)
	sum := h.Sum(nil)
	return sum[:4]
}

func NewNetwork(address string) *Network {
	ln, _ := net.Listen("tcp", address)

	return &Network{
		Listener: ln,
		Ops:      make(map[uint8]func(b []byte)),
	}
}

func (n *Network) Send(e Encodable, opCode uint8) {
	n.Conn.Write(Encode(e, opCode))
}

func (n *Network) Start() {
	n.Conn, _ = n.Listener.Accept()
	n.OnAccept()
}

func (n Network) Connect(host string) {
	c, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		headerBytes := make([]byte, 13)
		length, _ := c.Read(headerBytes)
		if length > 0 {
			header := parseHeader(headerBytes)
			payload := make([]byte, header.Length)
			c.Read(payload)
			n.Ops[header.Opcode](payload)
		}
	}
}

func parseHeader(b []byte) Header {
	r := bytes.NewReader(b)
	return Header{
		Magic:    ReadUInt32(r),
		Opcode:   ReadUInt8(r),
		Length:   ReadUInt32(r),
		CheckSum: ReadBytes(r, 4),
	}
}
