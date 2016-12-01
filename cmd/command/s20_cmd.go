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

	s20s, _ := s20.Discover(10)
	for _, s20 := range s20s {
		fmt.Println(s20)
	}
	os.Exit(0)
}
