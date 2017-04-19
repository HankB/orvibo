// Package send commands to the S20
package main

import (
	"fmt"
	"os"

	"github.com/HankB/orvibo/s20"
)

const replyTimeout = 2 // reply timeout in seconds

func usage(errstr string) {
	if len(errstr) > 0 {
		fmt.Println(errstr)
	}

}
func main() {
	// did the user pass arguments?

	if len(os.Args) == 1 {
		fmt.Println("Usage", os.Args[0], "[-d]|[-s][-c]")
		os.Exit(0)
	} else if os.Args[1] == "-d" { // exercise Discover command
		s20s, _ := s20.Discover(replyTimeout)
		for _, s20ds := range s20s {
			fmt.Println(s20ds)
		}
		os.Exit(0)
	} else if os.Args[1] == "-s" { // test Subscribe (first discovered S20)
		s20s, _ := s20.Discover(replyTimeout)
		if len(s20s) < 1 {
			fmt.Println("No S20s discovered")
			os.Exit(0)
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
	} else if os.Args[1] == "-c" { // test Subscribe (first discovered S20)
		s20s, _ := s20.Discover(replyTimeout)
		var cmd bool
		var state string

		if len(s20s) < 1 {
			fmt.Println("No S20s discovered")
			os.Exit(0)
		}
		e := s20.Subscribe(replyTimeout, &s20s[0])
		if e == nil {
			if s20s[0].IsOn {
				state = "on"
				cmd = false
			} else {
				state = "off"
				cmd = true
			}
			fmt.Println("subscribed and currently", state)
		} else {
			fmt.Println("subscription error:", e)
		}
		e = s20.Control(replyTimeout, &s20s[0], cmd)
		if s20s[0].IsOn {
			state = "on"
		} else {
			state = "off"
		}
		fmt.Println("commanded and currently", state)
		os.Exit(0)
	}
	fmt.Println("Unknown command arg", os.Args[1])
	os.Exit(1)
}
