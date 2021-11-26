package main

import (
	b64 "encoding/base64"

	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func readUserfile() user {
	b, err := os.ReadFile("./user.json")
	if err == nil {
		var readUser user
		err := json.Unmarshal(b, &readUser)
		if err != nil {
			fmt.Println(err)
		} else {
			decodedPw, _ := b64.StdEncoding.DecodeString(readUser.Password)
			readUser.Password = string(decodedPw)
			return readUser
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _, _ := reader.ReadLine()
	fmt.Print("Password: ")
	encodedPassword, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")

	password := b64.StdEncoding.EncodeToString(encodedPassword)
	newUser := user{string(username), string(password)}

	s, _ := json.Marshal(newUser)
	os.WriteFile("./user.json", []byte(s), 0666)

	return newUser
}
