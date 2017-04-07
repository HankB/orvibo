package s20_pair_test

import (
	"fmt"
	"os"
	"testing"

	s20 "github.com/HankB/orvibo/s20"
)

var ssid string
var pwd string
var m_id string

// fetch SSID and password from environment for testing, substitute
// some reasonable defaults if not provided.
func InitTest() {
	ssid = os.Getenv("SSID")
	if len(ssid) == 0 {
		ssid = "NO_SSID"
	}
	pwd = os.Getenv("PASSWORD")
	if len(pwd) == 0 {
		pwd = "NO_PWD"
	}
	m_id = os.Getenv("MODULE_ID")
	if len(m_id) == 0 {
		m_id = "NO_MODULE_ID"
	}
	fmt.Printf("SSID=\"%s\", PWD=\"%s\" MID=\"%s\"\n", ssid, pwd, m_id)
}

// TestPair try to pair - keep this last as it fails if
// the host is not associated with an S20 in AP mode.
func TestPair(t *testing.T) {
	InitTest()
	s20.Init(ssid, pwd, m_id)
	s20s := s20.Pair()
	if len(s20s) != 0 {
		t.Fail()
	}
}
