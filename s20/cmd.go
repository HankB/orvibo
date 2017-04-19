// Package s20 to support operations involving the Orvibo S20
package s20

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/HankB/txtutil"
)

// Code to find and send commands to Orvibo S20 WiFi controlled outlets
// The outlets need to brought into the network first (AKA paired)
// and for that see pair.go

// copyMac allocates space for and copies MAC addr from a byte slice
func copyMac(dst *net.HardwareAddr, src []byte) {
	*dst = make([]byte, len(src))
	copy(*dst, src)
}

// unpackDiscoverResp extracts required information from the Discover reply
// sent by the S20
func unpackDiscoverResp(ip *net.UDPAddr, buff []byte) Device {
	d := Device{IpAddr: *ip}
	copyMac(&d.Mac, buff[7:7+6])
	copyMac(&d.ReverseMac, buff[7+12:7+6+12])
	d.IsOn = buff[41] != 0
	return d
}

// Discover queries the local network to identify any
// S20s that have been paired and are listening. The timeout
// is how long the process will wait for a reply after sending
// or receiving the last reply. (S20 replies 10 times to this msg)
func Discover(timeout time.Duration) ([]Device, error) {
	inBuf := make([]byte, readBufLen)
	devices := make([]Device, 0)
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
				if reflect.DeepEqual(addrIter.IpAddr, *fromAddr) {
					found = true
				}
			}
			if !found {
				d := unpackDiscoverResp(fromAddr, inBuf)
				fmt.Println("unpacked", d)
				devices = append(devices, d)
				fmt.Println("adding", fromAddr, "count", len(devices), "on", inBuf[41], "mac", d.Mac)
				txtutil.Dump(string(inBuf[:readLen]))
			}
		}
	}
	return devices, nil
}

// Subscribe subscribes to the S20 and is required before sending
// further commands.
func Subscribe(timeout time.Duration, s20device *Device) error {
	inBuf := make([]byte, readBufLen)

	xmitBuf := bytes.NewBufferString(magic + subscribe)
	xmitBuf.Write(s20device.Mac)
	xmitBuf.WriteString(padding1)
	fmt.Println("building subscription")
	txtutil.Dump(xmitBuf.String())

	// get network connection, listen on udpDiscoverPort
	sender := fmt.Sprintf(":%d", udpDiscoverPort)
	ourAddr, err = net.ResolveUDPAddr("udp", sender)
	checkErr(err)
	conn, err = net.ListenUDP("udp", ourAddr)
	checkErr(err)
	defer conn.Close()

	// send the Subscribe message
	sendLen, err := conn.WriteToUDP(xmitBuf.Bytes(), &s20device.IpAddr)
	checkErr(err)
	fmt.Println("Sent", sendLen, "bytes")

	// read all replies
	err = conn.SetReadDeadline(time.Now().Add(timeout * time.Second))
	readLen, fromAddr, err := conn.ReadFromUDP(inBuf)
	if err != nil {
		return err
	}
	fmt.Println("Subscribe Reply", readLen, "bytes from ", fromAddr)
	txtutil.Dump(string(inBuf[:readLen]))
	s20device.IsOn = inBuf[23] != 0 // capture on/off state
	if bytes.Compare(inBuf[2:6], []byte(subscribeResp)) != 0 {
		fmt.Println("unexpected reply")
		return errors.New("unexpected response to subscribe")
	}
	s20device.subscriptionTime = time.Now()
	return nil
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
