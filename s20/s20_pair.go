// Package s20 to support operations involving the Orvibo S20
package s20

import (
	"fmt"
	"net"
	"os"
	"time"

	txt "github.com/HankB/txtutil"
)

//type bufferSender func([]byte)

var initiator = []byte("HF-A11ASSISTHREAD") // initiate conversation with s20
var sendSSID = "AT+WSSSID="                 // followed by SSID and '\r'
var sendPWD = "AT+WSKEY=WPA2PSK,AES,"       // followed by password and '\n'
var sendSTA = []byte("AT+WMODE=STA\n")      // complete to set s20 mode
var sendRST = []byte("AT+Z\n")              // request s20 to reset
var okReply = []byte("+ok")                 // ACK

var serverAddr *net.UDPAddr
var ourAddr *net.UDPAddr
var err error
var conn *net.UDPConn

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

	// tell s20 we want to talk
	sendRcv(initiator)
	send(okReply)

	// send SSID message
	sendSSID = sendSSID + ssid + "\r"
	fmt.Printf("SSID=\"%s\" result\"%s\"\n", ssid, sendSSID)
	sendRcv([]byte(sendSSID))

	// send WAP password
	sendPWD = sendPWD + pwd + "\n"
	fmt.Printf("PWD=\"%s\" result\"%s\"\n", pwd, sendPWD)
	sendRcv([]byte(sendPWD))

	// switch the s20 to station mode (from AP)
	sendRcv(sendSTA)

	// and restart
	sendRcv(sendRST)

	return found
}

//setupPairing will open a UDP port on which a broadcast can be sent
// and the response read.
func setupPairing() {
	// get server handle
	server := fmt.Sprintf("%s:%d", ip, udpSndPort)
	serverAddr, err = net.ResolveUDPAddr("udp", server)
	checkErr(err)
	client := fmt.Sprintf("%s:%d", "10.10.100.150", udpRcvPort)
	ourAddr, err = net.ResolveUDPAddr("udp", client)
	checkErr(err)
	fmt.Println("calling DialUDP()")
	conn, err = net.DialUDP("udp", ourAddr, serverAddr)
	checkErr(err)
	fmt.Println("Established UDP socket")
}

// Send a message and wait up to X seconds for a response
func sendRcv(b []byte) []byte {
	reply := make([]byte, 1024)
	var readLen int
	fmt.Println("sendRcv s:")
	txt.Dump(string(b))

	_, err = conn.Write(b)
	checkErr(err)
	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readLen, err = conn.Read(reply)
	checkErr(err)
	fmt.Printf("reply len %d, \"%s\"\n", readLen, reply)
	return reply

}

func send(b []byte) {
	reply := make([]byte, 1024)
	var readLen int

	_, err = conn.Write(b)
	checkErr(err)
	fmt.Printf("reply len %d, \"%s\"\n", readLen, reply)
}
