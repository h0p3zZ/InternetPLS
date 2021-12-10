package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/go-ping/ping"
)

type postObject struct {
	AuthUser string `json:"auth_user"`
	AuthPass string `json:"auth_pass"`
	Accept   string `json:"accept"`
}

const url = "http://10.10.0.251:8002/?zone=cp_htl"

func runPing(host string) (rtt time.Duration, pktLoss float64) {
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

func connect(userInfo *user) {
	localAddress := getDialInterfaceAddress()

	if localAddress == nil {
		panic("Could not find interface with DNS suffix htl.grieskirchen.local")
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			LocalAddr: localAddress,
		}).Dial, TLSHandshakeTimeout: 10 * time.Second}

	client := &http.Client{
		Transport: transport,
	}

	b, err := json.Marshal(&postObject{userInfo.Username, userInfo.Password, "Anmelden"})
	if err != nil {
		return
	}
	response, postErr := client.Post(url, "application/JSON", bytes.NewBuffer(b))
	if postErr != nil {
		fmt.Println(postErr)
	}
	body, _ := io.ReadAll(response.Body)
	isLoggedIn := bytes.Contains(body, []byte("freigeschalten"))
	if isLoggedIn {
		fmt.Println("Logged in!")
		return
	}
}
