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

	if len(os.Args) == 1 {
		fmt.Println("Usage", os.Args[0])
		os.Exit(0)
	} else if os.Args[1] == "-d" {
		s20s, _ := s20.Discover(replyTimeout)
		for _, s20s := range s20s {
			fmt.Println(s20s)
		}
		os.Exit(0)
	} else if os.Args[1] == "-s" {
		s20s, _ := s20.Discover(replyTimeout)
		e := s20.Subscribe(replyTimeout, s20s[0].IpAddr)
		fmt.Println("done", e)
		os.Exit(0)
	}
	fmt.Println("Unknown command arg", os.Args[1])
	os.Exit(1)
}
