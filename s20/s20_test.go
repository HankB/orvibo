package s20_test

import (
	"fmt"
	"net"
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

func ExampleInit() {
	s20.Init("127.0.0.1", "this is the password", "DeviceID")
	fmt.Println(s20.Get())
	// Output:
	// 127.0.0.1 this is the password 10.10.100.254:48899
}

func TestIsThisHost(t *testing.T) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for _, ifce := range addr {
		ourIP, _, err := net.ParseCIDR(ifce.String())

		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		isMe, err := s20.IsThisHost(ourIP)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		if !isMe {
			fmt.Printf("not host %s\n", ifce.String())
			t.Fail()
		}
	}
	// now try one that is Not our IP
	ourIP, _, err := net.ParseCIDR("172.217.4.238/24") // google
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	isMe, err := s20.IsThisHost(ourIP)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if isMe {
		fmt.Printf("is host 172.217.4.238/24\n")
		t.Fail()
	}

}

/*
// try to pair - keep this last as it fails if
// the host is not associated with an S20 in AP mode.
func TestPair(t *testing.T) {
	s20s := s20.Pair()
	if len(s20s) != 0 {
		t.Fail()
	}
}
*/
