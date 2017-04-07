package s20

import (
	"fmt"
	"net"
	"testing"
)

func ExampleInit() {
	Init("127.0.0.1", "this is the password", "DeviceID")
	fmt.Println(Get())
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
		isMe, err := IsThisHost(ourIP)
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
	isMe, err := IsThisHost(ourIP)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if isMe {
		fmt.Printf("is host 172.217.4.238/24\n")
		t.Fail()
	}

}
