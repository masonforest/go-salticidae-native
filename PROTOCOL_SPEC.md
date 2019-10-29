# The Salticidae Protocol

### Transports
Salticidae runs over raw [TCP](https://tools.ietf.org/html/rfc793) or [TLS](https://tools.ietf.org/html/rfc5246)
### Checksumming
Salticidae optionally include the first four bytes of a [SHA1](https://tools.ietf.org/html/rfc3174) hash of the message payload with each message to ensure against message corruption.

### Encoding
Numbers are little endian encoded.

### Messages

|                 | Magic          | Opcode   | Checksum (optional) | Payload |
| --------------- | -------------- | -------- |-------------------- | ------ |
| Type            | Unsigned Int32 | Unsigned Int8    | Unsigned Int32 | Bytes  |
| Number of bytes | 4              | Varies   | 4              | Varies |

### Payloads

Payload encoding is application specific but often has the form:

|                 | Length         | Data     | Length         | Data | ... |
| --------------- | -------------- | -------- |--------------- | ------ | --- |
| Type            | Unsigned Int32 | Bytes    | Unsigned Int32 | Bytes  | ... |
| Number of bytes | 4              | Varies   | 4              | Varies | ... |
