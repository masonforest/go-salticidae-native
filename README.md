# The Salticidae protocol in idiomatic golang

_This is just a proof of concept I wrote to learn how the protocol works. See: [salticidae-go](https://github.com/Determinant/salticidae-go)_ for a fully functional implementation.


### Install

    $ go get https://github.com/masonforest/go-salticidae-native

#### Test

    $ git clone https://github.com/masonforest/go-salticidae-native
    $ cd go-salticidae-native
    $ go test

### Demo

    $ git clone https://github.com/masonforest/go-salticidae-native
    $ cd go-salticidae-native
    $ go run test_msgnet/main.go test_msgnet/message_types.go
    alice accepted, waiting for greetings.
    bob accepted, waiting for greetings.
    alice connected, sending hello.
    bob connected, sending hello.
    [bob] alice says Hello there!
    [alice] bob says Hello there!
    [alice] the peer knows
    [bob] the peer knows

### Protocol documentation

See: [Protocol Spec](https://github.com/masonforest/go-salticidae-native/blob/master/PROTOCOL_SPEC.md)
