package s20_test

import (
	"fmt"
	"testing"

	s20 "github.com/HankB/orvibo/s20"
)

func ExampleDump() {
	// txt.Dump(s20.MAGIC)
	// 00000000  68 64                                             |hd|
	// 00000002
}

func ExampleInit() {
	s20.Init("127.0.0.1", "ssid", "this is the password")
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
