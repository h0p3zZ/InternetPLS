package main

import (
	//	b64 "encoding/base64"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-ping/ping"
)

const url = "http://10.10.0.251:8002/?zone=cp_htl"

// ToDo read console input
// ToDo write to config file
// ToDo split up in seperate go files
// ToDo encode password (on console as well as in the conf file)

func main() {
	b, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("File not found!")
	} else {
		fmt.Println(string(b))
	}

	for {
		rtt, pktLoss := pingRun("www.google.com")

		if pktLoss == float64(1) {
			fmt.Println("Connection lost. Reconnecting...")
			connect()
		} else {
			fmt.Printf("RTT: %v\n", rtt)
		}
	}
}

type PostObject struct {
	AuthUser string `json:"auth_user"`
	AuthPass string `json:"auth_pass"`
	Accept   string `json:"accpe"`
}

func pingRun(host string) (rtt time.Duration, pktLoss float64) {
	pinger, pingerErr := ping.NewPinger(host)
	if pingerErr != nil {
		fmt.Println(pingerErr)
		panic("Could not create new Pinger to 'www.google.com'")
	}

	pinger.SetPrivileged(true)
	pinger.Count = 4
	pingerRunErr := pinger.Run()
	if pingerRunErr != nil {
		panic("Could not run pinger for some reason. Please try again.")
	}

	stats := pinger.Statistics()
	return stats.AvgRtt, stats.PacketLoss
}

func connect() {
	const username = ""
	const password = ""
	b, err := json.Marshal(&PostObject{username, password, "Anemlden"})
	if err != nil {
		return
	}
	resp, postErr := http.Post(url, "application/JSON", bytes.NewBuffer(b))
	fmt.Println(resp)
	fmt.Println(postErr)
}
