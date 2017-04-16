# orvibo

Orvibo S20 in Go (golang)

## Warning

This is my first significant Go project and I'm sure I have much to learn. In
particular the file naming and organization (and also package naming) may be
unorthodox. Feel free to submit issues suggesting improvement.

## Status

Work in progress
Pairing is the present focus. Complete pairing sequence with a fake S20
(`https://bitbucket.org/HankB/orvibo-s20` S20-emulate.py) just accomplished.

Pairing with a real S20 just accomplished!

Testing factored and working again. &lt;sigh&gt;

Working with Discover() to get the information required for Subscribe().

## TODO

* Check to see if pairing works with already paired device.
* Work with Discovery.
* Implement on ON/OFF commands.
* Implement alternate pairing method. Present only works for hosts with WiFi
  and when associated with the Orvibo S20.
* Revisit testing and avoid the use of the environment variables.
* Test to see if pairing can be performed while the S20 is already paired.
  (e.g. will the S20 respond to the AT commands if it is not in AP mode)

## Purpose

Provide a reason to write some Go code. Provide capability to manage the Orvibo
S20 WiFi switch. Support comes in two phases. Initially the Orvibo must 'pair' with
the local WiFi network. Second, a program is required to turn the switch on and off.

The switch supports additional functionality (timed on/off.)

## Pairing

* One method is for another host on the network to send out various length messages
  where the password is encoded in the length of the messages (or something like that.)
  This is probably more appropriate for Android/IOS apps. This method is described at
  `http://blog.slange.co.uk/orvibo-s20-wifi-power-socket/`
* Press the button on the S20 for longer than 5 seconds to put the S20 into (open)
  Access Point mode. (In this mode a blue indicator will flash at about 5 Hz.
  If a red indicator is flashing at that frequency, press he button about 4
  seconds.) A PC can then associate with the S20 and exchange messages that
  share the desired SSID and password. Described at
  `https://stikonas.eu/wordpress/2015/02/24/reverse-engineering-orvibo-s20-socket/`
  (Note: This page states that communications with the S20 are via the UDP
  broadcast address. In this code the address 1 less the broadcast address -
  .e.g .254 - was found to work.) Maybe I have more to learn about networking.

## Security

Both pairing strategies expose the local WiFi password to snooping. Hopefully no one is
trying to get your credentials when this process is performed. (TODO: see if
a new SSID/password can be provided once paired.)

In normal operation it seems likely that the S20 communicates with a cloud server to
provide remote control of the switch. This exposes the network to ongoing security risk.
To mitigate that, the S20 can be blocked from access int the Internet. Local operation still works.

## Testing

Test files are presently separate from project sources and broken into two
groups. First are those that do not require particular environmental
configuration and can be expected to run without error. The second group
require an S20 on the network and will not pass without a properly configured
S20 on the LAN.

### Stand alone test

``` text
hbarta@olive:~/Documents/go-work/src/github.com/HankB/orvibo$ go test -v github.com/HankB/orvibo/s20
=== RUN   TestIsThisHost
--- PASS: TestIsThisHost (0.00s)
=== RUN   ExampleInit
--- PASS: ExampleInit (0.00s)
PASS
ok  	github.com/HankB/orvibo/s20	0.002s
hbarta@olive:~/Documents/go-work/src/github.com/HankB/orvib
```

### Test with S20

#### Pairing

This first test exercises the pairing code that is used when the S20 is operating
as an AP. To achieve this, power the S20 and long press the button twice until the
indicator LED is flashing blue rapidly. The PC needs a WiFi interface associated
with `WiWo-S20` (or access to a network that is connected to this AP.)
Without environment variables the code will
provide likely useless setup parameters. At the completion of the test the S20 will
reset and shut down AP operation. Test output looks like:

``` text
hbarta@olive:~/Documents/go-work/src/github.com/HankB/orvibo$ go test -v github.com/HankB/orvibo/s20_pair_test
=== RUN   TestPair
SSID="NO_SSID", PWD="NO_PWD" MID="NO_MODULE_ID"
S20 S/W Version '+ok=08 (2015-04-28 16:57 16B)

'
S20 MID '+ok=HF-LPB100

'
S20 send MID +ok

S20 MID +ok=NO_MODULE_ID

--- PASS: TestPair (0.38s)
PASS
ok  	github.com/HankB/orvibo/s20_pair_test	0.386s
hbarta@olive:~/Documents/go-work/src/github.com/HankB/orvibo$
```

