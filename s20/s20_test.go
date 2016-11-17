package main

import (
	"github.com/HankB/orvibo/s20"
	"fmt"
	"testing"
)



func main() {
	fmt.Println("Hello")
        // Output:
        // Hello
}

func TestXxx(*testing.T) {o
	if s20.Good() != 1 {
		Fail("not good")
	}
}

