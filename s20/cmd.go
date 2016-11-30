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

// Discover queries the local network to identify any
// S20s that have been paired and are listening
func Discover(timeout int) ([]string, error) {
	inBuf := make([]byte, 1024)
	devices := make([]string, 0)
	var readLen int
	var fromAddr *net.UDPAddr

	// get network connection
	ourAddr, err = net.ResolveUDPAddr("udp", ":10000")
	conn, err = net.ListenUDP("udp", ourAddr)
	checkErr(err)
	defer conn.Close()

	// send the Discover message
	server := fmt.Sprintf("%s:%d", bcastIP, udpDiscoverPort) // "255.255.255.255", 10000
	serverAddr, err = net.ResolveUDPAddr("udp", server)
	checkErr(err)
	checkErr(err)
	discoverMsg := []byte(magic)
	discoverMsg = append(discoverMsg, discovery...)
	sendLen, err := conn.WriteToUDP(discoverMsg, serverAddr)
	checkErr(err)
	fmt.Println("Sent", sendLen, "bytes")

	// read all replies
	for true {
		readLen, fromAddr, err = conn.ReadFromUDP(inBuf)
		fmt.Println("Read ", readLen, "bytesfrom ", fromAddr)
		txtutil.Dump(string(inBuf[:readLen]))
	}
	return devices, nil
}
