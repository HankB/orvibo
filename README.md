# orvibo
Orvibo S20 in Go (golang)

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
Point mode. A PC can then associate with the S20 and exchange messages that share the
desired SSID and password. Described at https://stikonas.eu/wordpress/2015/02/24/reverse-engineering-orvibo-s20-socket/

## Security
Both pairing strategies expose the local WiFi password to snooping. Hopefully no one is
trying to get your credentials when this process is performed.

In normal operation it seems likely that the S20 communicates with a cloud server to
provide remote control of the switch. This exposes the network to ongoing security risk. 
To midigate that, the S20 can be blocked from access int the Internet. Local operation is still
supported.

## Protocol
See details at http://pastebin.com/LfUhsbcS

