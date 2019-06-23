package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/SergeyShpak/TgAlerts/types"
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

func (c *Client) DoRequest(method string, endpoint string, params url.Values, response interface{}) error {
	r, err := http.NewRequest(method, c.composeURL(endpoint), nil)
	r.URL.RawQuery = params.Encode()
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if response == nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("body: ", string(body))
	var responseWrapper types.Response
	responseWrapper.Result = response
	if err := json.Unmarshal(body, &responseWrapper); err != nil {
		return err
	}
	return nil
}

func (c *Client) composeURL(endpoint string) string {
	return fmt.Sprintf("%s%s/%s", TelegramAPIAddr, c.token, endpoint)
}
