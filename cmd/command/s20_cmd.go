// Package send commands to the S20
package main

import (
"os"
"fmt"
)

func usage(errstr string) {
    if len(errstr) > 0 {
        fmt.Println(errstr)
    }
    
}
func main() {

	os.Exit(0)
}
