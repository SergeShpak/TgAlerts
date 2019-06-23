package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	token, err := readToken()
	if err != nil {
		panic(fmt.Sprintf("Failed to read the token file: %v", err))
	}
	botAPI, err := NewBotAPI(token)
	if err != nil {
		panic(fmt.Sprintf("Failed to create a new bot: %v", err))
	}
	if err := botAPI.GetMe(); err != nil {
		panic(err)
	}
}

func readToken() (string, error) {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		return "", err
	}
	return string(token), nil
}
