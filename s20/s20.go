// Package s20 supports operations involving the Orvibo S20
package s20

import "fmt"

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
const magic = "\x68\x64"
const discovery = "\x00\x06\x71\x61"
const discoveryResp = "\x00\x2a\x71\x61"
const subscribe = "\x00\x1e\x63\x6c"
const subscribeResp = "\x00\x18\x63\x6c"
const control = "\x00\x17\x64\x63"
const controlResp = "\x00\x17\x73\x66"
const padding1 = "\x20\x20\x20\x20\x20\x20"
const padding2 = "\x00\x00\x00\x00"
const on = "\x01"
const off = "\x00"

var ssid = ""  // SSID we will pair with
var pwd = ""   // password for SSID.
var swStr = "" // string used to establish connection
var mid = ""   // name module will use

const s20IP = "10.10.100.254"     // IP address used by the S20
const ourIP = "10.10.100.150"     // IP address S20 will assign to host
const bcastIP = "255.255.255.255" // broadcast IP address

const udpRcvPort = 9884       // port we listen on when pairing
const udpSndPort = 48899      // port S20 listens on when pairing
const udpDiscoverPort = 10000 // port to send discovery message

// Init saves network parameters for later usage and
// opens the port
func Init(SSID string, password string, module_ID string) {
	ssid = SSID
	// ip = IP
	pwd = password
	swStr = fmt.Sprintf("%s:%d", s20IP, udpSndPort)
	mid = module_ID
}

//Get returns 'object' parameters for testing
func Get() (string, string, string) {
	return ssid, pwd, swStr
}
