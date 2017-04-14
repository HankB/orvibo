// Package s20 to support operations involving the Orvibo S20
package s20

import (
	"bytes"
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/HankB/txtutil"
)

// Code to find and send commands to Orvibo S20 WiFi controlled outlets
// The outlets need to brought into the network first (AKA paired)
// and for that see pair.go

// Discover queries the local network to identify any
// S20s that have been paired and are listening
func Discover(timeout time.Duration) ([]net.UDPAddr, error) {
	inBuf := make([]byte, readBufLen)
	devices := make([]net.UDPAddr, 0)
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
	err = conn.SetReadDeadline(time.Now().Add(timeout * time.Second))
	noErr := true
	for noErr {
		readLen, fromAddr, err = conn.ReadFromUDP(inBuf)
		if err != nil {
			noErr = false
		} else {
			// fmt.Println("Read ", readLen, "bytes from ", fromAddr)
			isMe, err := IsThisHost(fromAddr.IP)
			checkErr(err)
			if isMe { // Seeing our own transmission?
				continue // just ignore it. We're not an S20. ;)
			}
			found := false
			for _, addrIter := range devices {
				// fmt.Println("comparing ", addrIter, *fromAddr)
				if reflect.DeepEqual(addrIter, *fromAddr) {
					found = true
				}
			}
			if !found {
				devices = append(devices, *fromAddr)
				fmt.Println("adding", fromAddr, "count", len(devices))
				txtutil.Dump(string(inBuf[:readLen]))
			}
		}
	}
	return devices, nil
}

// IsThisHost determine if the address belongs to localhost
// Perhaps should be moved to netutil package
func IsThisHost(check net.IP) (bool, error) {
	//  get a list of our IP addresses
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return true, err // bool true meaningless here
	}
	// compare to the provided address
	for _, thisHost := range addr {
		ourIP, _, err := net.ParseCIDR(thisHost.String())
		if err != nil {
			return true, err // bool true meaningless here
		}
		if bytes.Compare(ourIP, check) == 0 {
			return true, nil
		}
	}
	return false, nil
}
