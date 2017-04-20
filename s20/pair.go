// Package s20 to support operations involving the Orvibo S20
package s20

import (
	"fmt"
	"net"
	"os"
	"time"
)

//type bufferSender func([]byte)

var initiator = []byte("HF-A11ASSISTHREAD") // initiate conversation with s20
var sendSSID = "AT+WSSSID="                 // followed by SSID and '\r'
var sendPWD = "AT+WSKEY=WPA2PSK,AES,"       // followed by password and '\n'
var sendMID = "AT+WRMID="                   // complete to set s20 module ID
var queryMID = "AT+MID\n"                   // send to squery s20 module ID
var sendSTA = []byte("AT+WMODE=STA\n")      // complete to set s20 mode
var sendRST = []byte("AT+Z\n")              // request s20 to reset
var queryVER = []byte("AT+LVER\n")          // request S/W version
var okReply = []byte("+ok")                 // ACK (sent or received)

var serverAddr *net.UDPAddr
var ourAddr *net.UDPAddr
var err error
var conn *net.UDPConn

/* A Simple function to verify error */
/* copied from https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/ */
func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

// revised error handling ... Print out an error mewsage and exit. There is No
// recovery programmed for any errors that occur.

//Pair will send messages that will encode SSID and password in message length
func Pair() []string {
	found := make([]string, 0)
	setupPairing()

	// tell s20 we want to talk
	sendRcv(initiator)
	send(okReply)

	// request S/V version
	fmt.Printf("S20 S/W Version '%s'\n", sendRcv(queryVER))

	// request module ID
	fmt.Printf("S20 MID '%s'\n", sendRcv([]byte(queryMID)))

	// send module name
	sendMID = sendMID + mid + "\n"
	fmt.Printf("S20 send MID %s", sendRcv([]byte(sendMID)))

	// confirm new module ID
	fmt.Printf("S20 MID %s", sendRcv([]byte(queryMID)))

	// send SSID message
	sendSSID = sendSSID + ssid + "\r"
	sendRcv([]byte(sendSSID))

	// send WAP password
	sendPWD = sendPWD + pwd + "\n"
	sendRcv([]byte(sendPWD))

	// switch the s20 to station mode (from AP)
	sendRcv(sendSTA)

	// and restart
	send(sendRST)

	return found
}

//setupPairing will open a UDP port on which a broadcast can be sent
// and the response read.
func setupPairing() {
	// get server handle
	server := fmt.Sprintf("%s:%d", s20IP, udpSndPort)
	serverAddr, err = net.ResolveUDPAddr("udp", server)
	checkErr(err)
	client := fmt.Sprintf("%s:%d", ourIP, udpRcvPort)
	ourAddr, err = net.ResolveUDPAddr("udp", client)
	checkErr(err)
	// fmt.Println("calling DialUDP()")
	conn, err = net.DialUDP("udp", ourAddr, serverAddr)
	checkErr(err)
	// fmt.Println("Established UDP socket")
}

// Send a message and wait up to X seconds for a response
func sendRcv(b []byte) []byte {
	reply := make([]byte, 1024)
	var length int
	// fmt.Println("sendRcv s:")
	// txt.Dump(string(b))

	length, err = conn.Write(b)
	checkErr(err)
	if length != len(b) {
		fmt.Printf("Error: Write returned %d bytes, s/b %d\n", length, len(b))
		os.Exit(1)
	}
	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	checkErr(err)
	length, err = conn.Read(reply)
	checkErr(err)
	// fmt.Printf("reply length %d, \"%s\"\n", length, reply)
	return reply

}

func send(b []byte) {
	var length int
	length, err = conn.Write(b)
	checkErr(err)
	if length != len(b) {
		fmt.Printf("Error: Write returned %d bytes, s/b %d\n", length, len(b))
		os.Exit(1)
	}
	// fmt.Printf("reply len %d, \"%s\"\n", readLen, reply)
}