If the test is performed without an S20 connected the results will look like:

``` text
hbarta@olive:~/Documents/go-work/src/github.com/HankB/orvibo$ go test -v github.com/HankB/orvibo/s20_pair_test
=== RUN   TestPair
SSID="NO_SSID", PWD="NO_PWD" MID="NO_MODULE_ID"
Error:  dial udp 10.10.100.150:9884->10.10.100.254:48899: bind: cannot assign requested address
exit status 1
FAIL	github.com/HankB/orvibo/s20_pair_test	0.003s
hbarta@olive:~/Documents/go-work/src/github.com/HankB/orvibo$ 
```

If the following environment variables are set before running this test, the
S20 will be configured accordingly.

``` text
    export SSID=<your SSID>
    export PASSWORD="<password-for-your-AP"
    export MODULE_ID="hostname-for-S20"
```

The `MODULE_ID` will be sent to the S20 and it will use it when it uses DHCP to
request an IP address after associating with your WiFi AP. Some routers will
display this name and resolve DNS requests to allow access to the S20 using that
name.

#### Discovery

Identify the S20 devices on the network.

``` text
hbarta@olive:~/Documents/go-work/src/github.com/HankB/orvibo$ go run cmd/command/s20_cmd.go -d
Sent 6 bytes
adding 192.168.1.160:10000 count 1 on 1
00000000  68 64 00 2a 71 61 00 ac  cf 23 55 fe 22 20 20 20  |hd.*qa...#U."   |
00000010  20 20 20 22 fe 55 23 cf  ac 20 20 20 20 20 20 53  |   ".U#..      S|
00000020  4f 43 30 30 35 5e a5 19  bc 01                    |OC005^....|
0000002a
adding 192.168.1.212:10000 count 2 on 0
00000000  68 64 00 2a 71 61 00 ac  cf 23 36 02 0e 20 20 20  |hd.*qa...#6..   |
00000010  20 20 20 0e 02 36 23 cf  ac 20 20 20 20 20 20 53  |   ..6#..      S|
00000020  4f 43 30 30 35 52 ff 9b  dc 00                    |OC005R....|
0000002a
{{192.168.1.160 10000 } ac:cf:23:36:02:0e 0e:02:36:23:cf:ac true}
{{192.168.1.212 10000 } ac:cf:23:36:02:0e 0e:02:36:23:cf:ac false}
hbarta@olive:~/Documents/go-work/src/github.com/HankB/orvibo$
```

## Protocol

See details at `http://pastebin.com/LfUhsbcS` (Now part of project at orvibo_wifi_socket.txt)

## Errata

Returning to the project after 4 months of not working on it, I don't recall what
the different files are for. I do recall that I was grappling with package names,
testing, commands that use a 'library package' and my general ignorance of these
things WRT Go. Following is a list of the files and what I think they do.

* `cmd/` contains the programs that use the package to do things with the Orvibo S20.
* `cmd/commands/s20_cmd.go` Program to interact with an already paired S20.
* `cmd/pair/s20_pair.go` Program to pair an S20 (e.g. configure it for the local WiFi access point.)
* `s20/` contains the files that provide the `s20` package.
* `s20/cmd.go` Various functions used for interacting with a paired S20.
* `s20/pair.go` Functions used to pair with an S20 which is operating as an AP.
* `s20/s20.go` Predefined variables and utility functions such as Init() and Get().
* `s20/s20_test.go` provides tests that do not require a configured S20
* `s20_pair_test/` contains test that requires a configured S20
* `s20_pair_test/s20_pair.go` Test pairing with a real S20.

## References

* Protocol description kindly provided at `https://pastebin.com/LfUhsbcS` and
  copied here should the original ever be deleted.
* Python implementation `https://github.com/happyleavesaoc/python-orvibo`
* Partial Python implementation `https://bitbucket.org/HankB/orvibo-s20`