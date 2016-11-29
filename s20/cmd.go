// Package s20 to support operations involving the Orvibo S20
package s20

import (
	"fmt"
	"net"
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
	devices := make([]string, 0)
	// get server handle
	server := fmt.Sprintf("%s:%d", bcastIP, udpDiscoverPort)
	serverAddr, err = net.ResolveUDPAddr("udp", server)
	checkErr(err)
	client := fmt.Sprintf("%s:%d", "", 0)
	ourAddr, err = net.ResolveUDPAddr("udp", client)
	checkErr(err)
	// fmt.Println("calling DialUDP()")
	conn, err = net.DialUDP("udp", ourAddr, serverAddr)
	checkErr(err)
	// fmt.Println("Established UDP socket")
	return devices, nil
}
