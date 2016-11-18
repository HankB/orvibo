package s20

import (
	"fmt"
	"testing"

	s20 "github.com/HankB/orvibo/s20"
)

func main() {
	fmt.Println("Hello")
	// Output:
	// Hello

	s20.Dump(s20.MAGIC)
	// Output:
	// [ 86 64 ]
}

func TestIsPrint(t *testing.T) {
	if s20.IsPrint(10) {
		t.Fail()
	}

	if s20.IsPrint(0x20) {
		t.Fail()
	}
}

func TestDump(t *testing.T) {
	// this test does not work
	s20.Dump(s20.MAGIC)
	// Output:
	// [ 86 64 ]
	s20.Dump(s20.SUBSCRIBE)
	s20.Dump(s20.SUBSCRIBE + s20.CONTROL_RESP + s20.CONTROL_RESP + s20.CONTROL_RESP + s20.CONTROL_RESP)
}

func ExampleDump() {
	s20.Dump(s20.MAGIC)
	// Output:
	// [ 68 64 ]
}
