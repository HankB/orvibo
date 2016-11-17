package main

import (
	s20 "github.com/HankB/orvibo/s20"
	"fmt"
	"testing"
)



func main() {
	fmt.Println("Hello")
        // Output:
        // Hello

        s20.Dump(s20.MAGIC)
        // Output:
        // [ 86 64 ]
}

func TestDump(t *testing.T) {
        // this test does not work
        s20.Dump(s20.MAGIC)
        // Output:
        // [ 86 64 ]
}

func ExampleDump() {
       s20.Dump(s20.MAGIC)
        // Output:
        // [ 68 64 ]
}
