package s20_test

import (
	s20 "github.com/HankB/orvibo/s20"
	txt "github.com/HankB/txtutil"
)

func ExampleDump() {
	txt.Dump(s20.MAGIC)
	// 00000000  68 64                                             |hd|
	// 00000002
}
