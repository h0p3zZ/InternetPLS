package main

import (
	//	b64 "encoding/base64"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
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
		fmt.Println(b)
	}
	fmt.Println(string(b))
	for {
		start := time.Now()
		_, err := net.LookupIP("8.8.8.8")
		if err != nil {
			fmt.Println("Connection timed out")
			connect()
		}
		end := time.Now()
		fmt.Println(end.Sub(start))
		time.Sleep(1 * time.Second)
	}
}

type PostObject struct {
	AuthUser string `json:"auth_user"`
	AuthPass string `json:"auth_pass"`
	Accept   string `json:"accpe"`
}

func connect() {
	const username = ""
	const password = ""
	b, err := json.Marshal(&PostObject{username, password, "Anemlden"})
	if err != nil {
		return
	}
	resp, _ := http.Post(url, "application/JSON", bytes.NewBuffer(b))
	fmt.Println(resp)
}
