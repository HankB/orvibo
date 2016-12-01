package main

// Simple app to Pair with the Orvibo S20

import (
	"fmt"
	"os"

	s20 "github.com/HankB/orvibo/s20"
)

func usage(progname string, description string) {
	if description != "" {
		fmt.Println(description)
	}
	fmt.Println("Usage:", progname, "SSID password module_ID")
	os.Exit(1)
}

func main() {
	args := os.Args
	if len(args) != 4 {
		usage(args[0], "")
	}
	fmt.Printf("Associating with \"%s\" using \"%s\" naming as \"%s\"\n",
		args[1], args[2], args[3])
	s20.Init(args[1], args[2], args[3])
	s20.Pair()
}
