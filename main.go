package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Print(http.Get("https://google.com"))
}
