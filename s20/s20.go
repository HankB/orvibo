// Package to support operations involving the Orvibo S20
package s20

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

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
const MAGIC = []byte { 0x68, 0x64 }
const DISCOVERY = []byte {0x00, 0x06, 0x71, 0x61}
const DISCOVERY_RESP = []byte {0x00, 0x2a, 0x71, 0x61}
const SUBSCRIBE = []byte {x00, 0x1e\x63\x6c}
const SUBSCRIBE_RESP = []byte {x00, 0x18\x63\x6c}
const CONTROL = []byte {x00, 0x17, 0x64, 0x63}
const CONTROL_RESP = []byte {x00, 0x17, 0x73, 0x66}
const PADDING_1 = []byte {x20, 0x20, 0x20, 0x20, 0x20, 0x20}
const PADDING_2 = []byte {x00, 0x00, 0x00, 0x00}
const ON = []byte {x01}
const OFF = []byte {x00}

func Good() (int) {
	return 1
}
