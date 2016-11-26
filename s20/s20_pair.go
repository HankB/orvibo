// Package s20 to support operations involving the Orvibo S20
package s20

type bufferSender func([]byte)

//Pair will send messages that will encode SSID and password in message length
func Pair() []string {
	found := make([]string, 0)
	return found
}

// format and send messages that encode password
func sendPairMsgs(f bufferSender) {

}
func encodePasswordChar(c byte) []byte {

}
