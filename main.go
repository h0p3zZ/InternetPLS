package main

import (
	"fmt"
)

const url = "http://10.10.0.251:8002/?zone=cp_htl"

func main() {
	userInfo := readUserfile()

	for {
		rtt, pktLoss := pingRun("www.google.com")

		if pktLoss == float64(1) {
			fmt.Println("Connection lost. Reconnecting...")
			connect(userInfo)
		} else {
			fmt.Printf("RTT: %v\n", rtt)
		}
	}
}
