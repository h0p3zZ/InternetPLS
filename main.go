package main

import (
	"fmt"
)

// find network interface with dns suffic htl.grieskirchen.local
// create socket connection via this interace to the login-url

func main() {

	userInfo := readUserfile()
	connect(userInfo)

	for {
		rtt, pktLoss := runPing("www.google.com")

		if pktLoss == float64(1) {
			fmt.Println("Connection lost. Reconnecting...")
			connect(userInfo)
		} else {
			fmt.Printf("RTT: %v\n", rtt)
		}
	}
}
