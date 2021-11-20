package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	b, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("File not found!")
	}
	fmt.Println(string(b))
	for {
		start := time.Now()
		_, err := net.LookupIP("8.8.8.8")
		if err != nil {
			fmt.Println("Connection timed out")
		}
		end := time.Now()
		fmt.Println(end.Sub(start))
		time.Sleep(1 * time.Second)
	}
}
