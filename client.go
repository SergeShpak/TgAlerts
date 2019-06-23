package main

import (
	"fmt"
	"net/http"
)

const TelegramAPIAddr = "https://api.telegram.org/bot"

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) (*Client, error) {
	client := &Client{
		httpClient: &http.Client{},
		token:      token,
	}
	return client, nil
}

func (c *Client) DoRequest(method string, endpoint string) error {
	r, err := http.NewRequest(method, c.composeURL(endpoint), nil)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

func (c *Client) composeURL(endpoint string) string {
	return fmt.Sprintf("%s%s/%s", TelegramAPIAddr, c.token, endpoint)
}
