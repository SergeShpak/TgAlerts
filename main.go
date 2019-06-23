package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	updates, err := botAPI.GetUpdatesChan(nil)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		log.Println(update.Message.Text)
	}
}

func readToken() (string, error) {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		return "", err
	}
	return string(token), nil
}
