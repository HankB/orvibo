// Package to support operations involving the Orvibo S20
package s20

//"net"
//"os"
//"strconv"

/*
byte arrays in python used to build messages
MAGIC = b'\x68\x64'
DISCOVERY = b'\x00\x06\x71\x61'
DISCOVERY_RESP = b'\x00\x2a\x71\x61'
SUBSCRIBE = b'\x00\x1e\x63\x6c'
SUBSCRIBE_RESP = b'\x00\x18\x63\x6c'
CONTROL = b'\x00\x17\x64\x63'
CONTROL_RESP = b'\x00\x17\x73\x66'
PADDING_1 = b'\x20\x20\x20\x20\x20\x20'
PADDING_2 = b'\x00\x00\x00\x00'
ON = b'\x01'
OFF = b'\x00'
*/
// string literals used to build messages, syntax strikingly similar to
// python.
const MAGIC = "\x68\x64"
const DISCOVERY = "\x00\x06\x71\x61"
const DISCOVERY_RESP = "\x00\x2a\x71\x61"
const SUBSCRIBE = "\x00\x1e\x63\x6c"
const SUBSCRIBE_RESP = "\x00\x18\x63\x6c"
const CONTROL = "\x00\x17\x64\x63"
const CONTROL_RESP = "\x00\x17\x73\x66"
const PADDING_1 = "\x20\x20\x20\x20\x20\x20"
const PADDING_2 = "\x00\x00\x00\x00"
const ON = "\x01"
const OFF = "\x00"

func Good() int {
	return 1
}
