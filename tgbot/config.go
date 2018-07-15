package tgbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Token  string
	ChatID int
}

func NewConfig() (*config, error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, fmt.Errorf("Could not read config file %s", "config.json")
	}
	c := &config{}
	if err = json.Unmarshal(file, c); err != nil {
		return nil, fmt.Errorf("Could not parse config file %s", "config.json")
	}
	return c, nil
}
