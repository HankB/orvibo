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

// fetch SSID and password from environment for testing
func init() {
	ssid = os.Getenv("SSID")
	pwd = os.Getenv("PASSWORD")
	fmt.Printf("SSID=%s, PWD=\"%s\"\n", ssid, pwd)
}

func ExampleDump() {
	// txt.Dump(s20.MAGIC)
	// 00000000  68 64                                             |hd|
	// 00000002
}

func ExampleInit() {
	s20.Init(ssid, pwd)
	fmt.Println(s20.Get())
	// Output:
	// ssid 127.0.0.1 this is the password 127.0.0.1:48899
}

func TestIsThisHost(t *testing.T) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, ifce := range addr {
		ourIP, _, err := net.ParseCIDR(ifce.String())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		isMe, err := s20.IsThisHost(ourIP)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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
		os.Exit(1)
	}
	isMe, err := s20.IsThisHost(ourIP)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isMe {
		fmt.Printf("is host 172.217.4.238/24\n")
		t.Fail()
	}

}

// try to pair - keep this last as it fails if
// the host is not asociated with an S20 in AP mode.
func TestPair(t *testing.T) {
	s20s := s20.Pair()
	if len(s20s) != 0 {
		t.Fail()
	}
}
