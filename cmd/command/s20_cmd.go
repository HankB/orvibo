// Package send commands to the S20
package main

import (
	"fmt"
	"os"

	"github.com/HankB/orvibo/s20"
)

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
		s20s, _ := s20.Discover(2)
		for _, s20 := range s20s {
			fmt.Println(s20)
		}
		os.Exit(0)
	}
	fmt.Println("Unknown command arg", os.Args[1])
	os.Exit(1)
}
