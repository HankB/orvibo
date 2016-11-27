// Package s20 to support operations involving the Orvibo S20
package s20

import (
	"fmt"
	"net"
	"os"
	"time"
)

//type bufferSender func([]byte)

/* A Simple function to verify error */
/* copied from https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/ */
func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

//Pair will send messages that will encode SSID and password in message length
func Pair() []string {
	found := make([]string, 0)
	setupPairing()
	sendRcv(initiator)
	return found
}

var initiator = []byte("HF-A11ASSISTHREAD")   // initiate conversation with s20
var okReply = []byte("+ok")                   // ACK
var sendSSID = []byte("AT+WSSSID=")           // followed by SSID and '\r'
var sendPWD = []byte("AT+WSKEY=WPA2PSK,AES,") // followed by password and '\r'
var sendSTA = []byte("AT+WMODE=STA\n")        // complete to set s20 mode
var sendRST = []byte("AT+Z\n")

var serverAddr *net.UDPAddr
var ourAddr *net.UDPAddr
var err error
var conn *net.UDPConn

//setupPairing will open a UDP port on which a broadcast can be sent
// and the response read.
func setupPairing() {
	// get server handle
	server := fmt.Sprintf("127.0.0.1:%d", udpSndPort)
	serverAddr, err = net.ResolveUDPAddr("udp", server)
	checkErr(err)
	client := fmt.Sprintf("127.0.0.1:%d", udpRcvPort)
	ourAddr, err = net.ResolveUDPAddr("udp", client)
	checkErr(err)
	conn, err = net.DialUDP("udp", ourAddr, serverAddr)
	checkErr(err)
	fmt.Println("Established UDP socket")
}

// Send a message and waity up to X seconds for a response
func sendRcv(b []byte) []byte {
	reply := make([]byte, 1024)
	var readLen int

	_, err = conn.Write(b)
	checkErr(err)
	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readLen, err = conn.Read(reply)
	checkErr(err)
	fmt.Printf("reply len %d, \"%s\"\n", readLen, reply)
	return reply

}
