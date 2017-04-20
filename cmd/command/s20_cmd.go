// Package send commands to the S20
package main

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/HankB/orvibo/s20"
)

const replyTimeout = 1 // reply timeout in seconds
var s20s []s20.Device

func usage(errstr string) {
	if len(errstr) > 0 {
		fmt.Println(errstr)
	}
	fmt.Println("Usage", os.Args[0], "[-d]|[-s][-c [host|IP addr] [on|off]]")
	os.Exit(0)
}

func findDeviceByIP(addr *net.IPAddr) (s20.Device, error) {
	for _, device := range s20s {
		if bytes.Compare(addr.IP, device.IPAddr.IP) == 0 {
			return device, nil
		}
	}
	return s20s[0], errors.New("IP not found in discover list")
}

func main() {
	var cmd bool
	var state string

	// did the user pass arguments?
	if len(os.Args) == 1 {
		usage("")
	} else if os.Args[1] == "-d" { // exercise Discover command
		s20s, _ := s20.Discover(replyTimeout)
		for _, s20ds := range s20s {
			fmt.Println(s20ds)
		}
		os.Exit(0)
	} else if os.Args[1] == "-s" { // test Subscribe (first discovered S20)
		s20s, _ := s20.Discover(replyTimeout)
		if len(s20s) < 1 {
			usage("No S20s discovered")
		}
		e := s20.Subscribe(replyTimeout, &s20s[0])
		if e == nil {
			var state string
			if s20s[0].IsOn {
				state = "on"
			} else {
				state = "off"
			}
			fmt.Println("subscribed and currently", state)
		} else {
			fmt.Println("subscription error:", e)
		}
		os.Exit(0)
	} else if os.Args[1] == "-c" { // test Control (first discovered S20)
		if len(os.Args) < 4 {
			usage("")
		}

		// unpack on/off command
		if os.Args[3] == "on" {
			cmd = true
		} else if os.Args[3] == "off" {
			cmd = false
		} else {
			usage("command must be \"on\" or \"off\" (found \"" + os.Args[3] + "\")")
		}

		// resolve the IP address provided
		ipAddr, err := net.ResolveIPAddr("ip", os.Args[2])
		if err != nil {
			usage("Cannot resolve " + os.Args[2])
		}
		fmt.Println("resolved as ", ipAddr)

		s20s, _ = s20.Discover(replyTimeout)

		dev, err := findDeviceByIP(ipAddr)
		if len(s20s) < 1 {
			fmt.Println("No S20s discovered")
			os.Exit(0)
		}
		e := s20.Subscribe(replyTimeout, &dev)
		if e == nil {
			if dev.IsOn {
				state = "on"
			} else {
				state = "off"
			}
			fmt.Println("subscribed and currently", state)
		} else {
			fmt.Println("subscription error:", e)
		}
		e = s20.Control(replyTimeout, &dev, cmd)
		if e == nil {
			if dev.IsOn {
				state = "on"
			} else {
				state = "off"
			}
			fmt.Println("commanded and currently", state)
		} else {
			fmt.Println("control error:", e)
		}
		os.Exit(0)
	}
	fmt.Println("Unknown command arg", os.Args[1])
	os.Exit(1)
}
