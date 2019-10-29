package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestHelloRoundTrip(t *testing.T) {
	original := MsgHello{
		opcode: 1,
		name:   "alice",
		text:   "Hello there!",
	}
	encoded := original.Encode()
	fmt.Println(hex.EncodeToString(encoded))
	decoded := DecodeHello(bytes.NewReader(encoded))

	if !reflect.DeepEqual(original, decoded) {
		t.Errorf("expected %+v; got %+v", original, decoded)
	}
}
