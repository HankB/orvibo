package s20_test_pair

import (
	"fmt"
	//"net"
	"os"
	"testing"

	s20 "github.com/HankB/orvibo/s20"
)

var ssid string
var pwd string
var mid string

const magic = "\x68\x64" // copied rather than exported - just for testing

// fetch SSID and password from environment for testing, substitute
// some reasonable defaults if not provided.
func init() {
	ssid = os.Getenv("SSID")
	if len(ssid) == 0 {
		ssid = "NO_SSID"
	}
	pwd = os.Getenv("PASSWORD")
	if len(pwd) == 0 {
		pwd = "NO_PWD"
	}
	mid = os.Getenv("MODULE_ID")
	if len(mid) == 0 {
		mid = "NO_MODULE_ID"
	}
	fmt.Printf("SSID=\"%s\", PWD=\"%s\" MID=\"%s\"\n", ssid, pwd, mid)
}

// TestPair try to pair - keep this last as it fails if
// the host is not associated with an S20 in AP mode.
func TestPair(t *testing.T) {
	init()
	s20s := s20.Pair()
	if len(s20s) != 0 {
		t.Fail()
	}
}
