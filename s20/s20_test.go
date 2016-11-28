package s20_test

import (
	"fmt"
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
	s20.Init("127.0.0.1", ssid, pwd)
	fmt.Println(s20.Get())
	// Output:
	// ssid 127.0.0.1 this is the password 127.0.0.1:48899
}

func TestPair(t *testing.T) {
	s20s := s20.Pair()
	if len(s20s) != 0 {
		t.Fail()
	}
}
