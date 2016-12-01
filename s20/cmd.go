// Package s20 to support operations involving the Orvibo S20
package s20

import (
	"fmt"
	"net"
	"time"

	"github.com/HankB/txtutil"
)

// Code to find and send commands to Orvibo S20 WiFi controlled outlets
// The outlets need to brouight into the network first (AKA paired)
// and for that see pair.go

// Discover queries the local network to identify any
// S20s that have been paired and are listening
func Discover(timeout int) ([]*net.UDPAddr, error) {
	inBuf := make([]byte, 1024)
	devices := make([]*net.UDPAddr, 0)
	var readLen int
	var fromAddr *net.UDPAddr

	// get network connection
	sender := fmt.Sprintf(":%d", udpDiscoverPort)
	ourAddr, err = net.ResolveUDPAddr("udp", sender)
	checkErr(err)
	conn, err = net.ListenUDP("udp", ourAddr)
	checkErr(err)
	defer conn.Close()

	// send the Discover message
	server := fmt.Sprintf("%s:%d", bcastIP, udpDiscoverPort)
	serverAddr, err = net.ResolveUDPAddr("udp", server)
	checkErr(err)
	discoverMsg := []byte(magic)
	discoverMsg = append(discoverMsg, discovery...)
	sendLen, err := conn.WriteToUDP(discoverMsg, serverAddr)
	checkErr(err)
	fmt.Println("Sent", sendLen, "bytes")

	// read all replies
	err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	noErr := true
	for noErr {
		readLen, fromAddr, err = conn.ReadFromUDP(inBuf)
		if err != nil {
			noErr = false
		} else {
			fmt.Println("Read ", readLen, "bytes from ", fromAddr)
			found := false
			for _, addrIter := range devices {
				fmt.Println("comparing ", addrIter, fromAddr)
				if addrIter == fromAddr {
					found = true
				}
			}
			if !found {
				devices = append(devices, fromAddr)
				fmt.Println("adding", fromAddr, "len", len(devices))
			}
			txtutil.Dump(string(inBuf[:readLen]))
		}
	}
	return devices, nil
}
