// Package send commands to the S20
package main

import (
	"bytes"
	"errors"
	"net"
	"os"

	"github.com/HankB/orvibo/s20"
	"github.com/HankB/txtutil"
)

// output conventions -
// - by default output with priority >= 4 is displayed.
// - error messages are set to 6
// - minimal information to monitor progess set to 4 (display by default)
// - additional information at 3.
// - full dumps of sent/received messages at 1

const replyTimeout = 1 // reply timeout in seconds
var s20s []s20.Device

func usage(errstr string) {
	if len(errstr) > 0 {
		txtutil.PriFmtPrintln(6, errstr)
	}
	txtutil.PriFmtPrintln(6, "Usage", os.Args[0], "[-v|-vv][[-d]|[-s][-c [host|IP addr] [on|off]]]")
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

// process arguments (and return) until there are no more
// to process
func processArgs(args []string) []string {
	var cmd bool
	var state string

	if len(args) == 1 {
		usage("")
	} else if args[1] == "-v" { // increase verbosity
		if len(args) < 3 {
			usage("\"-v\" cannot be last argument")
		}
		args = append(args[:1], args[2:]...)
		txtutil.SetDumpPriority(4)
		// os.Exit(0)
	} else if args[1] == "-vv" { // increase verbosity even more
		if len(args) < 3 {
			usage("\"-vv\" cannot be last argument")
		}
		args = append(args[:1], args[2:]...)
		txtutil.SetDumpPriority(3)
		txtutil.SetDumpPriority(3)
		// os.Exit(0)
	} else if args[1] == "-d" { // exercise Discover command
		s20s, _ := s20.Discover(replyTimeout)
		for _, s20ds := range s20s {
			txtutil.PriFmtPrintln(4, s20ds)
		}
		os.Exit(0)
	} else if args[1] == "-s" { // test Subscribe (first discovered S20)
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
			txtutil.PriFmtPrintln(4, s20s[0].IPAddr.IP, "subscribed and currently", state)
		} else {
			txtutil.PriFmtPrintln(6, "subscription error:", e)
		}
		os.Exit(0)
	} else if args[1] == "-c" { // exercise Control 
		if len(args) < 4 {		// must include host or IP and on/off command 
			usage("")
		}

		// unpack on/off command
		if args[3] == "on" {
			cmd = true
		} else if args[3] == "off" {
			cmd = false
		} else {
			usage("command must be \"on\" or \"off\" (found \"" + args[3] + "\")")
		}

		// resolve the IP address provided
		ipAddr, err := net.ResolveIPAddr("ip", args[2])
		if err != nil {
			usage("Cannot resolve " + args[2])
		}
		txtutil.PriFmtPrintln(3, "resolved as ", ipAddr)

		s20s, _ = s20.Discover(replyTimeout)	// Discovery must precede subscription and command

		dev, err := findDeviceByIP(ipAddr)		// match desired device with those discovered.
		if err != nil {
			txtutil.PriFmtPrintln(4, "No S20s discovered at that IP", ipAddr.IP)
			os.Exit(0)
		}

		e := s20.Subscribe(replyTimeout, &dev)	// subscribe to desired device
		if e == nil {
			if dev.IsOn {
				state = "on"
			} else {
				state = "off"
			}
			txtutil.PriFmtPrintln(3, "subscribed and currently", state)
		} else {
			txtutil.PriFmtPrintln(3, "subscription error:", e)
		}
		e = s20.Control(replyTimeout, &dev, cmd)
		if e == nil {
			if dev.IsOn {
				state = "on"
			} else {
				state = "off"
			}
			txtutil.PriFmtPrintln(4, "commanded and currently", state)
		} else {
			txtutil.PriFmtPrintln(4, "control error:", e)
		}
		os.Exit(0)
	} else {
		// did the user pass arguments?
		usage("Unknown command arg \"" + os.Args[1] + "\"")
	}
	return args
}

func main() {
	args := os.Args
	for len(args) > 1 {
		args = processArgs(args)
	}
	usage("")
}
