# orvibo
Orvibo S20 in Go (golang)

## Status
Work in progress
Pairing is the present focus. Complete pairing sequence with a fake S20
(https://bitbucket.org/HankB/orvibo-s20 S20-emulate.py) just accomplished.

Pairing with a real S20 just accomplished!

## TODO
* Implement on ON/OFF commands.
* IOmplement alternate pairing method. Present only works for hosts with WiFi 
and when associated with the Orvibo S20.

## Purpose
Provide a reason to write some Go code. Provide capability to manage the Orvibo
S20 WiFi switch. Support comes in two phases. Initially the Orvibo must 'pair' with
the local WiFi network. Second, a program is required to turn the switch on and off.

The switch supports additional functionality (timed on/off.)

## Pairing
* One method is for another host on the network to send out various length messages
where the password is encoded in the length of the messages (or something like that.)
This is probably more appropriate for Android/IOS apps. This method is described at
http://blog.slange.co.uk/orvibo-s20-wifi-power-socket/
* Press the button on the S20 for longer than 5 seconds to put the S20 into (open) Access
Point mode. (In this mode a blue indicator will flash at about 5 Hz. If a red indicator is flashing at that frequency, 
press he button about 4 seconds.) A PC can then associate with the S20 and exchange messages that share the
desired SSID and password. Described at https://stikonas.eu/wordpress/2015/02/24/reverse-engineering-orvibo-s20-socket/ (Note: This page states that communications with the S20 are via the UDP broadcast address. In this code the 
address 1 less the broadcast address - .e.g .254 - was found to work.)

## Security
Both pairing strategies expose the local WiFi password to snooping. Hopefully no one is
trying to get your credentials when this process is performed.

In normal operation it seems likely that the S20 communicates with a cloud server to
provide remote control of the switch. This exposes the network to ongoing security risk.
To mitigate that, the S20 can be blocked from access int the Internet. Local operation is still
supported.

## testing
Some tests are provided in `s20_test.go`. Before running them the following
environment variables should be set.

    export SSID=<your SSID>
    export PASSWORD="<password-for-your-AP"
    export MODULE_ID="hostname-for-S20"

    go test -v s20/s20_test.go

The best part about the test is that if your PC has associated with the S20 in pairing 
mode, it will actually pair the S20 to your network! (Hmmm... Maybe that test should be 
removed.) If the host running the test is not associated with the S20 the test will normally end with

    Error:  dial udp 10.10.100.150:9884->10.10.100.254:48899: bind: cannot assign requested address

## Protocol
See details at http://pastebin.com/LfUhsbcS
