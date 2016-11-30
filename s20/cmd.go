// Package s20 to support operations involving the Orvibo S20
package s20

import (
	"fmt"
	"net"

	"github.com/HankB/txtutil"
)

// Code to find and send commands to Orvibo S20 WiFi controlled outlets
// The outlets need to brouight into the network first (AKA paired)
// and for that see pair.go

// import (
// "fmt"
// "net"
// "os"
// "time"
// )

// Discover queries the local network to identify any
// S20s that have been paired and are listening
func Discover(timeout int) ([]string, error) {
	inBuf := make([]byte, 1024)
	devices := make([]string, 0)
	var readLen int
	var fromAddr *net.UDPAddr

	// get server handle
	server := fmt.Sprintf("%s:%d", bcastIP, udpDiscoverPort)
	serverAddr, err = net.ResolveUDPAddr("udp", server)
	checkErr(err)
	// client := fmt.Sprintf("%s:%d", "0.0.0.0", udpDiscoverPort)
	// client := fmt.Sprintf(":%d", udpDiscoverPort)
	ourAddr, err = net.ResolveUDPAddr("udp", "192.168.1.132:10000")
	checkErr(err)
	// fmt.Println("calling DialUDP()")
	conn, err = net.DialUDP("udp", ourAddr, serverAddr)
	checkErr(err)
	defer conn.Close()
	// fmt.Println("Established UDP socket")

	discoverMsg := []byte(magic)
	discoverMsg = append(discoverMsg, discovery...)
	sendLen, err := conn.Write(discoverMsg)
	checkErr(err)
	fmt.Println("Sent", sendLen, "bytes")
	// send the Discover message
	readLen, fromAddr, err = conn.ReadFromUDP(inBuf)
	fmt.Println("Read ", readLen, "bytesfrom ", fromAddr)
	txtutil.Dump(string(inBuf[:readLen]))
	return devices, nil
}
