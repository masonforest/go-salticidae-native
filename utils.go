package salticidae

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func WriteUInt32(n uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, n)
	return b
}

func WriteString(s string) []byte {
	b := WriteUInt32(uint32(len(s)))
	return append(b, []byte(s)...)
}

func ReadToEnd(r io.Reader) []byte {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func ReadUInt8(r io.Reader) uint8 {
	var n uint8
	binary.Read(r, binary.LittleEndian, &n)
	return n
}

func ReadUInt32(r io.Reader) uint32 {
	var len uint32
	binary.Read(r, binary.LittleEndian, &len)
	return len
}

func ReadBytes(r io.Reader, n int) []byte {
	buf := make([]byte, n)
	_, err := r.Read(buf)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return buf
}

func ReadString(r io.Reader) string {
	s := make([]byte, ReadUInt32(r))
	r.Read(s)
	return string(s)
}
